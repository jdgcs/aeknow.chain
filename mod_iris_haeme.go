package main

import (
	"bufio"
	//"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"io"
	"io/ioutil"

	//"unsafe"

	//"net/url"
	"os"
	"os/exec"

	//"path"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/kataras/iris/v12"
)

type PageBlog struct {
	Account         string
	PageContent     template.HTML
	PageTitle       string
	PageDescription string
	PageTags        string
	PageCategories  string
	EditPath        string
	PreLink         template.HTML
	NextLink        template.HTML
	TagsLink        template.HTML
	PubTime         string
	AuthorLink      template.HTML
}

type PageList struct {
	Account         string
	PageContent     template.HTML
	PageTitle       string
	PageDescription string
	PageTags        string
	PageCategories  string
	EditPath        string
	PreLink         template.HTML
	NextLink        template.HTML
	TagsLink        template.HTML
	PubTime         string
	AuthorLink      template.HTML
}

type IPFSConfig struct {
	Identity  IPFSIdentity
	Datastore IPFSDatastore
}

type IPFSIdentity struct {
	PeerID  string
	PrivKey string
}

type IPFSDatastore struct {
	StorageMax         string
	StorageGCWatermark string
	GCPeriod           string
	Bootstrap          []string
}

type SiteConfig struct {
	Title             string
	Subtitle          string
	Description       string
	Author            string
	AuthorDescription string
	Theme             string
}

var MyIPFSConfig IPFSConfig
var MySiteConfig SiteConfig
var lastIPFS string
var MyUsername string
var MyAENS string

func getSiteConfig() SiteConfig {

	configFilePath := ""
	if ostype == "windows" {
		configFilePath = "data\\site\\" + globalAccount.Address + "\\site.json"
	} else {
		configFilePath = "./data/site/" + globalAccount.Address + "/site.json"
	}
	_, err := os.Stat(configFilePath)

	if err != nil {
		configFilePath = "./data/config_default.json"
	}

	JsonParse := NewJsonStruct()
	readConfigfile := SiteConfig{}
	JsonParse.Load(configFilePath, &readConfigfile)

	return readConfigfile

}

func getIPFSConfig() IPFSConfig {
	configFilePath := ""
	if ostype == "windows" {
		configFilePath = "data\\repo\\config"
	} else {
		configFilePath = "./data/repo/config"
	}
	_, err := os.Stat(configFilePath)

	if err != nil {
		configFilePath = "./data/config_default.json"
	}

	JsonParse := NewJsonStruct()
	readConfigfile := IPFSConfig{}
	JsonParse.Load(configFilePath, &readConfigfile)

	return readConfigfile

}

//change the site settings, such as theme, name and description
func iSetSite(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}

	ctx.ViewData("", MySiteConfig)
	ctx.View("haeme_settings.php")
}

func iSaveSetSite(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	title := ctx.FormValue("title")
	title = strings.TrimSpace(title)
	subtitle := ctx.FormValue("subtitle")
	subtitle = strings.TrimSpace(subtitle)
	sitedescription := ctx.FormValue("sitedescription")
	author := ctx.FormValue("author")
	authordescription := ctx.FormValue("authordescription")
	theme := ctx.FormValue("theme")

	jsoncontent := `{
  "Title":"` + title + `", 
  "Subtitle":"` + subtitle + `",
  "Description":"` + sitedescription + `",
  "Author":"` + author + `",
  "AuthorDescription":"` + authordescription + `",
  "Theme":"` + theme + `"
}`

	if ostype == "windows" {
		//save site config file
		targetFile := ".\\data\\site\\" + globalAccount.Address + "\\site.json"
		err := ioutil.WriteFile(targetFile, []byte(jsoncontent), 0644)
		if err != nil {
			panic(err)
		}
	} else {
		//save site config file
		targetFile := "./data/site/" + globalAccount.Address + "/site.json"
		err := ioutil.WriteFile(targetFile, []byte(jsoncontent), 0644)
		if err != nil {
			panic(err)
		}
	}
	MySiteConfig = getSiteConfig() //重新读取网站设置
	configHugo()
	ctx.HTML("<h1>Site config has been saved.</h1>")
	//ctx.ViewData("", MySiteConfig)
	//ctx.View("haeme_settings.php")
}

//show the defult homepage
func iHaeme(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	//fmt.Println("Haeme")
	myPage := PageWallet{PageId: 23, Account: globalAccount.Address, PageTitle: "Haeme"}
	ctx.ViewData("", myPage)
	ctx.View("transaction.php")
}

func iBlog(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	fmt.Println("Haeme")
	myPage := PageBlog{Account: globalAccount.Address, PageTitle: "", EditPath: ""}

	ctx.ViewData("", myPage)
	ctx.View("haeme_blog.php")
}

func iView(ctx iris.Context) {
	//View the page content with 2 parameters.
	if !checkLogin(ctx) {
		//	return
	}

	hash := ctx.FormValue("hash")
	pubkey := ctx.FormValue("pubkey")
	tag := ctx.FormValue("tag")
	viewtype := ctx.FormValue("viewtype")

	dbpath := "./data/accounts/" + pubkey + "/public.db"
	//fmt.Println(dbpath)

	if hash != "" && pubkey != "" {
		ViewContent(ctx, hash, pubkey, dbpath)
	}

	if tag != "" && pubkey != "" {
		ViewTag(ctx, tag, pubkey, dbpath)
	}
	if viewtype == "author" && pubkey != "" {
		ViewHome(ctx, pubkey, dbpath)
	}

}

func ViewHome(ctx iris.Context, pubkey string, dbpath string) {
	//View the Homepage of a user

	if !FileExist(dbpath) {
		ctx.HTML("<h1>No such sqlite db." + dbpath + "</h1>")
		return
	}
	querystr := "SELECT name FROM author ORDER BY aid DESC LIMIT 1"
	db, err := sql.Open("sqlite3", dbpath)
	checkError(err)
	rows, err := db.Query(querystr)
	checkError(err)

	//var pubkey string
	//var bio string
	//var ipns string
	var name string
	PageContent := ""

	for rows.Next() {
		err = rows.Scan(&name)
		checkError(err)
		PageContent = "PubKey: " + pubkey + ", Author: " + name

	}

	myPage := PageList{Account: globalAccount.Address, PageTitle: name, PageContent: template.HTML(PageContent)}

	ctx.ViewData("", myPage)
	ctx.View("haeme_index.php")
	db.Close()

	//TODO: 用户页面布局，按时间展示
}

func ViewTag(ctx iris.Context, tag string, pubkey string, dbpath string) {
	//View the tag list

	if !FileExist(dbpath) {
		ctx.HTML("<h1>No such sqlite db." + dbpath + "</h1>")
		return
	}

	db, err := sql.Open("sqlite3", dbpath)
	checkError(err)

	querystr := "SELECT title,author,keywords,aid,abstract,filetype,pubtime,hash FROM aek WHERE keywords LIKE '%" + tag + "%' OR title LIKE '%" + tag + "%' ORDER BY pubtime DESC"
	fmt.Println(querystr)
	rows, err := db.Query(querystr)
	checkError(err)
	var title string
	var author string
	var keywords string
	var aid int
	var abstract sql.NullString
	var filetype sql.NullString
	var pubtime string
	var hash string

	PageContent := ""

	for rows.Next() {
		err = rows.Scan(&title, &author, &keywords, &aid, &abstract, &filetype, &pubtime, &hash)
		checkError(err)
		PageContent = PageContent + "<li><a href=/view?pubkey=" + pubkey + "&hash=" + hash + ">" + title + "</a> - <div class=pubtime>" + pubtime + "</div></li>"
		//fmt.Println(title + author + keywords + strconv.Itoa(aid) + abstract.String + filetype.String + pubtime)

	}
	db.Close()

	title = "List of tag: " + tag
	myPage := PageList{Account: globalAccount.Address, PageTitle: title, PageContent: template.HTML(PageContent)}

	ctx.ViewData("", myPage)
	ctx.View("haeme_pagelist.php")
	fmt.Println("Haeme")
}

func ViewContent(ctx iris.Context, hash string, pubkey string, dbpath string) {
	if hash == "" || pubkey == "" {
		ctx.HTML("<h1>No such parameter.</h1> ")
		return
	}

	if !FileExist(dbpath) {
		ctx.HTML("<h1>No such sqlite db." + dbpath + "</h1>")
		return
	}

	db, err := sql.Open("sqlite3", dbpath)
	checkError(err)

	querystr := "SELECT title,author,keywords,aid,abstract,filetype,pubtime,authorname FROM aek WHERE hash='" + hash + "'"
	//fmt.Println(querystr)
	rows, err := db.Query(querystr)

	checkError(err)
	var title string
	var author string
	var keywords string
	var aid int
	var abstract sql.NullString
	var filetype sql.NullString
	var pubtime string
	var authorname string

	for rows.Next() {
		err = rows.Scan(&title, &author, &keywords, &aid, &abstract, &filetype, &pubtime, &authorname)
		checkError(err)
	}
	db.Close()

	myTime, err := strconv.ParseInt(pubtime, 10, 64)
	tm := time.Unix(myTime, 0)
	pubtime = tm.Format("2006-01-02 15:04:05") //2018-07-11 15:10:19

	pubtime = strings.Replace(pubtime, "T", " ", -1)
	pubtime = strings.Replace(pubtime, "Z", " ", -1)

	PreLink := template.HTML(GetPreLink(aid, pubkey))
	NextLink := template.HTML(GetNextLink(aid, pubkey))
	TagsLink := template.HTML(GetTagLink(keywords, pubkey))
	AuthorLink := template.HTML("<a href=/view?pubkey=" + pubkey + "&viewtype=author>" + authorname + "</a>")
	sh := shell.NewShell(NodeConfig.IPFSAPI)
	rc, err := sh.Cat("/ipfs/" + hash)
	s, err := copyToString(rc)
	checkError(err)
	//fmt.Println(s)

	myPage := PageBlog{AuthorLink: AuthorLink, PubTime: pubtime, TagsLink: TagsLink, PreLink: PreLink, NextLink: NextLink, Account: globalAccount.Address, PageTitle: title, PageContent: template.HTML(s), PageTags: keywords, EditPath: author}

	ctx.ViewData("", myPage)
	//if filetype.String == "markdown" {
	ctx.View("haeme_page.php")
	//}

	if filetype.String == "html" {
		ctx.View("haeme_page_html.php")
	}
	fmt.Println("Haeme")
}

func copyToString(r io.Reader) (res string, err error) {
	var sb strings.Builder
	if _, err = io.Copy(&sb, r); err == nil {
		res = sb.String()
	}
	return
}

func GetTagLink(tags string, pubkey string) string {
	//Generate tag links
	TmpStr := strings.Split(tags, ",")
	taglink := ""
	for i := 0; i < len(TmpStr); i++ {
		taglink = taglink + "<a href=/view?pubkey=" + pubkey + "&tag=" + TmpStr[i] + ">" + TmpStr[i] + "</a>,"
	}

	return taglink
}

func GetPreLink(aid int, pubkey string) string {
	//Get the previous page and link
	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite3", dbpath)
	checkError(err)
	querystr := "SELECT title,hash FROM aek WHERE aid<" + strconv.Itoa(aid) + " ORDER BY aid desc LIMIT 1;"

	rows, err := db.Query(querystr)
	checkError(err)
	title := ""
	hash := ""

	for rows.Next() {
		err = rows.Scan(&title, &hash)
		checkError(err)
	}

	db.Close()
	return "<a href=/view?pubkey=" + pubkey + "&hash=" + hash + ">" + title + "</a>"
}

func GetNextLink(aid int, pubkey string) string {
	//Get the next page and link
	dbpath := "./data/accounts/" + pubkey + "/public.db"
	db, err := sql.Open("sqlite3", dbpath)
	checkError(err)
	querystr := "SELECT title,hash FROM aek WHERE aid>" + strconv.Itoa(aid) + " ORDER BY aid asc LIMIT 1;"

	rows, err := db.Query(querystr)
	checkError(err)
	title := ""
	hash := ""

	for rows.Next() {
		err = rows.Scan(&title, &hash)
		checkError(err)
	}

	db.Close()
	return "<a href=/view?pubkey=" + pubkey + "&hash=" + hash + ">" + title + "</a>"
}
func iSaveBlog(ctx iris.Context) {

	if !checkLogin(ctx) {
		return
	}

	title := ctx.FormValue("title")
	title = strings.TrimSpace(title)

	categories := ctx.FormValue("categories")
	categories = strings.Replace(categories, "，", ",", -1)
	tags := ctx.FormValue("tags")
	tags = strings.Replace(tags, "，", ",", -1)
	content := ctx.FormValue("content")
	description := ctx.FormValue("description")
	//editpath := ctx.FormValue("editpath")
	//draft := ctx.FormValue("draft")

	body := html.EscapeString(content)
	description = html.EscapeString(description)

	sh := shell.NewShell(NodeConfig.IPFSAPI)

	hash, err := sh.Add(strings.NewReader(body))
	fmt.Println("posted hash: " + hash)
	checkError(err)

	dbpath := "./data/accounts/" + globalAccount.Address + "/public.db"
	db, err := sql.Open("sqlite3", dbpath)
	checkError(err)
	//pubtime := t.UTC().Format(time.UnixDate)
	//pubtime := time.Now().Unix()
	filesize := strconv.Itoa(strings.Count(body, ""))
	pubtime := strconv.FormatInt(time.Now().Unix(), 10)
	sql_insert := "INSERT INTO aek(title,abstract,hash,keywords,author,pubtime,filetype,authorname,filesize) VALUES('" + title + "','" + description + "','" + hash + "','" + tags + "','" + globalAccount.Address + "'," + pubtime + ",'markdown','" + MyUsername + "'," + filesize + ")"
	fmt.Println(sql_insert)
	_, err = db.Exec(sql_insert)

	checkError(err)
	db.Close()
	//sh := shell.NewShell(NodeConfig.IPFSAPI)
	pubfile, err := os.Open(dbpath)
	cid, err := sh.Add(pubfile)
	lastIPFS = cid
	UpdateConfigs(globalAccount.Address, "LastIPFS", cid) //upfate lastipfs
	fmt.Println("Pub: " + cid)
	//err := ioutil.WriteFile(targetFile, []byte(body), 0644)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Posted")
		//iBuildSite(ctx)
		//iBuildSite(ctx)
	}
	/*
		f, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_TRUNC, 0666)
		defer f.Close()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			_, err = f.Write([]byte(body))
			iUpdateStatic(ctx)
		}*/

	//TODO:detail the section of title and other field in UI, and polish UI
}

func iNewBlog(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	myPage := PageBlog{Account: globalAccount.Address, PageTitle: "", PageContent: "", PageTags: "", PageCategories: "", EditPath: ""}
	//myPage := PageWallet{PageId: 23, Account: globalAccount.Address, PageTitle: "Haeme"}
	ctx.ViewData("", myPage)
	ctx.View("haeme_newblog.php")
}

func iDelBlog(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	pageaddress := ctx.URLParam("pageaddress")
	fmt.Println("pageaddress: " + pageaddress)
	tmpPath := strings.Split(pageaddress, "/")
	fileName := ""

	if ostype == "windows" {
		fileName = ".\\data\\site\\" + globalAccount.Address + "\\content\\post\\" + tmpPath[len(tmpPath)-2] + ".md"
	} else {
		fileName = "./data/site/" + globalAccount.Address + "/content/post/" + tmpPath[len(tmpPath)-2] + ".md"
	}

	ctx.HTML(fileName + " was not deleted yet, it's a test.")
}

func iEditBlog(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	pageaddress := ctx.URLParam("pageaddress")
	fmt.Println("pageaddress: " + pageaddress)
	tmpPath := strings.Split(pageaddress, "/")
	fileName := ""

	if ostype == "windows" {
		fileName = ".\\data\\site\\" + globalAccount.Address + "\\content\\post\\" + tmpPath[len(tmpPath)-2] + ".md"
	} else {
		fileName = "./data/site/" + globalAccount.Address + "/content/post/" + tmpPath[len(tmpPath)-2] + ".md"
	}

	mdstr := ""
	title := ""
	categories := ""
	tags := ""
	description := ""
	editpath := tmpPath[len(tmpPath)-2] + ".md"

	headerCount := 0
	if FileExist(fileName) {
		fi, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		defer fi.Close()

		br := bufio.NewReader(fi)

		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}

			if headerCount > 1 { //get main body
				mdstr = mdstr + string(a) + "\n"
			}

			if strings.Index(string(a), "---") > -1 { //count header info
				headerCount++
			}

			if headerCount < 2 { //get title info
				if strings.Index(string(a), "title: \"") > -1 {
					title = strings.Replace(string(a), "title: ", "", 1)
					title = strings.Replace(title, "\"", "", -1)
				}

				if strings.Index(string(a), "description: \"") > -1 {
					description = strings.Replace(string(a), "description: ", "", 1)
					description = strings.Replace(description, "\"", "", -1)
				}

				if strings.Index(string(a), "categories: [") > -1 {
					categories = strings.Replace(string(a), "categories: ", "", 1)
					categories = strings.Replace(categories, "[\"", "", -1)
					categories = strings.Replace(categories, "\"", "", -1)
					categories = strings.Replace(categories, "]", "", -1)

				}

				if strings.Index(string(a), "tags: [") > -1 {
					tags = strings.Replace(string(a), "tags: ", "", 1)
					tags = strings.Replace(tags, "[\"", "", -1)
					tags = strings.Replace(tags, "\"", "", -1)
					tags = strings.Replace(tags, "]", "", -1)
				}

			}

			//fmt.Println(string(a))
		}
	} else {
		fmt.Println(fileName)
	}

	fmt.Println(fileName)
	myPage := PageBlog{Account: globalAccount.Address, PageTitle: title, PageDescription: description, PageContent: template.HTML(mdstr), PageTags: tags, PageCategories: categories, EditPath: editpath}
	ctx.ViewData("", myPage)
	ctx.View("haeme_editblog.php")
}

func iBlogUploadFile(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	//filename := ctx.FormValue("filename")
	//fmt.Println("\n\nfile:" + filename + "\n\n")
	file, info, err := ctx.FormFile("editormd-image-file")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}

	defer file.Close()
	fname := info.Filename
	fmt.Println(fname)
	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile("./uploads/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer out.Close()

	io.Copy(out, file)

	//urlString := NodeConfig.LocalWeb + "/uploads/" + fname
	//url, err := url.Parse(urlString)
	//myapi, err := cmdenv.GetApi(myenv, myreq)
	//enc, err := cmdenv.GetCidEncoder(myreq)
	if err != nil {
		//return err
		fmt.Println("Enc failed")
	}
	//fileExt :=fname strings.ToLower(path.Ext(fname))
	/*
		opts := []options.UnixfsAddOption{

			options.Unixfs.CidVersion(1),
			options.Unixfs.RawLeaves(true),
			options.Unixfs.Nocopy(false),
		}
		filepost := files.NewWebFile(url)

		path, err := myapi.Unixfs().Add(myreq.Context, filepost, opts...)
		if err != nil {
			//return err
			fmt.Println("Post file failed")
		} else {
			//fmt.Println("Posted file" + fname + enc.Encode(path.Cid()))
		}
	*/
	uploadedImageValue := `{
  "success": 1,
"message" : "` + fname + `", 
  "url": "/ipfs/` + "" + `"  //////////////////////////////////////////////path
}`

	fmt.Println(uploadedImageValue)
	ctx.Writef(uploadedImageValue)
	//remove the temp uploaded file
	if ostype == "windows" {
		err = os.Remove(".\\uploads\\" + fname)
		if err != nil {
			fmt.Println("Delete uplaod file failed.")
		}
	} else {
		err = os.Remove("./uploads/" + fname)
		if err != nil {
			fmt.Println("Delete uplaod file failed.")
		}
	}
}

func iBuildSite(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}

	configHugo()

	if ostype == "windows" {
		fileExec := "..\\..\\..\\bin\\hugo.exe"
		c := fileExec + " --theme=aeknow"

		cmd := exec.Command("cmd", "/c", c)
		cmd.Dir = ".\\data\\site\\" + globalAccount.Address
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		addstr := string(out)
		fmt.Println(string(out))
		fileExec = ".\\bin\\ipfs.exe"
		c = "set IPFS_PATH=.\\data\\site\\" + globalAccount.Address + "\\repo\\&& " + fileExec + " add -r .\\data\\site\\" + globalAccount.Address + "\\public"
		fmt.Println(c)
		cmd = exec.Command("cmd", "/c", c)
		out, _ = cmd.Output()

		addstr = string(out)
		strArrayNew := strings.Split(addstr, "\n")
		//fmt.Println(strArrayNew[len(strArrayNew)-2], len(strArrayNew))
		strArrayNew = strings.Split(strArrayNew[len(strArrayNew)-2], " ")
		lastIPFS = strArrayNew[len(strArrayNew)-2]

		//Add last IPFS to the page link.
		targetFile := ".\\data\\site\\" + globalAccount.Address + "\\lastIPFS"
		err = ioutil.WriteFile(targetFile, []byte(lastIPFS), 0644)
		if err != nil {
			panic(err)
		}

		//publish IPNS
		c = "set IPFS_PATH=.\\data\\site\\" + globalAccount.Address + "\\repo\\&& " + fileExec + " name publish " + lastIPFS
		cmd = exec.Command("cmd", "/c", c)
		out, err = cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))

		//publish to pubsub
		updateMSG := globalAccount.Address + ":IPFS:" + lastIPFS + ":IPNS:" + MyIPFSConfig.Identity.PeerID
		updateMSG = updateMSG + sigMSG(updateMSG)
		updateMSG = base64.StdEncoding.EncodeToString([]byte(updateMSG))
		PubMSGTo(updateMSG, "update")
		ctx.HTML(lastIPFS + " have been successfully published to: " + string(out) + "<br /> <br /><a href=" + NodeConfig.IPFSNode + "/ipns/" + MyIPFSConfig.Identity.PeerID + ">My IPNS address: </a><br /><br />" + NodeConfig.IPFSNode + "/ipns/" + MyIPFSConfig.Identity.PeerID)

	} else {
		fileExec := "../../../bin/hugo"
		c := fileExec + " --theme=aeknow"
		fmt.Println(c)
		cmd := exec.Command("sh", "-c", c)
		cmd.Dir = "./data/site/" + globalAccount.Address
		//cmd.Run()
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		addstr := string(out)
		fmt.Println(string(out))

		fileExec = "./bin/ipfs"
		c = "export IPFS_PATH=./data/site/" + globalAccount.Address + "/repo && " + fileExec + " add -r ./data/site/" + globalAccount.Address + "/public"
		fmt.Println(c)
		//c := "/home/ae/dev/go/go-ipfs/cmd/ipfs/ipfs add -r /home/ae/dev/go/go-ipfs/cmd/ipfs/data/site/default"
		cmd = exec.Command("sh", "-c", c)
		out, _ = cmd.Output()
		fmt.Println(string(out))
		addstr = string(out)
		strArrayNew := strings.Split(addstr, "\n")
		//fmt.Println(strArrayNew[len(strArrayNew)-2], len(strArrayNew))
		strArrayNew = strings.Split(strArrayNew[len(strArrayNew)-2], " ")
		lastIPFS = strArrayNew[len(strArrayNew)-2]

		//Add last IPFS to the page link.
		targetFile := "./data/site/" + globalAccount.Address + "/lastIPFS"
		err = ioutil.WriteFile(targetFile, []byte(lastIPFS), 0644)
		if err != nil {
			panic(err)
		}

		//publish IPNS
		c = "export IPFS_PATH=./data/site/" + globalAccount.Address + "/repo && " + fileExec + " name publish " + lastIPFS
		fmt.Println(c)
		cmd = exec.Command("sh", "-c", c)
		out, err = cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
		//publish to pubsub
		updateMSG := globalAccount.Address + ":IPFS:" + lastIPFS + ":IPNS:" + MyIPFSConfig.Identity.PeerID
		updateMSG = updateMSG + sigMSG(updateMSG)
		updateMSG = base64.StdEncoding.EncodeToString([]byte(updateMSG))
		PubMSGTo(updateMSG, "update")

		ctx.HTML(lastIPFS + " have been successfully published to: " + string(out) + "<br /> <br /><a href=" + NodeConfig.IPFSNode + "/ipns/" + MyIPFSConfig.Identity.PeerID + ">My IPNS address: </a><br /><br />" + NodeConfig.IPFSNode + "/ipns/" + MyIPFSConfig.Identity.PeerID)
	}

	//DONE:Different users' site and Windows
}
func iUpdateStatic(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}

	if ostype == "windows" {

	} else {
		fileExec, _ := exec.LookPath(os.Args[0])
		fileExec = "./ipfs"
		c := fileExec + " add -r ./data/site/" + globalAccount.Address + "/public"
		fmt.Println(c)
		//c := "/home/ae/dev/go/go-ipfs/cmd/ipfs/ipfs add -r /home/ae/dev/go/go-ipfs/cmd/ipfs/data/site/default"
		cmd := exec.Command("sh", "-c", c)
		out, _ := cmd.Output()
		//fmt.Println(string(out))
		addstr := string(out)
		strArrayNew := strings.Split(addstr, "\n")
		//fmt.Println(strArrayNew[len(strArrayNew)-2], len(strArrayNew))
		strArrayNew = strings.Split(strArrayNew[len(strArrayNew)-2], " ")
		hashstr := strArrayNew[len(strArrayNew)-2]

		//c = "/home/ae/dev/go/go-ipfs/cmd/ipfs/ipfs name publish " + hashstr
		c = fileExec + " name publish " + hashstr
		cmd = exec.Command("sh", "-c", c)
		out, _ = cmd.Output()
		fmt.Println(string(out))
		ctx.HTML(hashstr + "have been successfully published to: " + string(out))
	}

}

func iGoAENS(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	aensname := ctx.URLParam("aensname")
	refresh := ctx.URLParam("refresh")
	gohome := ctx.URLParam("gohome")

	//Go home firstly
	if gohome == "gohome" {
		//ctx.Redirect(NodeConfig.IPFSNode + "/ipns/" + MyIPFSConfig.Identity.PeerID)
		ctx.Redirect(NodeConfig.IPFSNode + "/ipfs/" + lastIPFS)
		return
	}

	//Do normal AENS resolve
	myurl := NodeConfig.PublicNode + "/v2/names/" + aensname

	str := httpGet(myurl)
	fmt.Println(myurl)

	var s AENSInfo
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}

	var myPagedata PageUpdateAENS

	myPagedata.NameID = s.ID
	myPagedata.NameTTL = s.TTL
	myPagedata.NameJson = template.HTML(str)
	myPagedata.AENSName = aensname
	myPagedata.Account = globalAccount.Address

	myPointers := s.Pointers

	var i int

	for i = 0; i < len(myPointers); i++ {
		if myPointers[i].Key == "account_pubkey" {
			myPagedata.AEAddress = myPointers[i].ID
		}
		if myPointers[i].ID == "ak_ipfsD1iUfRLdnJjQMEczjSzzphPbNnSQudnqUAe1vPJetmMK9" {
			myPagedata.IPFSAddress = myPointers[i].Key
		}
		if myPointers[i].ID == "ak_ipnsoMiJmYq1joKGXFtLRDrSJ3mUjapNB7gcPud7mmpVUXssM" {
			myPagedata.IPNSAddress = myPointers[i].Key
		}
		if myPointers[i].Key == "contract_pubkey" {
			myPagedata.ContractAddress = myPointers[i].ID
		}
		if myPointers[i].Key == "oracle_pubkey" {
			myPagedata.OracleAddress = myPointers[i].ID
		}
		if myPointers[i].ID == "ak_btcqM2NycfJaeLYhYY9uPGKj98iVkwL9VLw7ZP5WzzWHHj2sP" {
			myPagedata.BTCAddress = myPointers[i].Key
		}
		if myPointers[i].ID == "ak_ethe795mCkWMAkguuc3ay9k2JSMikZ61L6VfEMDrujEwCiaiB" {
			myPagedata.ETHAddress = myPointers[i].Key
		}

		if myPointers[i].ID == "ak_em3io3Ntov4qJ1y9mDoyQgHTaWBnBZd1CBu7wnH6iyuF5jf5m" {
			myPagedata.EmailAddress = myPointers[i].Key
		}

		if myPointers[i].ID == "ak_webcVNwKZujeYcxDMjAH5ZUPNwCdcFL4QgYD34pFHZi6KEnzS" {
			myPagedata.WebAddress = myPointers[i].Key
		}

	}
	redirecturl := NodeConfig.IPFSNode
	IsRedirect := false

	if myPagedata.IPNSAddress != "" {
		redirecturl = NodeConfig.IPFSNode + "/" + "ipns/" + myPagedata.IPNSAddress
		if refresh == "refresh" {
			t := time.Now()
			timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
			redirecturl = NodeConfig.IPFSNode + "/" + "ipns/" + myPagedata.IPNSAddress + "/?" + timestamp
		}
		IsRedirect = true
	}

	if myPagedata.IPFSAddress != "" {
		redirecturl = NodeConfig.IPFSNode + "/" + "ipfs/" + myPagedata.IPFSAddress
		if refresh == "refresh" {
			t := time.Now()
			timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
			redirecturl = NodeConfig.IPFSNode + "/" + "ipfs/" + myPagedata.IPFSAddress + "/?" + timestamp
		}
		IsRedirect = true
	}

	if IsRedirect {
		fmt.Println(redirecturl)
		ctx.Redirect(redirecturl)
	}

	ctx.HTML("No IPFS or IPNS pointer,AENS info:<br/>" + str)
}

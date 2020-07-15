package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"

	//"net/url"
	"os"
	"os/exec"

	//"path"
	"strconv"
	"strings"
	"time"

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

var MyIPFSConfig IPFSConfig
var lastIPFS string

func getIPFSConfig() IPFSConfig {
	configFilePath := "./data/site/" + globalAccount.Address + "/repo/config"

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
	myPage := PageWallet{Account: globalAccount.Address, PageTitle: "Setting"}
	ctx.ViewData("", myPage)
	ctx.View("haeme_settings.php")
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
	editpath := ctx.FormValue("editpath")
	draft := ctx.FormValue("draft")

	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)

	targetFile := "./data/site/" + globalAccount.Address + "/content/post/" + timestamp + ".md"
	if editpath != "" {
		targetFile = "./data/site/" + globalAccount.Address + "/content/post/" + editpath
	}

	//generate proper tags & categories
	catsstr := "["
	tmpArray := strings.Split(categories, ",")
	var i int
	for i = 0; i < len(tmpArray); i++ {
		catsstr = catsstr + "\"" + tmpArray[i] + "\","
	}
	catsstr = catsstr + "]"
	catsstr = strings.Replace(catsstr, ",]", "]", -1)

	tagstr := "["
	tmpArray = strings.Split(tags, ",")
	for i = 0; i < len(tmpArray); i++ {
		tagstr = tagstr + "\"" + tmpArray[i] + "\","
	}
	tagstr = tagstr + "]"
	tagstr = strings.Replace(tagstr, ",]", "]", -1)

	header := `---
title: "` + title + `"
date: ` + t.UTC().Format(time.UnixDate) + `
categories: ` + catsstr + `
tags: ` + tagstr + `
draft: ` + draft + `
description: "` + description + `"
---`
	//TODO:ADD proper fields to UI
	/*
	   ---
	   # Common-Defined params
	   title: "Example article title"
	   date: "2017-08-21"
	   description: "Example article description"
	   categories:
	     - "Category 1"
	     - "Category 2"
	   tags:
	     - "Test"
	     - "Another test"
	   menu: main # Optional, add page to a menu. Options: main, side, footer

	   # Theme-Defined params
	   thumbnail: "img/placeholder.jpg" # Thumbnail image
	   lead: "Example lead - highlighted near the title" # Lead text
	   comments: false # Enable Disqus comments for specific page
	   authorbox: true # Enable authorbox for specific page
	   pager: true # Enable pager navigation (prev/next) for specific page
	   toc: true # Enable Table of Contents for specific page
	   mathjax: true # Enable MathJax for specific page
	   ---*/
	body := header + "\n" + content
	fmt.Println(targetFile, body)

	err := ioutil.WriteFile(targetFile, []byte(body), 0644)
	if err != nil {
		panic(err)
	} else {
		iBuildSite(ctx)
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

		c = "set IPFS_PATH=.\\data\\site\\" + globalAccount.Address + "\\repo\\&& " + fileExec + " name publish " + lastIPFS
		cmd = exec.Command("cmd", "/c", c)
		out, err = cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
		ctx.HTML(lastIPFS + "have been successfully published to: " + string(out) + "<br /> <br /><a href=" + NodeConfig.IPFSNode + "/ipns/" + MyIPFSConfig.Identity.PeerID + ">My IPNS address: </a><br /><br />" + NodeConfig.IPFSNode + "/ipns/" + MyIPFSConfig.Identity.PeerID)

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

		//c = "/home/ae/dev/go/go-ipfs/cmd/ipfs/ipfs name publish " + hashstr
		//fileExec = "./ipfs"
		c = "export IPFS_PATH=./data/site/" + globalAccount.Address + "/repo && " + fileExec + " name publish " + lastIPFS
		fmt.Println(c)
		cmd = exec.Command("sh", "-c", c)
		out, err = cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
		ctx.HTML(lastIPFS + "have been successfully published to: " + string(out) + "<br /> <br /><a href=" + NodeConfig.IPFSNode + "/ipns/" + MyIPFSConfig.Identity.PeerID + ">My IPNS address: </a><br /><br />" + NodeConfig.IPFSNode + "/ipns/" + MyIPFSConfig.Identity.PeerID)
		//myapi, err := cmdenv.GetApi(myenv, myreq)
		//fmt.Println(MyIPFSConfig.Identity.PeerID)
	}

	//TODO:Different users' site and Windows
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
		ctx.Redirect(NodeConfig.IPFSNode + "/ipns/" + MyIPFSConfig.Identity.PeerID)
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
	if myPagedata.IPFSAddress != "" {
		redirecturl = NodeConfig.IPFSNode + "/" + "ipfs/" + myPagedata.IPFSAddress
		if refresh == "refresh" {
			t := time.Now()
			timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
			redirecturl = NodeConfig.IPFSNode + "/" + "ipfs/" + myPagedata.IPFSAddress + "/?" + timestamp
		}
		IsRedirect = true
	}

	if myPagedata.IPNSAddress != "" {
		redirecturl = NodeConfig.IPFSNode + "/" + "ipns/" + myPagedata.IPNSAddress
		if refresh == "refresh" {
			t := time.Now()
			timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
			redirecturl = NodeConfig.IPFSNode + "/" + "ipns/" + myPagedata.IPNSAddress + "/?" + timestamp
		}
		IsRedirect = true
	}

	if IsRedirect {
		fmt.Println(redirecturl)
		ctx.Redirect(redirecturl)
	}

	ctx.HTML("No IPFS or IPNS pointer,AENS info:<br/>" + str)
}

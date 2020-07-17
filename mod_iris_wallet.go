package main

import (
	//"encoding/json"
	crypto_rand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aeternity/aepp-sdk-go/v7/account"
	aeconfig "github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	qrcode "github.com/skip2/go-qrcode"

	"github.com/jdgcs/ed25519/extra25519"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/nacl/box"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

type PageWallet struct {
	PageId       int
	PageContent  template.HTML
	PageTitle    string
	Account      string
	Balance      string
	Nonce        uint64
	Recipient_id string
	Payload      string
	Amount       string
}

var NodeOnline bool

func iRegisterNew(ctx iris.Context) {
	var myPage PageReg
	myPage.PageTitle = "Registering Page"
	myPage.SubTitle = "Decentralized knowledge system without barrier."
	myPage.Register = "Register"

	//myPage.Lang = getPageString(getPageLang(r))

	ctx.ViewData("", myPage)
	ctx.View("register.php")
}

func iImportUI(ctx iris.Context) {
	ctx.View("import.php")
}
func iExportFromMnemonic(ctx iris.Context) {

	var curve25519Private []byte
	var recipientPrivateKeySlice [64]byte

	entropy := globalAccount.SigningKey
	copy(recipientPrivateKeySlice[0:64], entropy[0:64])
	myrecipientPrivateKey := &recipientPrivateKeySlice

	_, recipientPrivateKey, err := box.GenerateKey(crypto_rand.Reader)
	if err != nil {
		panic(err)
	}

	extra25519.PrivateKeyToCurve25519(recipientPrivateKey, myrecipientPrivateKey)
	curve25519Private = recipientPrivateKey[:]
	fmt.Println(hex.EncodeToString(curve25519Private))
	fromHex, _ := hex.DecodeString(hex.EncodeToString(curve25519Private))
	mnemomic, _ := bip39.NewMnemonic(fromHex)
	fmt.Println(mnemomic)

	/*entropy := globalAccount.SigningKey

	fmt.Println(hex.EncodeToString(globalAccount.SigningKey))

	var recipientPrivateKeySlice [64]byte
	copy(recipientPrivateKeySlice[0:64], entropy[0:64])

	myrecipientPrivateKey := &recipientPrivateKeySlice

	extra25519.PrivateKeyToCurve25519(recipientPrivateKey, myrecipientPrivateKey)

	mypkk := myrecipientPrivateKey
	mnemomic, _ := bip39.NewMnemonic(mypkk)
	fmt.Println(mnemomic)
	//fmt.Println(string(mypkk))
	//ctx.View("import.php")*/

}

func iImportFromMnemonic(ctx iris.Context) {
	password := ctx.FormValue("password")
	password_repeat := ctx.FormValue("password_repeat")
	mnemonic := ctx.FormValue("mnemonic")
	account_index, _ := strconv.ParseInt(ctx.FormValue("account_index"), 10, 32)
	address_index, _ := strconv.ParseInt(ctx.FormValue("address_index"), 10, 32)

	if (password == password_repeat) && len(password) > 1 {
		seed, err := account.ParseMnemonic(mnemonic)
		if err != nil {
			fmt.Println(err)
		}

		// Derive the subaccount m/44'/457'/3'/0'/1'
		key, err := account.DerivePathFromSeed(seed, uint32(account_index), uint32(address_index))
		if err != nil {
			fmt.Println(err)
		}

		// Deriving the aeternity Account from a BIP32 Key is a destructive process
		mykey, err := account.BIP32KeyToAeKey(key)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(mykey.Address)
		accountFileName := "./data/accounts/" + mykey.Address
		if !FileExist(accountFileName) {
			account.StoreToKeyStoreFile(mykey, password, accountFileName)
		} else {
			ctx.HTML("<h1>Account Exist</h1>")
		}
		ctx.Redirect("/")
	} else {
		ctx.HTML("<h1>Passwords must be the same.</h1>")
	}
}

func iDoRegister(ctx iris.Context) {
	password := ctx.FormValue("password")
	password_repeat := ctx.FormValue("password_repeat")
	if (password == password_repeat) && len(password) > 1 {
		acc, _ := account.New()
		accountFileName := "tmpAccount"
		f, _ := account.StoreToKeyStoreFile(acc, password, accountFileName)
		//fmt.Println(acc.Address)
		//fmt.Println(f)
		newFile := "./data/accounts/" + acc.Address
		os.Rename(f, newFile)
		ctx.Redirect("/")
	} else {
		ctx.HTML("<h1>Passwords must be the same.</h1>")
	}
}
func iLogOut(ctx iris.Context) {
	//c := "killall ipfs && ./ipfs daemon"
	//_ = exec.Command("sh", "-c", c)
	//
	globalAccount.Address = ""
	session := sess.Start(ctx)

	//notifyStopping()
	//aerepo.Close() //close the repo

	// Revoke users authentication
	//session.Set("authenticated", false)
	// Or to remove the variable:
	//session.Delete("authenticated")
	// Destroy the whole session:
	session.Destroy()
	ctx.Redirect("/")
	//stop current daemon
	//<-myreq.Context.Done()

	go killIPFS()
}

func killIPFS() {

	if ostype == "windows" {
		c := "TASKKILL /IM ipfs.exe /F"
		fmt.Println(c)
		cmd := exec.Command("cmd", "/c", c)
		output, err := cmd.Output()

		if err != nil {
			fmt.Printf("Execute Shell:%s failed with error:%s", c, err.Error())
			return
		}
		fmt.Printf("Execute Shell:%s finished with output:\n%s", c, string(output))
	} else {
		//kill ipfs firstly
		c := `killall ipfs`
		fmt.Println(c)
		cmd := exec.Command("sh", "-c", c)
		output, err := cmd.Output()

		if err != nil {
			fmt.Printf("Execute Shell:%s failed with error:%s", c, err.Error())
			return
		}
		fmt.Printf("Execute Shell:%s finished with output:\n%s", c, string(output))
	}
}
func iHomePage(ctx iris.Context) {

	needReg := true
	ak := ""
	AccountsLists := ""
	//myLang := getPageString(getPageLang(ctx.Request()))
	//language := ctx.GetLocale().Language()
	//fmt.Println(myLang.Register)

	merr := filepath.Walk("data/accounts/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "ak_") {

			ak = filepath.Base(path)
			if len(ak) > 0 {
				AccountsLists = AccountsLists + "<option>" + ak + "</option>\n"
			}

			needReg = false
		}

		return nil
	})
	//fmt.Println("address:" + globalAccount.Address)
	if len(globalAccount.Address) > 1 {
		if !checkLogin(ctx) {
			return
		}

		needReg = false
		ak := globalAccount.Address

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: MyIPFSConfig.Identity.PeerID}
		ctx.ViewData("", myPage)
		ctx.View("dashboard.php")

		err := qrcode.WriteFile(ak, qrcode.Medium, 256, "./views/qr_ak.png")
		err = qrcode.WriteFile("https://www.aeknow.org/v2/accounts/"+ak, qrcode.Medium, 256, "./views/qr_account.png")
		if err != nil {
			fmt.Println("write error")
		}
	} else {

		var myoption template.HTML
		myoption = template.HTML(AccountsLists)
		myPage := PageLogin{Options: myoption}
		ctx.ViewData("", myPage)
		ctx.View("login.php")
	}

	if merr != nil {
		fmt.Println("error")
	}

	if needReg {

		var myPage PageReg
		myPage.PageTitle = "Registering Page"
		myPage.SubTitle = "Decentralized knowledge system without barrier."
		myPage.Register = "Register"

		//myPage.Lang = getPageString(getPageLang(r))

		//myPage = getPageString(getPageLang(r), "register")
		ctx.ViewData("", myPage)
		ctx.View("register.php")
	}
}

func iCheckLogin(ctx iris.Context) {
	accountname := ctx.FormValue("accountname")
	password := ctx.FormValue("password")
	myAccount, err := account.LoadFromKeyStoreFile("data/accounts/"+accountname, password)
	if err != nil {
		fmt.Println("Could not create myAccount's Account:", err)
		myPage := PageWallet{PageTitle: "Password error:Could not Read Account"}
		ctx.ViewData("", myPage)
		ctx.View("error.php")

	} else { //init the settings
		globalAccount = *myAccount //作为呈现账号
		signAccount = myAccount    //作为签名账号
		IPFS_PATH := "./data/site/" + globalAccount.Address + "/repo/"
		_ = os.Setenv("IPFS_PATH", IPFS_PATH)
		checkHugo()

		checkIPFSRepo(globalAccount.Address)
		// Set user as authenticated
		session := sess.Start(ctx)
		session.Set("authenticated", true)

		NodeConfig = getConfigString() //读取节点设置
		MyIPFSConfig = getIPFSConfig() //读取IPFS节点配置
		MySiteConfig = getSiteConfig() //读取网站设置
		lastIPFS = ""
		configHugo() //登录成功初始化
		go bootIPFS()
		NodeOnline = true

		// Authentication goes here
		// ...

	}

	ctx.Redirect("/")

}

func bootIPFS() { //boot IPFS independently
	if ostype == "windows" {
		fileExec := ".\\bin\\ipfs.exe"
		c := "set IPFS_PATH=data\\site\\" + globalAccount.Address + "\\repo\\&& " + fileExec + " daemon"
		fmt.Println(c)
		cmd := exec.Command("cmd", "/c", c)
		out, _ := cmd.Output()
		fmt.Println(string(out))

	} else {
		fileExec := "./bin/ipfs"

		c := "export IPFS_PATH=./data/site/" + globalAccount.Address + "/repo/&& " + fileExec + " daemon"
		cmd := exec.Command("sh", "-c", c)
		fmt.Println(c)
		out, _ := cmd.Output()
		fmt.Println(string(out))

	}
}

func iMakeTranscaction(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	sender_id := ctx.FormValue("sender_id")
	recipient_id := ctx.FormValue("recipient_id")
	amount := ctx.FormValue("amount")
	payload := ctx.FormValue("payload")
	password := ctx.FormValue("password")

	famount, err := strconv.ParseFloat(amount, 64)
	bigfloatAmount := big.NewFloat(famount)
	imultiple := big.NewFloat(1000000000000000000) //18 dec
	fmyamount := big.NewFloat(1)
	fmyamount.Mul(bigfloatAmount, imultiple)

	myamount := new(big.Int)
	fmyamount.Int(myamount)
	//transfer tokens to .chain name
	if strings.Index(recipient_id, ".chain") > -1 {
		recipient_id = getAccountFromAENS(recipient_id)
	}

	alice, err := account.LoadFromKeyStoreFile("data/accounts/"+sender_id, password)
	if err != nil {

		ak := globalAccount.Address

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Password error:Could not Read Account"}
		ctx.ViewData("", myPage)
		ctx.View("error.php")
		return
	}

	bobAddress := recipient_id

	// create a connection to a node, represented by *Node
	node := naet.NewNode(NodeConfig.PublicNode, false)

	// create the closures that autofill the correct account nonce and transaction TTL
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)

	// create the SpendTransaction

	tx, err := transactions.NewSpendTx(alice.Address, bobAddress, myamount, []byte(payload), ttlnoncer)
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	} else {
		fmt.Println(tx)
	}

	_, myTxhash, _, _, _, err := SignBroadcastWaitTransaction(tx, alice, node, aeconfig.Node.NetworkID, 10)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
		ak := globalAccount.Address

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}
		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	} else {
		ak := globalAccount.Address
		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: myTxhash}
		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	}
}

func iWallet(ctx iris.Context) {
	if !checkLogin(ctx) {
		return
	}
	needReg := true
	ak := ""
	AccountsLists := ""
	recipient_id := ""
	payload := ""
	amountstr := ""

	recipient_id = ctx.URLParam("recipient_id")
	//payloadByte = ctx.URLParam("payload")
	payloadByte, _ := base64.StdEncoding.DecodeString(ctx.URLParam("payload"))
	payload = string(payloadByte)

	amountstr = ctx.URLParam("amount")

	node := naet.NewNode(NodeConfig.PublicNode, false)

	akBalance, err := node.GetAccount(globalAccount.Address)
	var thisamount string
	var myNonce uint64
	if err != nil {
		fmt.Println("Account not exist.")
		thisamount = "0"
		myNonce = 0
	} else {
		bigstr := akBalance.Balance.String()
		myBalance := ToBigFloat(bigstr)
		imultiple := big.NewFloat(0.000000000000000001) //18 dec
		thisamount = new(big.Float).Mul(myBalance, imultiple).String()
		myNonce = *akBalance.Nonce

	}

	merr := filepath.Walk("data/accounts/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "ak_") {

			ak = filepath.Base(path)
			if len(ak) > 0 {
				AccountsLists = AccountsLists + "<option>" + ak + "</option>\n"
			}

			needReg = false
		}
		//fmt.Println(path)
		return nil
	})
	//fmt.Println("address:" + globalAccount.Address)
	if len(globalAccount.Address) > 1 {
		needReg = false
		ak := globalAccount.Address

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Wallet", Balance: thisamount, Nonce: myNonce, Recipient_id: recipient_id, Amount: amountstr, Payload: payload}
		ctx.ViewData("", myPage)
		ctx.View("wallet.php")

		err := qrcode.WriteFile(ak, qrcode.Medium, 256, "./views/qr_ak.png")
		err = qrcode.WriteFile("https://www.aeknow.org/v2/accounts/"+ak, qrcode.Medium, 256, "./views/qr_account.png")
		if err != nil {
			fmt.Println("write error")
		}
	} else {

		var myoption template.HTML
		myoption = template.HTML(AccountsLists)
		myPage := PageLogin{Options: myoption}
		ctx.ViewData("", myPage)
		ctx.View("login.php")
	}

	if merr != nil {
		fmt.Println("error")
	}

	if needReg {

		var myPage PageReg
		myPage.PageTitle = "Registering Page"
		myPage.SubTitle = "Decentralized knowledge system without barrier."
		myPage.Register = "Register"

		myPage.Lang = getPageString(getPageLang(ctx.Request()))

		ctx.ViewData("", myPage)
		ctx.View("register.php")
	}
}
func checkHugo() {
	if ostype == "windows" {
		hugoDir := ".\\data\\site\\" + globalAccount.Address
		if !FileExist(hugoDir) {
			fileExec := "..\\..\\bin\\hugo.exe"
			c := fileExec + " new site " + globalAccount.Address
			fmt.Println(c)
			cmd := exec.Command("cmd", "/c", c)
			cmd.Dir = ".\\data\\site"
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}
			//mkdir post for hugo
			c = "md post"
			cmd = exec.Command("cmd", "/c", c)
			cmd.Dir = ".\\data\\site\\" + globalAccount.Address + "\\content"
			out, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}

			//mkdir and cp theme
			c = "md data\\site\\" + globalAccount.Address + "\\themes\\aeknow"
			fmt.Println(c)
			cmd = exec.Command("cmd", "/c", c)
			out, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}

			c = "xcopy /e /r /y data\\themes\\aeknow  data\\site\\" + globalAccount.Address + "\\themes\\aeknow"
			fmt.Println(c)
			cmd = exec.Command("cmd", "/c", c)
			out, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(out))
			//cp default site config
			c = "copy data\\site.json  data\\site\\" + globalAccount.Address + "\\"
			fmt.Println(c)
			cmd = exec.Command("cmd", "/c", c)
			out, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(out))
		}
	} else {
		hugoDir := "./data/site/" + globalAccount.Address
		if !FileExist(hugoDir) {
			fileExec := "../../bin/hugo"
			c := fileExec + " new site " + globalAccount.Address
			fmt.Println(c)
			cmd := exec.Command("sh", "-c", c)
			cmd.Dir = "./data/site"
			//cmd.Run()
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}
			//mkdir post for hugo
			c = "mkdir post"
			cmd = exec.Command("sh", "-c", c)
			cmd.Dir = "./data/site/" + globalAccount.Address + "/content"
			out, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}
			//cp theme
			c = "cp ./data/themes/aeknow/ -r ./data/site/" + globalAccount.Address + "/themes/"
			fmt.Println(c)
			cmd = exec.Command("sh", "-c", c)
			out, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(out))

			//cp default site config
			c = "cp ./data/site.json ./data/site/" + globalAccount.Address + "/"
			fmt.Println(c)
			cmd = exec.Command("sh", "-c", c)
			out, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(out))
		}
	}
}

func configHugo() {
	if ostype == "windows" {
		//init theme config file
		themeConfigFile := readFileStr(".\\data\\themes\\config.toml")
		targetFile := ".\\data\\site\\" + globalAccount.Address + "\\config.toml"
		err := ioutil.WriteFile(targetFile, []byte(themeConfigFile), 0644)
		if err != nil {
			panic(err)
		}

		fmt.Println(targetFile + "...done.")

		//config search page
		srcFile := ".\\data\\search.html"
		targetFile = ".\\data\\site\\" + globalAccount.Address + "\\content\\search.html"
		if contents, err := ioutil.ReadFile(srcFile); err == nil {
			MyContents := strings.Replace(string(contents), "{{.PeerID}}", MyIPFSConfig.Identity.PeerID, -1)
			err := ioutil.WriteFile(targetFile, []byte(MyContents), 0644)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println("Read search page failed")
		}

	} else {
		//init theme config file
		themeConfigFile := readFileStr("./data/themes/config.toml")
		targetFile := "./data/site/" + globalAccount.Address + "/config.toml"
		err := ioutil.WriteFile(targetFile, []byte(themeConfigFile), 0644)
		if err != nil {
			panic(err)
		}

		fmt.Println(targetFile + "...done.")
		//TODO:	1.mkdir post;2.copy themes;3.init config files;4.init ipns node info;5.add search and remove about page;6.add default help link
		//Done:
		//addstr := string(out)

		//config search page
		srcFile := "./data/search.html"
		targetFile = "./data/site/" + globalAccount.Address + "/content/search.html"
		if contents, err := ioutil.ReadFile(srcFile); err == nil {
			MyContents := strings.Replace(string(contents), "{{.PeerID}}", MyIPFSConfig.Identity.PeerID, -1)
			err := ioutil.WriteFile(targetFile, []byte(MyContents), 0644)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println("Read search page failed")
		}

	}
}

func readFileStr(fileName string) string {
	//TODONE: how to config the site
	if contents, err := ioutil.ReadFile(fileName); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		//fmt.Println(MyIPFSConfig.Identity.PeerID)
		MyContents := strings.Replace(string(contents), "{{.SiteTitle}}", MySiteConfig.Title, -1)
		MyContents = strings.Replace(MyContents, "{{.Author}}", MySiteConfig.Author, -1)
		MyContents = strings.Replace(MyContents, "{{.AuthorDescription}}", MySiteConfig.AuthorDescription, -1)
		MyContents = strings.Replace(MyContents, "{{.Subtitle}}", MySiteConfig.Subtitle, -1)
		MyContents = strings.Replace(MyContents, "{{.SiteDescription}}", MySiteConfig.Description, -1)

		return strings.Replace(MyContents, "{{.Baseurl}}", NodeConfig.IPFSNode+"/ipns/"+MyIPFSConfig.Identity.PeerID+"/", -1)
	}
	return ""
}

//TODO: Post article hash automatically

//Simple version login check for local user
func checkLogin(ctx iris.Context) bool {
	if len(globalAccount.Address) > 1 {
		return true
	}

	return false
}

func checkIPFSRepo(RepoName string) {
	IPFSCheck := "./data/site/" + RepoName + "/repo/version"

	if !FileExist(IPFSCheck) {
		if ostype == "windows" {
			IPFS_PATH := "data\\site\\" + RepoName + "\\repo"
			c := "mkdir " + IPFS_PATH + " && set IPFS_PATH=" + IPFS_PATH + "\\&& bin\\ipfs.exe init &&copy data\\swarm.key " + IPFS_PATH
			fmt.Println(c)
			cmd := exec.Command("cmd", "/c", c)
			out, _ := cmd.Output()

			fmt.Println(string(out))
		} else {
			IPFS_PATH := "./data/site/" + RepoName + "/repo"
			c := "mkdir " + IPFS_PATH + "&& export IPFS_PATH=" + IPFS_PATH + "/&& ./bin/ipfs init && cp ./data/swarm.key " + IPFS_PATH
			fmt.Println(c)
			cmd := exec.Command("sh", "-c", c)
			out, _ := cmd.Output()

			fmt.Println(string(out))

		}
	}

}

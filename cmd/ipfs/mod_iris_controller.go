package main

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

type clientPage struct {
	Title string
	Host  string
}

func myiris() {
	app := iris.New()
	fmt.Println("Web UI is booting...")
	app.HandleDir("/views", "./views")
	app.HandleDir("/uploads", "./uploads")
	app.RegisterView(iris.HTML("./views", ".php"))
	//System
	app.Get("/", iHomePage)
	app.Post("/register", iDoRegister)
	app.Get("/registernew", iRegisterNew)
	app.Post("/login", iCheckLogin)
	app.Get("/logout", iLogOut)
	//import
	app.Get("/import", iImportUI)
	app.Post("/doimport", iImportFromMnemonic)
	//export
	app.Get("/export", iExportFromMnemonic)

	//Wallet
	app.Get("/wallet", iWallet)
	app.Post("/transaction", iMakeTranscaction)

	//Haeme
	app.Get("/haeme", iHaeme)
	app.Get("/updatestatic", iUpdateStatic)
	app.Get("/blog", iBlog)
	app.Get("/newblog", iNewBlog)
	app.Post("/uploadblogimage", iBlogUploadFile)
	app.Post("/saveblog", iSaveBlog)
	app.Get("/buildblog", iBuildSite)
	app.Get("/editpage", iEditBlog)
	app.Get("/delpage", iDelBlog)
	app.Get("/setsite", iSetSite)

	app.Get("/goaens", iGoAENS)

	//AENS
	app.Get("/aens", getAENS)
	app.Get("/aensbidding", getAENSBidding)
	app.Post("/queryaens", iQueryAENS)
	app.Post("/regaens", iDoRegAENS)
	app.Post("/bidaens", iDoBidAENS)
	app.Get("/transfername", iTransferAENS)
	app.Post("/dotransferaens", iDoTransferAENS)
	app.Get("/updatename", iUpdateAENS)
	app.Post("/updatenamepointer", iDoUpdateAENS)
	app.Post("/expertupdatenamepointer", iExpertDoUpdateAENS)

	//Contracts
	app.Get("/contracts", iContractsHome)
	app.Get("/deploy", iDeployContractUI)
	app.Post("/dodeploy", iDoDeployContract)
	app.Get("/call", iCallContractUI)
	app.Post("/docall", iDoCallContract)
	app.Get("/decodecall", iDecodeContractCall)
	app.Post("/dodecode", iDoDecodeContractCall)
	//aex-9 tokens
	app.Get("/deploytoken", iDeployTokenUI)
	app.Post("/dodeploytoken", iDoDeployToken)
	//aex-9 token
	app.Get("/viewtoken", iToken)
	app.Get("/token", getToken)
	app.Post("/transfertoken", iTokenTransfer)

	//Handle chat files
	app.Post("/uploadimage", iChatUploadImage)
	app.Post("/uploadfile", iChatUploadFile)
	//Chat
	app.Get("/chat", iChatUI)
	app.Get("/getchatdata", iGetChatData)

	ws := websocket.New(websocket.DefaultGorillaUpgrader, websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			fmt.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())
			//msg.To = globalAccount.Address
			handleChatMsg(msg, nsConn)

			return nil
		},
	})

	ws.OnConnect = func(c *websocket.Conn) error {
		fmt.Printf("[%s] Connected to server!", c.ID())
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		fmt.Printf("[%s] Disconnected from server", c.ID())
	}

	app.Get("/websocket", websocket.Handler(ws))

	app.Get("/chattest", func(ctx iris.Context) {
		ctx.View("client.php", clientPage{"Client Page", "localhost:8888"})
	})

	//handle proxy ipfs content for editor.md
	app.Get("/ipfs/{anythingparameter:path}", func(ctx iris.Context) {
		paramValue := ctx.Params().Get("anythingparameter")
		ipfsUrl := NodeConfig.IPFSNode + "/ipfs/" + paramValue
		resp, err := http.Get(ipfsUrl)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		//分片逐步写入
		buf := make([]byte, 4096)
		for {
			n, err := resp.Body.Read(buf)
			ctx.Write(buf[:n])
			if err != nil {
				break
			}

		}
	})

	//handle proxy ipns content for editor.md
	app.Get("/ipns/{anythingparameter:path}", func(ctx iris.Context) {
		paramValue := ctx.Params().Get("anythingparameter")
		ipnsUrl := NodeConfig.IPFSNode + "/ipns/" + paramValue
		//fmt.Println("ipnsurl:" + ipnsUrl)
		resp, err := http.Get(ipnsUrl)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		//分片逐步写入
		buf := make([]byte, 4096)
		for {
			n, err := resp.Body.Read(buf)
			ctx.Write(buf[:n])
			if err != nil {
				break
			}

		}
	})

	//test functions for ipfs
	app.Get("/ipks/{anythingparameter:path}", func(ctx iris.Context) {
		paramValue := ctx.Params().Get("anythingparameter")
		ipfsUrl := NodeConfig.IPFSNode + "/ipfs/" + paramValue
		fmt.Println("http.Get =>", ipfsUrl)

		resp, err := http.Get(ipfsUrl)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		//分片逐步写入
		buf := make([]byte, 4096)
		for {
			n, err := resp.Body.Read(buf)
			if err != nil {
				break
			}
			ctx.Write(buf[:n])
		}

		//ctx.Exec("GET", ipfsUrl)
		//fmt.Println("Got?")

	})

	app.Run(iris.Addr("localhost:8888"))
}

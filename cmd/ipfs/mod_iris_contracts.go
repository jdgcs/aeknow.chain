package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"strconv"
	"strings"

	//"github.com/aeternity/aepp-sdk-go/v7/account"
	aeconfig "github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
	"github.com/kataras/iris/v12"

	//"github.com/kataras/iris/v12/sessions"
	qrcode "github.com/skip2/go-qrcode"
)

type TokenInfo struct {
	Tokenname string
	Decimal   int64
	Contract  string
	Balance   string
}

type TokenSlice struct {
	Tokens []TokenInfo
}

type CallResutInfo struct {
	Caller_id    string
	Caller_nonce uint64
	Contract_id  string
	Gas_price    uint64
	Gas_used     uint64
	Height       uint64
	Log          []string
	Return_type  string
	Return_value string
}

type CallResutSlice struct {
	Call_info CallResutInfo
}

type PageDeployContract struct {
	Options   template.HTML
	Account   string
	PageTitle string
}

var ostype = runtime.GOOS

//decode the cb_strings
func iDoDecodeContractCall(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	contract_name := ctx.FormValue("contract_name")
	callfunc := ctx.FormValue("callfunc")
	call_result := strings.Trim(ctx.FormValue("call_result"), "\n")

	if callfunc == "" {
		callfunc = "transfer"
	}

	callStr := strings.Replace(call_result, `"`, `\"`, -1)
	callData := getCallResult(callStr, contract_name, callfunc)

	var myoption template.HTML
	myoption = template.HTML(callData)
	myPage := PageDeployContract{Options: myoption, Account: globalAccount.Address}
	ctx.ViewData("", myPage)
	ctx.View("contract_decoded.php")
}

func iDecodeContractCall(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	contract_name := ""
	ContractsLists := ""
	filepath.Walk("./contracts/decode/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".aes") {

			contract_name = filepath.Base(path)
			if len(contract_name) > 0 {
				ContractsLists = ContractsLists + "<option>" + contract_name + "</option>\n"
			}

		}

		return nil
	})

	var myoption template.HTML
	myoption = template.HTML(ContractsLists)
	myPage := PageDeployContract{Options: myoption, Account: globalAccount.Address}
	ctx.ViewData("", myPage)
	ctx.View("contract_decode.php")
}

//deploy AEX-9 token UI
func iDeployTokenUI(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	myPage := PageWallet{PageId: 23, Account: globalAccount.Address, PageTitle: "AEX-9 Token"}
	ctx.ViewData("", myPage)
	ctx.View("token_create.php")

}

//do deploy AEX-9 token
func iDoDeployToken(ctx iris.Context) {
	// NewContractCreateTx(ownerID string, bytecode string, vmVersion, abiVersion uint16, deposit, amount, gasLimit, gasPrice *big.Int, callData string, ttlnoncer TTLNoncer) (tx *ContractCreateTx, err error)

	name := ctx.FormValue("name")
	symbol := ctx.FormValue("symbol")
	decimals := ctx.FormValue("decimals")
	total_supply := ctx.FormValue("total_supply")
	contract_name := ctx.FormValue("contract_name")

	decimals_int, _ := strconv.Atoi(decimals)
	decimals_long := "000000000000000000000000000000"
	total_supply = total_supply + decimals_long[1:decimals_int]

	//callData := getCallData("init(\""+name+"\","+decimals+",\""+symbol+"\","+total_supply+")", contract_name)
	//callStr := "init(\\\"" + name + "\\\"," + decimals + ",\\\"" + symbol + "\\\",Some(" + total_supply + "))"
	callStr := `init("` + name + `",` + decimals + `,"` + symbol + `",Some(` + total_supply + `))`
	//fmt.Println(callStr)

	callData := getCallData(callStr, contract_name)
	vmVersion := uint16(5)
	abiVersion := uint16(3)
	deposit := big.NewInt(0)
	amount := big.NewInt(0)
	gasLimit := big.NewInt(10000)
	gasPrice := big.NewInt(1000000000)

	byteCode := getByteCode(contract_name)
	node := naet.NewNode(NodeConfig.PublicNode, false)
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)
	ownerID := globalAccount.Address

	tx, err := transactions.NewContractCreateTx(ownerID, byteCode, vmVersion, abiVersion, deposit, amount, gasLimit, gasPrice, callData, ttlnoncer)
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	} else {
		//fmt.Println(tx)
	}

	_, myTxhash, _, _, _, err := SignBroadcastWaitTransaction(tx, signAccount, node, aeconfig.Node.NetworkID, 10)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
		ak := globalAccount.Address

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}
		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	} else {
		ak := globalAccount.Address
		//TODO:get contract id and redirect to another submit token to browser page
		//contract_id := getContractIDFromHash(myTxhash)
		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: myTxhash}
		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	}
}

func iCallContractUI(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	ContractsLists := ""
	contract_name := ""
	//myLang := getPageString(getPageLang(ctx.Request()))
	//language := ctx.GetLocale().Language()
	//fmt.Println(myLang.Register)

	filepath.Walk("./contracts/call/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".aes") {

			contract_name = filepath.Base(path)
			if len(contract_name) > 0 {
				ContractsLists = ContractsLists + "<option>" + contract_name + "</option>\n"
			}

		}

		return nil
	})

	var myoption template.HTML
	myoption = template.HTML(ContractsLists)
	myPage := PageDeployContract{Options: myoption, Account: globalAccount.Address}
	ctx.ViewData("", myPage)
	ctx.View("contract_call.php")

}

func iDoCallContract(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	callfunc := ctx.FormValue("callfunc")
	contract_id := ctx.FormValue("contract_id")
	contract_name := ctx.FormValue("contract_name")

	node := naet.NewNode(NodeConfig.PublicNode, false)
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)
	callerID := globalAccount.Address

	abiVersion := uint16(3)
	amount := big.NewInt(0)
	gasLimit := big.NewInt(10000)
	gasPrice := big.NewInt(1000000000)
	callStr := strings.Replace(callfunc, `"`, `\"`, -1)
	callData := getCallData(callStr, contract_name)

	tx, err := transactions.NewContractCallTx(callerID, contract_id, amount, gasLimit, gasPrice, abiVersion, callData, ttlnoncer)
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	} else {
		//fmt.Println(tx)
	}

	_, myTxhash, _, _, _, err := SignBroadcastWaitTransaction(tx, signAccount, node, aeconfig.Node.NetworkID, 10)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
		ak := globalAccount.Address

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}
		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	} else {
		//TODO:return call result
		call_result_url := NodeConfig.PublicNode + "/v2/transactions/" + myTxhash + "/info"
		fmt.Println(call_result_url)
		txinfo := httpGet(call_result_url)
		//fmt.Println(txinfo)
		var s CallResutSlice
		err = json.Unmarshal([]byte(txinfo), &s)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(s.Call_info.Return_value)
		braIndex := strings.Index(callfunc, "(")

		CallResutStr := getCallResult(s.Call_info.Return_value, contract_name, Substr(callfunc, 0, braIndex))
		fmt.Println(CallResutStr)

		var myoption template.HTML
		myoption = template.HTML(CallResutStr)
		myPage := PageWallet{PageId: 23, Account: globalAccount.Address, PageTitle: myTxhash, PageContent: myoption}
		ctx.ViewData("", myPage)

		ctx.View("transaction.php")
	}
}

//deploy any contracts UI
func iDeployContractUI(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	ContractsLists := ""
	contract_name := ""
	//myLang := getPageString(getPageLang(ctx.Request()))
	//language := ctx.GetLocale().Language()
	//fmt.Println(myLang.Register)

	filepath.Walk("./contracts/deploy/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".aes") {

			contract_name = filepath.Base(path)
			if len(contract_name) > 0 {
				ContractsLists = ContractsLists + "<option>" + contract_name + "</option>\n"
			}

		}

		return nil
	})
	var myoption template.HTML
	myoption = template.HTML(ContractsLists)
	myPage := PageDeployContract{Options: myoption, Account: globalAccount.Address}
	ctx.ViewData("", myPage)
	ctx.View("contract_deploy.php")
}

//deploy any contracts
func iDoDeployContract(ctx iris.Context) {
	// NewContractCreateTx(ownerID string, bytecode string, vmVersion, abiVersion uint16, deposit, amount, gasLimit, gasPrice *big.Int, callData string, ttlnoncer TTLNoncer) (tx *ContractCreateTx, err error)
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	contract_name := ctx.FormValue("contract_name")
	init := ctx.FormValue("init")

	callStr := strings.Replace(init, `"`, `\"`, -1)

	fmt.Println(callStr)

	callData := getCallData(callStr, contract_name)
	vmVersion := uint16(5)
	abiVersion := uint16(3)
	deposit := big.NewInt(0)
	amount := big.NewInt(0)
	gasLimit := big.NewInt(10000)
	gasPrice := big.NewInt(1000000000)

	byteCode := getByteCode(contract_name)
	node := naet.NewNode(NodeConfig.PublicNode, false)
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)
	ownerID := globalAccount.Address

	tx, err := transactions.NewContractCreateTx(ownerID, byteCode, vmVersion, abiVersion, deposit, amount, gasLimit, gasPrice, callData, ttlnoncer)
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	} else {
		//fmt.Println(tx)
	}

	_, myTxhash, _, _, _, err := SignBroadcastWaitTransaction(tx, signAccount, node, aeconfig.Node.NetworkID, 10)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
		ak := globalAccount.Address

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}
		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	} else {
		ak := globalAccount.Address
		//TODO:get contract id and redirect to another submit token to browser page
		//contract_id := getContractIDFromHash(myTxhash)
		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: myTxhash}
		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	}
}

//build transfering token transaction and post it
func iTokenTransfer(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	//sender_id := ctx.FormValue("sender_id")
	recipient_id := ctx.FormValue("recipient_id")
	//transferamount := ctx.FormValue("amount")
	amountstr := ctx.FormValue("amount")
	contractID := ctx.FormValue("contractID")
	//password := ctx.FormValue("password")

	//convert transfer amout to bigint string
	famount, err := strconv.ParseFloat(amountstr, 64)
	bigfloatAmount := big.NewFloat(famount)
	imultiple := big.NewFloat(1000000000000000000) //18 dec
	fmyamount := big.NewFloat(1)
	fmyamount.Mul(bigfloatAmount, imultiple)
	myamount := new(big.Int)
	fmyamount.Int(myamount)

	transferamount := myamount.String()

	node := naet.NewNode(NodeConfig.PublicNode, false)
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)
	ownerID := globalAccount.Address
	//bytecode := ""
	//contractID = "ct_M9yohHgcLjhpp1Z8SaA1UTmRMQzR4FWjJHajGga8KBoZTEPwC"
	//vmVersion := uint16(5)
	abiVersion := uint16(3)
	//deposit := big.NewInt(0)
	amount := big.NewInt(0)
	gasLimit := big.NewInt(10000)
	gasPrice := big.NewInt(1000000000)
	//callData := "cb_KxGEoV2hK58AoMLlAP6SFrYeiuRrxi5A5rNjruumGuhbIsuZStUbvgZYbyS7wfSV"
	callData := getCallData("transfer("+recipient_id+","+transferamount+")", "aex9.aes")
	//callData := getCallData("meta_info()")

	//NewContractCallTx(callerID string, contractID string, amount, gasLimit, gasPrice *big.Int, abiVersion uint16, callData string, ttlnoncer TTLNoncer) (tx *ContractCallTx, err error) {

	tx, err := transactions.NewContractCallTx(ownerID, contractID, amount, gasLimit, gasPrice, abiVersion, callData, ttlnoncer)
	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	} else {
		//fmt.Println(tx)
	}

	_, myTxhash, _, _, _, err := SignBroadcastWaitTransaction(tx, signAccount, node, aeconfig.Node.NetworkID, 10)
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

//get the deployed contract's hash by txhash
//func getContractIDFromHash(myTxhash string) string {
//myurl := NodeConfig.PublicNode + "/v2/transactions/" + myTxhash + "/info"
//str := httpGet(myurl)
//}

//get decoded call result
func getCallResult(callStr, callContract, callfunc string) string {
	if ostype == "windows" {
		c := "bin\\sophia\\erts\\bin\\escript.exe bin\\sophia\\aesophia_cli  contracts\\decode\\" + callContract + " -b fate --call_result " + callStr + " --call_result_fun " + callfunc
		cmd := exec.Command("cmd", "/c", c)
		fmt.Println(c)
		out, _ := cmd.Output()
		callData := strings.Trim(strings.Replace(string(out), "Calldata:", "", 1), "\n")
		fmt.Println(callData)
		return callData
	} else {
		c := "./bin/sophia/erts/bin/escript ./bin/sophia/aesophia_cli  ./contracts/decode/" + callContract + " -b fate --call_result " + callStr + " --call_result_fun " + callfunc
		cmd := exec.Command("sh", "-c", c)
		fmt.Println(c)
		out, _ := cmd.Output()
		callData := strings.Trim(strings.Replace(string(out), "Calldata:", "", 1), "\n")
		fmt.Println(callData)
		return callData
	}

	//cmd := exec.Command("sh", "-c", c)
	return ""
}

//get call data
func getCallData(callStr, callContract string) string {

	if ostype == "windows" {
		c := "bin\\sophia\\erts\\bin\\escript.exe bin\\sophia\\aesophia_cli --create_calldata contracts\\deploy\\" + callContract + " --call " + callStr
		cmd := exec.Command("cmd", "/c", c)
		fmt.Println(c)
		out, _ := cmd.Output()
		callData := strings.Trim(strings.Replace(string(out), "Calldata:", "", 1), "\n")
		fmt.Println("Exec result:" + string(out))
		fmt.Println(callData)
		return callData
	} else {
		c := "./bin/sophia/erts/bin/escript ./bin/sophia/aesophia_cli --create_calldata ./contracts/deploy/" + callContract + " --call \"" + callStr + "\""
		cmd := exec.Command("sh", "-c", c)
		out, _ := cmd.Output()
		callData := strings.Trim(strings.Replace(string(out), "Calldata:", "", 1), "\n")
		fmt.Println(callData)
		return callData
	}

	//cmd := exec.Command("sh", "-c", c)
	return ""

}

//compie bytecode of the contract
func getByteCode(callContract string) string {
	if ostype == "windows" {
		c := "bin\\sophia\\erts\\bin\\escript.exe bin\\sophia\\aesophia_cli contracts\\deploy\\" + callContract
		cmd := exec.Command("cmd", "/c", c)
		out, _ := cmd.Output()
		outStr := strings.Trim(strings.Replace(string(out), "Bytecode:", "", 1), "\n")
		fmt.Println(outStr)

		return outStr
	} else {
		c := "./bin/sophia/erts/bin/escript ./bin/sophia/aesophia_cli ./contracts/deploy/" + callContract
		cmd := exec.Command("sh", "-c", c)
		out, _ := cmd.Output()
		outStr := strings.Trim(strings.Replace(string(out), "Bytecode:", "", 1), "\n")
		fmt.Println(outStr)

		return outStr
	}

	return ""

}

//Token main page
func getToken(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
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

	//fmt.Println(myNonce)
	//fmt.Println(thisamount)

	myurl := NodeConfig.APINode + "/api/token/" + globalAccount.Address
	str := httpGet(myurl)
	var s TokenSlice
	err = json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}

	var i int
	myNames := ""
	for i = 0; i < len(s.Tokens); i++ {
		//fmt.Println(s.Names[i].Aensname)
		bigstr := s.Tokens[i].Balance
		myBalance := ToBigFloat(bigstr)
		imultiple := big.NewFloat(0.000000000000000001) //18 dec
		//thisamount = new(big.Float).Mul(myBalance, imultiple).String()
		thistokenamount := fmt.Sprintf("%.2f", new(big.Float).Mul(myBalance, imultiple))

		myNames = myNames + `<tr>                    
                    <td><a href=` + NodeConfig.APINode + `/token/view/` + s.Tokens[i].Tokenname + ` target=_blank>` + s.Tokens[i].Tokenname + `</a></td>  
                    <td><a href="">` + strconv.FormatInt(s.Tokens[i].Decimal, 10) + `</a></td>    
                    <td><a href="">` + thistokenamount + `</a></td>              
                    <td align="center">
                      <div class="btn-group">
						  <a href=/viewtoken?contractid=` + s.Tokens[i].Contract + `><button type="button" class="btn btn-success">Transfer</button></a> &nbsp;
						 
						</div>
                    </td>
                  </tr>`
	}

	myAENSLists := template.HTML(myNames)
	aensCount := len(s.Tokens)
	myPage := PageAENS{PageId: aensCount, Account: globalAccount.Address, PageTitle: "Wallet", Balance: thisamount, Nonce: myNonce, PageContent: myAENSLists}

	ctx.ViewData("", myPage)
	ctx.View("tokenhome.php")
	//TODO:get the balance of each account quickly.
}

//Token management page
func iToken(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	contractid := ctx.URLParam("contractid")

	needReg := true
	ak := ""
	AccountsLists := ""
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

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: contractid, Balance: thisamount, Nonce: myNonce}
		ctx.ViewData("", myPage)
		ctx.View("token.php")

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

func iContractsHome(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	myPage := PageWallet{Account: globalAccount.Address}
	ctx.ViewData("", myPage)
	ctx.View("contract_home.php")
}

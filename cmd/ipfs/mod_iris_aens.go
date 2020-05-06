package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/big"
	"strconv"
	"strings"

	aeconfig "github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
	"github.com/kataras/iris/v12"
)

type PageUpdateAENS struct {
	AENSName        string
	NameID          string
	PointerJson     template.HTML
	NameJson        template.HTML
	NameTTL         uint64
	Account         string
	Balance         string
	Nonce           uint64
	AEAddress       string
	IPFSAddress     string
	IPNSAddress     string
	ContractAddress string
	OracleAddress   string
	BTCAddress      string
	ETHAddress      string
	EmailAddress    string
	WebAddress      string
}

func iExpertDoUpdateAENS(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	myPointerJson := ctx.FormValue("pointerjson")
	aensname := ctx.FormValue("aensname")
	ak := globalAccount.Address

	//fmt.Println(myPointerJson)

	var s []*transactions.NamePointer

	err := json.Unmarshal([]byte(myPointerJson), &s)
	if err != nil {
		fmt.Println(err)
	}

	node := naet.NewNode(NodeConfig.PublicNode, false)
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)

	tx, err := transactions.NewNameUpdateTx(globalAccount.Address, aensname, s, 50000, ttlnoncer)

	//fmt.Println(tx)

	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	}

	_, myTxhash, _, _, _, err := SignBroadcastWaitTransaction(tx, signAccount, node, aeconfig.Node.NetworkID, 10)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}
		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	} else {
		//fmt.Println("TxHash:" + myTxhash)

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: myTxhash}
		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	}
}

func iDoUpdateAENS(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	aensname := ctx.FormValue("aensname")
	aeaddress := ctx.FormValue("aeaddress")
	ipfsaddress := ctx.FormValue("ipfsaddress")
	ipnsaddress := ctx.FormValue("ipnsaddress")

	contractaddress := ctx.FormValue("contractaddress")
	oracleaddress := ctx.FormValue("oracleaddress")
	btcaddress := ctx.FormValue("btcaddress")
	ethaddress := ctx.FormValue("ethaddress")
	emailaddress := ctx.FormValue("emailaddress")
	webaddress := ctx.FormValue("webaddress")

	ak := globalAccount.Address

	//fmt.Println(aensname)

	myPointerJson := "["
	if strings.TrimSpace(aeaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"` + aeaddress + `","key":"account_pubkey"},`
	} else {
		myPointerJson = myPointerJson + `{"id":"` + globalAccount.Address + `","key":"account_pubkey"},`
	}

	if strings.TrimSpace(ipfsaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"ak_ipfsD1iUfRLdnJjQMEczjSzzphPbNnSQudnqUAe1vPJetmMK9","key":"` + ipfsaddress + `"},`
	} else {
		if lastIPFS != "" {
			myPointerJson = myPointerJson + `{"id":"ak_ipfsD1iUfRLdnJjQMEczjSzzphPbNnSQudnqUAe1vPJetmMK9","key":"` + lastIPFS + `"},`
		}
	}

	if strings.TrimSpace(ipnsaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"ak_ipnsoMiJmYq1joKGXFtLRDrSJ3mUjapNB7gcPud7mmpVUXssM","key":"` + ipnsaddress + `"},`
	} else {
		if MyIPFSConfig.Identity.PeerID != "" {
			myPointerJson = myPointerJson + `{"id":"ak_ipnsoMiJmYq1joKGXFtLRDrSJ3mUjapNB7gcPud7mmpVUXssM","key":"` + MyIPFSConfig.Identity.PeerID + `"},`
		}
	}

	if strings.TrimSpace(contractaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"` + contractaddress + `","key":"contract_pubkey"},`
	}
	if strings.TrimSpace(oracleaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"` + oracleaddress + `","key":"oracle_pubkey"},`
	}
	if strings.TrimSpace(btcaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"ak_btcqM2NycfJaeLYhYY9uPGKj98iVkwL9VLw7ZP5WzzWHHj2sP","key":"` + btcaddress + `"},`
	}
	if strings.TrimSpace(ethaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"ak_ethe795mCkWMAkguuc3ay9k2JSMikZ61L6VfEMDrujEwCiaiB","key":"` + ethaddress + `"},`
	}

	if strings.TrimSpace(emailaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"ak_em3io3Ntov4qJ1y9mDoyQgHTaWBnBZd1CBu7wnH6iyuF5jf5m","key":"` + emailaddress + `"},`
	}
	if strings.TrimSpace(webaddress) != "" {
		myPointerJson = myPointerJson + `{"id":"ak_webcVNwKZujeYcxDMjAH5ZUPNwCdcFL4QgYD34pFHZi6KEnzS","key":"` + webaddress + `"},`
	}

	myPointerJson = myPointerJson + "]"
	myPointerJson = strings.Replace(myPointerJson, ",]", "]", -1)
	//fmt.Println(myPointerJson)
	var s []*transactions.NamePointer

	err := json.Unmarshal([]byte(myPointerJson), &s)
	if err != nil {
		fmt.Println(err)
	}

	node := naet.NewNode(NodeConfig.PublicNode, false)
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)
	//p := s
	tx, err := transactions.NewNameUpdateTx(globalAccount.Address, aensname, s, 50000, ttlnoncer)

	//fmt.Println(tx)

	if err != nil {
		fmt.Println("Could not create the SpendTx:", err)
	}

	_, myTxhash, _, _, _, err := SignBroadcastWaitTransaction(tx, signAccount, node, aeconfig.Node.NetworkID, 10)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}
		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	} else {

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: myTxhash}

		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	}
}

func iUpdateAENS(ctx iris.Context) {
	aensname := ctx.URLParam("aensname")
	myurl := NodeConfig.PublicNode + "/v2/names/" + aensname
	str := httpGet(myurl)
	//fmt.Println(myurl)

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

	ctx.ViewData("", myPagedata)
	ctx.View("aens_update.php")
}

func iDoTransferAENS(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	aensname := ctx.FormValue("aensname")
	toaddress := ctx.FormValue("toaddress")
	ak := globalAccount.Address
	node := naet.NewNode(NodeConfig.PublicNode, false)
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)

	tx, err := transactions.NewNameTransferTx(globalAccount.Address, aensname, toaddress, ttlnoncer)

	_, myTxhash, _, _, _, err := SignBroadcastWaitTransaction(tx, signAccount, node, aeconfig.Node.NetworkID, 10)
	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}

		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	} else {

		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: myTxhash}

		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	}
}

func iTransferAENS(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	aensname := ctx.URLParam("aensname")
	myPage := PageAENS{PageContent: template.HTML(aensname)}
	ctx.ViewData("", myPage)
	ctx.View("aens_transfer.php")
}

func iDoBidAENS(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	aensname := ctx.FormValue("aensname")
	aensprice := ctx.FormValue("aensprice")
	recommendprice := ctx.FormValue("recommendprice")

	var myprice float64

	if strings.TrimSpace(aensprice) == "" {
		myprice, _ = strconv.ParseFloat(recommendprice, 64)
	} else {
		myprice, _ = strconv.ParseFloat(aensprice, 64)
	}

	bigfloatAmount := big.NewFloat(myprice)
	imultiple := big.NewFloat(1000000000000000000) //18 dec
	fmyamount := big.NewFloat(1)
	fmyamount.Mul(bigfloatAmount, imultiple)

	myamount := new(big.Int)
	fmyamount.Int(myamount)

	node := naet.NewNode(NodeConfig.PublicNode, false)
	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)
	nameSalt := big.NewInt(0)
	tx, _ := transactions.NewNameClaimTx(globalAccount.Address, aensname, nameSalt, myamount, ttlnoncer)

	_, claimTxhash, _, _, _, _ := SignBroadcastWaitTransaction(tx, signAccount, node, aeconfig.Node.NetworkID, 10)

	hashInfo := "Bidding Tx hash is <a href=" + NodeConfig.APINode + "/block/transaction/" + claimTxhash + ">" + claimTxhash + "</a><br /><br /><br />"
	ak := globalAccount.Address
	myPage := PageWallet{PageId: 23, Account: ak, PageTitle: claimTxhash, PageContent: template.HTML(hashInfo)}

	ctx.ViewData("", myPage)
	ctx.View("transaction.php")
}
func iDoRegAENS(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	aensname := ctx.FormValue("aensname")
	aensprice := ctx.FormValue("aensprice")

	var myprice float64

	if strings.TrimSpace(aensprice) == "" {
		myprice = calcAENSFee(aensname)
	} else {
		myprice, _ = strconv.ParseFloat(aensprice, 64)
	}

	bigfloatAmount := big.NewFloat(myprice)
	imultiple := big.NewFloat(1000000000000000000) //18 dec
	fmyamount := big.NewFloat(1)
	fmyamount.Mul(bigfloatAmount, imultiple)

	myamount := new(big.Int) //regfee
	fmyamount.Int(myamount)

	node := naet.NewNode(NodeConfig.PublicNode, false)

	_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)
	tx_pre, nameSalt, err := transactions.NewNamePreclaimTx(globalAccount.Address, aensname, ttlnoncer)
	fmt.Println("Ready to Precliam the AENS Name " + aensname)
	_, preClaimTxhash, _, _, _, err := SignBroadcastWaitTransaction(tx_pre, signAccount, node, aeconfig.Node.NetworkID, 10)
	hashInfo := "Preclaim Tx hash is <a href=" + NodeConfig.APINode + "/block/transaction/" + preClaimTxhash + ">" + preClaimTxhash + "</a><br /><br /><br />"
	fmt.Println(aensname + " was precliamed.\nReady to claim the AENS Name " + aensname + "\n\nPlease waiting for several minutes for 1 block...")

	tx, err := transactions.NewNameClaimTx(globalAccount.Address, aensname, nameSalt, myamount, ttlnoncer)

	fmt.Println(aensname + " was registered successfully.")

	_, claimTxhash, _, err := SignBroadcastTransaction(tx, signAccount, node, aeconfig.Node.NetworkID)
	hashInfo = hashInfo + "Claim Tx hash is <a href=" + NodeConfig.APINode + "/block/transaction/" + claimTxhash + ">" + claimTxhash + "</a><br />Please <b>WAIT ~1 block </b> time and check this transaction."

	if err != nil {
		fmt.Println("SignBroadcastTransaction failed with:", err)
		ak := globalAccount.Address
		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: "Failed"}

		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	} else {
		ak := globalAccount.Address
		myPage := PageWallet{PageId: 23, Account: ak, PageTitle: claimTxhash, PageContent: template.HTML(hashInfo)}

		ctx.ViewData("", myPage)
		ctx.View("transaction.php")
	}
}

func iQueryAENS(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	aensname := ctx.FormValue("aensname")
	if strings.Index(aensname, ".chain") == -1 {
		aensname = aensname + ".chain"
	}
	myurl := NodeConfig.APINode + "/api/aensquery/" + aensname
	//fmt.Println(myurl)
	str := httpGet(myurl)
	status := ""
	if strings.Index(str, "NONE") > -1 {
		regFee := calcAENSFeeStr(aensname)
		status = `  <div class="box"><div class="col-md-9"> <div class="box-footer">
		<form action="/regaens" method="post">
<input type="hidden" name="aensname" value="` + aensname + `">
                    <div class="input-group">
                      ` + aensname + `:<input type="text" name="aensprice" placeholder="Default price: ` + regFee + ` AE" class="form-control">
                      <br/><br/><br/>                    
                            <button type="submit" class="btn btn-warning btn-flat">Register & Wait several minutes</button>
                         
                    </div>
                  </form></div></div></div>`
		//status = aensname + "=>" + str
	}
	if strings.Index(str, "DONE") > -1 {
		s := strings.Split(str, ":")
		myBalance := ToBigFloat(s[1])
		imultiple := big.NewFloat(0.000000000000000001) //18 dec
		thisamount := new(big.Float).Mul(myBalance, imultiple).String()
		status = "<a href=" + NodeConfig.APINode + "/" + aensname + ">" + aensname + "</a>" + "=>Registed.\n<br /><br />Price=>" + thisamount + " AE\n<br /><br /><a href=" + NodeConfig.APINode + "/aens/viewbids/" + aensname + ">Check bidding details</a>"

		//status = aensname + "=>" + thisamount
	}

	if strings.Index(str, "BIDDING") > -1 {
		s := strings.Split(str, ":")
		myBalance := ToBigFloat(s[1])
		imultiple := big.NewFloat(0.000000000000000001) //18 dec
		thisamount := new(big.Float).Mul(myBalance, imultiple).String()

		imultiple = big.NewFloat(0.00000000000000000105) //18 dec
		recommendamount := new(big.Float).Mul(myBalance, imultiple).String()

		status = `  <div class="box"><div class="col-md-9"> <div class="box-footer">
		<form action="/bidaens" method="post">
<input type="hidden" name="aensname" value="` + aensname + `">
<input type="hidden" name="recommendprice" value="` + recommendamount + `">
                    <div class="input-group">
                      ` + aensname + "=>Last bidding price:" + thisamount + ` AE(<a href=` + NodeConfig.APINode + `/aens/viewbids/` + aensname + ` target=_blank>View bidding details</a>) <input type="text" name="aensprice" placeholder="Recommend bidding price: ` + recommendamount + ` AE" class="form-control">
                      <br/><br/><br/>                    
                            <button type="submit" class="btn btn-warning btn-flat">Bidding with my price</button>
                         
                    </div>
                  </form></div></div></div>`

	}

	queryResults := template.HTML(status)

	myPage := PageAENS{PageContent: queryResults}

	ctx.ViewData("", myPage)
	ctx.View("aens_query.php")
}

func getAENSBidding(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	ak := globalAccount.Address
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

	myurl := NodeConfig.APINode + "/api/aensbidding/" + ak
	str := httpGet(myurl)
	//fmt.Println(myurl)
	//fmt.Println(str)
	var s NameBidSlice
	err = json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}

	var i int
	myNames := ""
	for i = 0; i < len(s.Names); i++ {
		//fmt.Println(s.Names[i].Aensname)
		bidBalance := ToBigFloat(s.Names[i].Lastprice)
		myimultiple := big.NewFloat(0.000000000000000001) //18 dec
		mylastprice := new(big.Float).Mul(bidBalance, myimultiple).String()
		if s.Names[i].Lastbidder == globalAccount.Address {
			myNames = myNames + `<tr>                    
                    <td><a href=` + NodeConfig.APINode + `/aens/viewbids/` + s.Names[i].Aensname + ` target=_blank>` + s.Names[i].Aensname + `</a></td>                    
                    <td><a href=` + NodeConfig.APINode + `/address/wallet/` + s.Names[i].Lastbidder + ` target=_blank>` + s.Names[i].Lastbidder + `</a></td>
 					<td>` + mylastprice + ` AE </td>
                    <td align="center">
                      <div class="btn-group">
                     <button type="button" class="btn btn-success">OK</button> &nbsp;
						
						</div>
                    </td>
                  </tr>`
		} else {
			myNames = myNames + `<tr>                    
                    <td><a href=` + NodeConfig.APINode + `/aens/viewbids/` + s.Names[i].Aensname + ` target=_blank>` + s.Names[i].Aensname + `</a></td>                    
                    <td><a href=` + NodeConfig.APINode + `/address/wallet/` + s.Names[i].Lastbidder + ` target=_blank>` + s.Names[i].Lastbidder + `</a></td>
 					<td>` + mylastprice + ` AE </td>
                    <td align="center">
                      <div class="btn-group">
					<form action="/queryaens" method="post">
                      <input type="hidden" name="aensname" value="` + s.Names[i].Aensname + `" class="form-control"><button type="submit"  type="button" class="btn btn-warning">Add price</button></a> &nbsp;
					</form>
						</div>
                    </td>
                  </tr>`
		}
	}

	myAENSLists := template.HTML(myNames)
	aensCount := len(s.Names)
	myPage := PageAENS{PageId: aensCount, Account: ak, PageTitle: "Wallet", Balance: thisamount, Nonce: myNonce, PageContent: myAENSLists}
	ctx.ViewData("", myPage)
	ctx.View("aens_bidding.php")
}

func getAENS(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	node := naet.NewNode(NodeConfig.PublicNode, false)

	akBalance, err := node.GetAccount(globalAccount.Address)
	topHeight, _ := node.GetHeight()
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

	myurl := NodeConfig.APINode + "/api/aens/" + globalAccount.Address
	str := httpGet(myurl)
	var s NameSlice
	err = json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}

	var i int
	myNames := ""
	for i = 0; i < len(s.Names); i++ {
		//fmt.Println(s.Names[i].Aensname)
		num1 := float64(s.Names[i].Expire_height - int64(topHeight))
		num2 := float64(480)

		d1 := num1 / num2
		days := fmt.Sprintf("%.1f", d1)
		myNames = myNames + `<tr>
                    <td><a href="">` + strconv.FormatInt(s.Names[i].Expire_height, 10) + `</a>(~` + days + ` days)</td>
                    <td><a href=` + NodeConfig.APINode + `/` + s.Names[i].Aensname + ` target=_blank>` + s.Names[i].Aensname + `</a></td>                    
                    <td align="center">
                      <div class="btn-group">
						  <a href=/updatename?aensname=` + s.Names[i].Aensname + `><button type="button" class="btn btn-success">Update</button></a> &nbsp;
						  <a href=/transfername?aensname=` + s.Names[i].Aensname + `><button type="button" class="btn btn-info pull-right">Transfer</button></a>
						</div>
                    </td>
                  </tr>`
	}

	myAENSLists := template.HTML(myNames)
	aensCount := len(s.Names)
	myPage := PageAENS{PageId: aensCount, Account: globalAccount.Address, PageTitle: "Wallet", Balance: thisamount, Nonce: myNonce, PageContent: myAENSLists}

	ctx.ViewData("", myPage)
	ctx.View("aens.php")
}

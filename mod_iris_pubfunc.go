package main

import (
	//"context"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"

	//"sync"
	"time"

	"github.com/aeternity/aepp-sdk-go/v7/account"
	aeconfig "github.com/aeternity/aepp-sdk-go/v7/config"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
)

const (
	dnsResolveTimeout = 10 * time.Second
)

func SmartPrint(i interface{}) {
	var kv = make(map[string]interface{})
	vValue := reflect.ValueOf(i)
	vType := reflect.TypeOf(i)
	for i := 0; i < vValue.NumField(); i++ {
		kv[vType.Field(i).Name] = vValue.Field(i)
	}
	fmt.Println("获取到数据:")
	for k, v := range kv {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
	}
}

func Openhome() error {
	uri := "./jump.html"
	run, ok := aecommands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	fmt.Println("Aeternity is bridged..." + runtime.GOOS + run + " Opening browser...")

	cmd := exec.Command("cmd", "/c start "+uri)
	return cmd.Start()
}

func getConfigString() AeknowConfig {
	configFilePath := "./data/config.json"

	_, err := os.Stat(configFilePath)

	if err != nil {
		configFilePath = "./data/config_default.json"
	}

	JsonParse := NewJsonStruct()
	readConfigfile := AeknowConfig{}
	JsonParse.Load(configFilePath, &readConfigfile)

	return readConfigfile

}

type GetAnythingByNameFunc func(name, key string) (results []string, err error)

// GenerateGetAnythingByName is the underlying implementation of Get*ByName
func GenerateGetAnythingByName(n naet.GetNameEntryByNamer) GetAnythingByNameFunc {
	return func(name string, key string) (results []string, err error) {
		nameEntry, err := n.GetNameEntryByName(name)
		if err != nil {
			return []string{}, err
		}
		for _, p := range nameEntry.Pointers {
			if *p.Key == key {
				results = append(results, *p.ID)
			}
		}
		return results, nil
	}
}

// GetAccountsByName returns any account_pubkey entries that it finds in a
// name's Pointers.
func GetAccountsByName(n GetAnythingByNameFunc, name string) (addresses []string, err error) {
	return n(name, "account_pubkey")
}

// GetOraclesByName returns any oracle_pubkey entries that it finds in a name's
// Pointers.
func GetOraclesByName(n GetAnythingByNameFunc, name string) (oracleIDs []string, err error) {
	return n(name, "oracle_pubkey")
}

// GetContractsByName returns any contract_pubkey entries that it finds in a
// name's Pointers.
func GetContractsByName(n GetAnythingByNameFunc, name string) (contracts []string, err error) {
	return n(name, "contract_pubkey")
}

// GetChannelsByName returns any channel entries that it finds in a name's
// Pointers.
func GetChannelsByName(n GetAnythingByNameFunc, name string) (channels []string, err error) {
	return n(name, "channel")
}

// getTransactionByHashHeighter is used by WaitForTransactionForXBlocks to
// specify that the node/mock node passed in should support
// GetTransactionByHash() and GetHeight()
type getTransactionByHashHeighter interface {
	naet.GetTransactionByHasher
	naet.GetHeighter
}

// ErrWaitTransaction is returned by WaitForTransactionForXBlocks() to let
// callers distinguish between network errors and transaction acceptance errors.
type ErrWaitTransaction struct {
	NetworkErr     bool
	TransactionErr bool
	Err            error
}

func (b ErrWaitTransaction) Error() string {
	var errType string
	if b.TransactionErr {
		errType = "TransactionErr"
	} else {
		errType = "NetworkErr"
	}

	return fmt.Sprintf("%s: %s", errType, b.Err.Error())
}

// WaitForTransactionForXBlocks blocks until a transaction has been mined or X
// blocks have gone by, after which it returns an error. The node polling
// interval can be config.Configured with config.Tuning.ChainPollInterval.
func WaitForTransactionForXBlocks(c getTransactionByHashHeighter, txHash string, x uint64) (blockHeight uint64, blockHash string, wtError error) {
	nodeHeight, err := c.GetHeight()
	if err != nil {
		wtError = ErrWaitTransaction{
			NetworkErr:     true,
			TransactionErr: false,
			Err:            err,
		}
		return
	}
	endHeight := nodeHeight + x
	for nodeHeight <= endHeight {
		nodeHeight, err = c.GetHeight()
		if err != nil {
			wtError = ErrWaitTransaction{
				NetworkErr:     true,
				TransactionErr: false,
				Err:            err,
			}
			return
		}
		tx, err := c.GetTransactionByHash(txHash)
		if err != nil {
			wtError = ErrWaitTransaction{
				NetworkErr:     false,
				TransactionErr: true,
				Err:            err,
			}
			return
		}

		if tx.BlockHeight.LargerThanZero() {
			bh := big.Int(tx.BlockHeight)
			return bh.Uint64(), *tx.BlockHash, nil
		}
		time.Sleep(time.Millisecond * time.Duration(aeconfig.Tuning.ChainPollInterval))
	}
	wtError = ErrWaitTransaction{
		NetworkErr:     false,
		TransactionErr: true,
		Err:            fmt.Errorf("%v blocks have gone by and %v still isn't in a block", x, txHash),
	}
	return
}

// SignBroadcastTransaction signs a transaction and broadcasts it to a node.
func SignBroadcastTransaction(tx transactions.Transaction, signingAccount *account.Account, n naet.PostTransactioner, networkID string) (signedTxStr, hash, signature string, err error) {

	signedTx, hash, signature, err := transactions.SignHashTx(signingAccount, tx, networkID)
	if err != nil {
		return
	}
	fmt.Println(hash)
	signedTxStr, err = transactions.SerializeTx(signedTx)

	//fmt.Println(signedTxStr)
	if err != nil {
		return
	}

	err = n.PostTransaction(signedTxStr, hash)
	if err != nil {
		return
	}
	return
}

type broadcastWaitTransactionNodeCapabilities interface {
	naet.PostTransactioner
	getTransactionByHashHeighter
}

// SignBroadcastWaitTransaction is a convenience function that combines
// SignBroadcastTransaction and WaitForTransactionForXBlocks.
func SignBroadcastWaitTransaction(tx transactions.Transaction, signingAccount *account.Account, n broadcastWaitTransactionNodeCapabilities, networkID string, x uint64) (signedTxStr, hash, signature string, blockHeight uint64, blockHash string, err error) {
	signedTxStr, hash, signature, err = SignBroadcastTransaction(tx, signingAccount, n, networkID)
	fmt.Println(signedTxStr)
	fmt.Println(hash)

	if err != nil {
		return
	}
	blockHeight, blockHash, err = WaitForTransactionForXBlocks(n, hash, x)
	return
}

func httpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(body)
}

func IPFSAPIPost(data string, postfunc string) string {
	request, _ := http.NewRequest("POST", NodeConfig.IPFSAPI+"/api/"+postfunc, strings.NewReader(data))
	request.Header.Add("content-type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("post data error:%v\n", err)
		return "post data error"
	} else {
		//fmt.Println("post a data successful.")
		respBody, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("response data:%v\n", string(respBody))
		return string(respBody)
	}
}
func calcAENSFeeStr(aensname string) string {
	if len(aensname) == 7 {
		return "570.3"
	}
	if len(aensname) == 8 {
		return "352.5"
	}
	if len(aensname) == 9 {
		return "217.9"
	}
	if len(aensname) == 10 {
		return "134.7"
	}
	if len(aensname) == 11 {
		return "83.3"
	}
	if len(aensname) == 12 {
		return "51.5"
	}
	if len(aensname) == 13 {
		return "31.8"
	}
	if len(aensname) == 14 {
		return "19.7"
	}
	if len(aensname) == 15 {
		return "12.2"
	}
	if len(aensname) == 16 {
		return "7.51"
	}
	if len(aensname) == 17 {
		return "4.64"
	}

	return "2.9"
}

func calcAENSFee(aensname string) float64 {
	if len(aensname) == 7 {
		return 570.3
	}
	if len(aensname) == 8 {
		return 352.5
	}
	if len(aensname) == 9 {
		return 217.9
	}
	if len(aensname) == 10 {
		return 134.7
	}
	if len(aensname) == 11 {
		return 83.3
	}
	if len(aensname) == 12 {
		return 51.5
	}
	if len(aensname) == 13 {
		return 31.8
	}
	if len(aensname) == 14 {
		return 19.7
	}
	if len(aensname) == 15 {
		return 12.2
	}
	if len(aensname) == 16 {
		return 7.51
	}
	if len(aensname) == 17 {
		return 4.64
	}

	return 2.9
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
func ToBigFloat(str string) *big.Float {
	f, _, _ := big.ParseFloat(str, 10, 256, big.ToNearestEven)
	return f
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

//start：正数 - 在字符串的指定位置开始,超出字符串长度强制把start变为字符串长度
//       负数 - 在从字符串结尾的指定位置开始
//       0 - 在字符串中的第一个字符处开始
//length:正数 - 从 start 参数所在的位置返回
//       负数 - 从字符串末端返回

func Substr(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	rune_str := []rune(str)
	len_str := len(rune_str)

	if start < 0 {
		start = len_str + start
	}
	if start > len_str {
		start = len_str
	}
	end := start + length
	if end > len_str {
		end = len_str
	}
	if length < 0 {
		end = len_str + length
	}
	if start > end {
		start, end = end, start
	}
	return string(rune_str[start:end])
}

func ConnetDefaultNodes() {
	//set test ipfs nodes for global and CN
	seednode1 := "/ip4/104.156.239.14/tcp/4001/p2p/QmXZUVYs6SNHNJqFNwTKVRsRB2Lom5N2RtfH8T8CT3sFfU"
	seednode2 := "/ip4/111.231.110.42/tcp/4001/p2p/QmXiowBAKzjKXjkRKWJRFZXkS6BsKbYXgXHmoWp4hSSCsD"
	//Do connect once firstly
	time.Sleep(20 * time.Second)
	DoConnect(seednode1)
	DoConnect(seednode2)
	go ReadPubsub("update") //listening update channel
	//Reconnect continuously every 30 secs(?)
	for {
		time.Sleep(30 * time.Second)
		if NodeOnline {
			fmt.Println("Connect to seeds...")
			DoConnect(seednode1)
			DoConnect(seednode2)
		}
	}
}

func DoConnect(addr string) {
	IPFSAPIPost("", "v0/swarm/connect?arg="+addr)
}

func PubMSGTo(msg string, topic string) {
	//"http://127.0.0.1:5001/api/v0/pubsub/pub?arg=<topic>&arg=<data>"
	IPFSAPIPost("", "v0/pubsub/pub?arg="+topic+"&arg="+msg)
}

func sigMSG(msg string) string {
	//mysignAccount :=account.FromHexString(signAccount.Sign())
	signed := base64.StdEncoding.EncodeToString(signAccount.Sign([]byte(msg)))
	fmt.Println(signed)
	return ":SIG:" + signed
}

var (
	httpClient *http.Client
)

func ReadPubsub(topic string) {
	fmt.Println("Start listening..." + topic)
	//curl -X POST "http://127.0.0.1:5001/api/v0/pubsub/sub?arg=<topic>&discover=<value>"
	var endPoint string = "http://127.0.0.1:5001/api/v0/pubsub/sub?arg=" + topic

	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer([]byte("")))
	if err != nil {
		fmt.Println("Error Occured. %+v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	fmt.Println(endPoint)
	// use httpClient to send request
	response, err := httpClient.Do(req)
	if err != nil && response == nil {
		fmt.Println("Error sending request to API endpoint. %+v", err)
	} else {
		// Close the connection to reuse it
		defer response.Body.Close()

		// Let's check if the work actually is done
		// We have seen inconsistencies even when we get 200 OK response
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Couldn't parse response body. %+v", err)
		}

		fmt.Println("Response Body:", string(body))
	}
}

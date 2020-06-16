package main

import (
	//"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/aeternity/aepp-sdk-go/v7/account"
	"github.com/jdgcs/ed25519"

	cmds "github.com/ipfs/go-ipfs-cmds"
	files "github.com/ipfs/go-ipfs-files"
	cmdenv "github.com/ipfs/go-ipfs/core/commands/cmdenv"
	options "github.com/ipfs/interface-go-ipfs-core/options"

	"github.com/kataras/iris/v12"
	iriswebsocket "github.com/kataras/iris/v12/websocket"
	"golang.org/x/net/websocket"
)

type pubsubMessage struct {
	From     []byte   `json:"from,omitempty"`
	Data     []byte   `json:"data,omitempty"`
	Seqno    []byte   `json:"seqno,omitempty"`
	TopicIDs []string `json:"topicIDs,omitempty"`
}

const (
	pubsubDiscoverOptionName = "discover"
)

type Message struct {
	Username string
	Message  string
	Topic    string
}

type User struct {
	Username string
}

type Datas struct {
	Messages []Message
	Users    []User
}

type PageChat struct {
	Account     string
	PageContent template.HTML
	PageTitle   string
	ChatTopic   string
}

type ChatMsg struct {
	Mine SenderMsg
	To   ReceiverMsg
}

type SenderMsg struct {
	Username  string
	Groupname string
	Avatar    string
	Id        string
	Type      string
	Content   string
	Cid       string
	Mine      bool
	Fromid    string
	Timestamp int64
	Name      string
}

type ReceiverMsg struct {
	Username  string
	Groupname string
	Avatar    string
	Id        string
	Type      string
	Content   string
	Cid       string
	Mine      bool
	Fromid    string
	Timestamp uint64
	Sign      string
	Name      string
}

// 全局信息
var datas Datas
var users map[*websocket.Conn]string
var myreq *cmds.Request
var myres cmds.ResponseEmitter
var myenv cmds.Environment
var lastTimestamp int64

//所有聊天channel，默认启动一个
var chatChan = [1000]string{"ae"}

func handleChatMsg(msg iriswebsocket.Message, nsConn *iriswebsocket.NSConn) {
	//
	/*利用Message定义聊天消息的结构，需要包含
	From string
	To string
	Signature string //如果是公开消息
	Msgtype string //0-palin, 1-crypted
	Body string//json string, 包含nonce
	*/
	//SmartPrint(msg)
	topic := "ak_fCCw1JEkvXdztZxk8FRGNAkvmArhVeow89e64yX4AxbCPrVh5"
	msgBody := string(msg.Body)
	if msgBody != "ping" {
		fmt.Println("Ready to decode msg....")
		var s ChatMsg
		err := json.Unmarshal([]byte(msgBody), &s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Msg From " + s.Mine.Id + " to...." + s.To.Id)
		//SmartPrint(s)

		//topic := msg.To
		//if strings.Index(msgBody, "avatar") > 0 {
		if s.Mine.Id == globalAccount.Address {
			myapi, err := cmdenv.GetApi(myenv, myreq)
			if err = myapi.PubSub().Publish(myreq.Context, topic, []byte(msgBody)); err != nil {
				fmt.Println("Publish failed")
				//return err
			}
		} else {
			fmt.Println("Received Msg:" + msgBody)
		}
	} else {
		myapi, err := cmdenv.GetApi(myenv, myreq)
		msgBody = "ping from:" + globalAccount.Address

		if err = myapi.PubSub().Publish(myreq.Context, topic, []byte(msgBody)); err != nil {
			fmt.Println("Online braoadcast failed.")
			//return err
		}
		fmt.Println("Broadcast ping:" + msgBody)
		RecordPingTimestamp()
	}

	if msgBody == globalAccount.Address+" Online" {

		//fmt.Println(lastTimestamp)
		myapi, err := cmdenv.GetApi(myenv, myreq)
		if err = myapi.PubSub().Publish(myreq.Context, topic, []byte(msgBody)); err != nil {
			fmt.Println("Online braoadcast failed.")
			//return err
		}
	}

	nsConn.Conn.Server().Broadcast(nsConn, msg)
}

//记录最后一次ping的时间，便于下次推送
func RecordPingTimestamp() {
	lastTimestamp = time.Now().UnixNano() / 1e6
}

func iChatUploadImage(ctx iris.Context) {
	//filename := ctx.FormValue("filename")
	//fmt.Println("\n\nfile:" + filename + "\n\n")
	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}

	defer file.Close()
	fname := info.Filename
	fmt.Println("\n\nfile:" + fname + "\n\n")
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
	//fmt.Println("uploaded?")

	urlString := NodeConfig.LocalWeb + "/uploads/" + fname
	url, err := url.Parse(urlString)
	myapi, err := cmdenv.GetApi(myenv, myreq)
	enc, err := cmdenv.GetCidEncoder(myreq)
	if err != nil {
		//return err
		fmt.Println("Enc failed")
	}

	opts := []options.UnixfsAddOption{
		options.Unixfs.Pin(true),
		options.Unixfs.CidVersion(1),
		options.Unixfs.RawLeaves(true),
		options.Unixfs.Nocopy(false),
	}

	filepost := files.NewWebFile(url)

	path, err := myapi.Unixfs().Add(myreq.Context, filepost, opts...)
	if err != nil {
		//return err
		fmt.Println("Post file failed", err)
	} else {
		//fmt.Println("Posted file" + fname + enc.Encode(path.Cid()))
	}

	uploadedImageValue := `{
  "code": 0 
  ,"msg": "" 
  ,"data": {
    "src": "` + NodeConfig.IPFSNode + `/ipfs/` + enc.Encode(path.Cid()) + `" 
  }
}`
	ctx.Writef(uploadedImageValue)
}
func iChatUploadFile(ctx iris.Context) {

	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}

	defer file.Close()
	fname := info.Filename

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

	urlString := NodeConfig.LocalWeb + "/uploads/" + fname
	url, err := url.Parse(urlString)
	myapi, err := cmdenv.GetApi(myenv, myreq)
	enc, err := cmdenv.GetCidEncoder(myreq)
	if err != nil {
		//return err
		fmt.Println("Enc failed")
	}

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
		fmt.Println("Posted file" + fname + enc.Encode(path.Cid()))
	}
	uploadedFileValue := `{
  "code": 0 
  ,"msg": "" 
  ,"data": {
    "src": "` + NodeConfig.IPFSNode + `/ipfs/` + enc.Encode(path.Cid()) + `" 
    ,"name": "` + fname + `"
  }
}`
	ctx.Writef(uploadedFileValue)

}

func iGetChatData(ctx iris.Context) {
	action := ctx.URLParam("action")

	if action == "get_user_data" {

	}

	if action == "groupMembers" {

	}

}

func iChatUI(ctx iris.Context) {
	//mytopic := req.FormValue("topic")
	mytopic := ctx.URLParam("topic")
	needGo := true
	lastNull := 0
	for i := range chatChan {
		if chatChan[i] == mytopic {
			fmt.Println("Topic exist:" + mytopic)
			needGo = false
			break
		}

	}

	for i := range chatChan {
		if chatChan[i] == "" {
			lastNull = i
			break
		}
	}

	if needGo {
		chatChan[lastNull] = mytopic
		go toTopic(myreq, myres, myenv, mytopic)
	}

	myPage := PageChat{Account: globalAccount.Address, ChatTopic: mytopic, PageTitle: "Private Data Digging"}
	ctx.ViewData("", myPage)
	ctx.View("chat.php")
}

func toTopic(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment, topic string) error {
	fmt.Println("Subscrbing..." + topic)
	api, err := cmdenv.GetApi(env, req)
	if err != nil {
		fmt.Println("Init failed..." + topic)
		return err
	}

	discover, _ := req.Options[pubsubDiscoverOptionName].(bool)

	var origin = NodeConfig.LocalWeb + "/"
	var url = "ws://127.0.0.1:8888/websocket"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("Websocket connect failed..." + topic)
	} else {
		fmt.Println("Connected..." + topic)
	}

	sub, err := api.PubSub().Subscribe(req.Context, topic, options.PubSub.Discover(discover))
	if err != nil {
		return err
	}
	defer sub.Close()

	if f, ok := res.(http.Flusher); ok {
		f.Flush()
	}

	for {
		msg, err := sub.Next(req.Context)
		if err == io.EOF || err == context.Canceled {
			return nil
		} else if err != nil {
			return err
		}
		//fmt.Println("From:" + string(msg.From()))
		fmt.Println("Msg in:" + string(msg.Data()))

		var s ChatMsg
		err = json.Unmarshal(msg.Data(), &s)
		if err != nil {
			fmt.Println(err)
		}

		if s.To.Id == globalAccount.Address { //私聊
			s.Mine.Mine = false
			s.Mine.Type = "friend"
			s.Mine.Timestamp = time.Now().UnixNano() / 1e6

			//msgStr := string(s.Mine)
			msgTmpStr, err := json.Marshal(s.Mine)
			msgStr := string(msgTmpStr)

			msgStr = strings.Replace(msgStr, "Username", "username", 1)
			msgStr = strings.Replace(msgStr, "Type", "type", 1)
			msgStr = strings.Replace(msgStr, "Avatar", "avatar", 1)
			msgStr = strings.Replace(msgStr, "Id", "id", 1)
			msgStr = strings.Replace(msgStr, "Content", "content", 1)
			msgStr = strings.Replace(msgStr, "Mine", "mine", 1)
			msgStr = strings.Replace(msgStr, "Timestamp", "timestamp", 1)
			msgStr = strings.Replace(msgStr, "Fromid", "fromid", 1)
			msgStr = strings.Replace(msgStr, "Cid", "cid", 1)
			fmt.Printf("Pre Broadcast: %s\n", msgStr)
			//b, err := json.Marshal(s.Mine)
			//b, err := json.Marshal([]byte(msgStr))
			//_, err = ws.Write(msg.Data())

			_, err = ws.Write([]byte(msgStr))
			if err != nil {
				fmt.Println(err)
			} else {
				//fmt.Printf("Broadcast: %s\n", string(msg.Data()))
				fmt.Printf("Broadcast: %s\n", msgStr)
			}

		}

		if s.To.Type == "group" && s.Mine.Id != globalAccount.Address { //群聊基础版本
			s.Mine.Mine = false
			s.Mine.Type = "group"
			s.Mine.Id = s.To.Id
			s.Mine.Name = s.To.Name
			s.Mine.Groupname = s.To.Groupname
			s.Mine.Timestamp = time.Now().UnixNano() / 1e6

			//msgStr := string(s.Mine)
			msgTmpStr, err := json.Marshal(s.Mine)
			msgStr := string(msgTmpStr)

			msgStr = strings.Replace(msgStr, "Username", "username", 1)
			msgStr = strings.Replace(msgStr, "Type", "type", 1)
			msgStr = strings.Replace(msgStr, "Avatar", "avatar", 1)
			msgStr = strings.Replace(msgStr, "Id", "id", 1)
			msgStr = strings.Replace(msgStr, "Content", "content", 1)
			msgStr = strings.Replace(msgStr, "Mine", "mine", 1)
			msgStr = strings.Replace(msgStr, "Timestamp", "timestamp", 1)
			msgStr = strings.Replace(msgStr, "Fromid", "fromid", 1)
			msgStr = strings.Replace(msgStr, "Cid", "cid", 1)
			msgStr = strings.Replace(msgStr, "Group", "group", 1)
			msgStr = strings.Replace(msgStr, "Groupname", "groupname", 1)
			msgStr = strings.Replace(msgStr, "Name", "name", 1)
			//fmt.Printf("Pre Broadcast: %s\n", msgStr)
			//b, err := json.Marshal(s.Mine)
			//b, err := json.Marshal([]byte(msgStr))
			//_, err = ws.Write(msg.Data())

			_, err = ws.Write([]byte(msgStr))
			if err != nil {
				fmt.Println(err)
			} else {
				//fmt.Printf("Broadcast: %s\n", string(msg.Data()))
				fmt.Printf("Group Broadcast: %s\n", msgStr)
			}

		}
		/*if !strings.Contains(string(msg.Data()), "username") {
			msgBody := "{\"username\":\"localakak\",\"message\":\"" + string(msg.Data()) + "\"}"
			_, err = ws.Write([]byte(msgBody))
			fmt.Println("Msg in json:" + msgBody)
			if err != nil {
				//log.Fatal(err)
			}
			fmt.Printf("Send: %s\n", string(msg.Data()))
		}*/
		//fmt.Println("Topics:" + string(msg.Topics()))
		//fmt.Println("Seqno:" + string(msg.Seq()))
		if err := res.Emit(&pubsubMessage{
			Data:     msg.Data(),
			From:     []byte(msg.From()),
			Seqno:    msg.Seq(),
			TopicIDs: msg.Topics(),
		}); err != nil {
			return err
		}
	}
	ws.Close()
	return nil
}

func ChatSocket(ws *websocket.Conn) {
	var message Message
	var data string

	myapi, err := cmdenv.GetApi(myenv, myreq)
	if err != nil {
		//return err
	}

	topic := "ae"

	myAccount, err := account.LoadFromKeyStoreFile("data/config/ak_CdCnkxwJg472TgzHgFUhvNazoraGqLLBWH6wABgHbYPmxfHeh", "aep")

	for {
		// 接收数据
		err := websocket.Message.Receive(ws, &data)
		if err != nil {
			// 移除出错的连接
			delete(users, ws)
			fmt.Println("连接异常")
			break
		}

		// 解析信息
		err = json.Unmarshal([]byte(data), &message)
		if err != nil {
			fmt.Println("解析数据异常")
		}
		topic = message.Topic
		// 添加新用户到map中,已经存在的用户不必添加
		if _, ok := users[ws]; !ok {
			users[ws] = message.Username

			// 添加用户到全局信息
			datas.Users = append(datas.Users, User{Username: message.Username})
		}
		fmt.Println("check data:" + data)
		//time.Sleep(3 * time.Second)
		//if !strings.Contains(data, "localakak") {
		var recipientPrivateKeySlice [64]byte
		var recipientPublicKeySlice [32]byte
		copy(recipientPrivateKeySlice[0:64], myAccount.SigningKey[0:64])
		copy(recipientPublicKeySlice[0:32], myAccount.SigningKey[32:64])
		chatSig := ed25519.Sign(&recipientPrivateKeySlice, []byte(data))

		sum := recipientPrivateKeySlice
		//hashChannel <- sum[:]
		tmpStr := hex.EncodeToString(sum[:])
		fmt.Println("prvk:" + tmpStr)

		sum32 := recipientPublicKeySlice
		//hashChannel <- sum[:]
		tmpStr = hex.EncodeToString(sum32[:])
		fmt.Println("pubk:" + tmpStr)

		fmt.Println("ChatSig:")
		fmt.Println(chatSig)

		if ed25519.Verify(&recipientPublicKeySlice, []byte(data), chatSig) {
			fmt.Println("Signature checked.")
		} else {
			fmt.Println("Signature failed.")
		}

		if err = myapi.PubSub().Publish(myreq.Context, topic, []byte(data)); err != nil {
			fmt.Println("failed???")
			//return err

		}
		//}

		// 添加聊天记录到全局信息
		datas.Messages = append(datas.Messages, message)

		// 通过webSocket将当前信息分发
		for key := range users {
			err := websocket.Message.Send(key, data)
			if err != nil {
				// 移除出错的连接
				delete(users, key)
				fmt.Println("发送出错: " + err.Error())
				break
			}
		}
	}
}

func defaultSub(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
	myreq = req
	myres = res
	myenv = env
	fmt.Println("Subscrbing...")
	api, err := cmdenv.GetApi(env, req)
	if err != nil {
		return err
	}

	topic := "ae"
	discover, _ := req.Options[pubsubDiscoverOptionName].(bool)

	var origin = "http://127.0.0.1:8888/"
	var url = "ws://127.0.0.1:8888/websocket"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		//log.Fatal(err)
	}

	sub, err := api.PubSub().Subscribe(req.Context, topic, options.PubSub.Discover(discover))
	if err != nil {
		return err
	}
	defer sub.Close()

	if f, ok := res.(http.Flusher); ok {
		f.Flush()
	}

	for {
		msg, err := sub.Next(req.Context)
		if err == io.EOF || err == context.Canceled {
			return nil
		} else if err != nil {
			return err
		}
		//fmt.Println("From:" + string(msg.From()))
		fmt.Println("Msg in:" + string(msg.Data()))

		if !strings.Contains(string(msg.Data()), "username") {
			msgBody := "{\"username\":\"localakak\",\"message\":\"" + string(msg.Data()) + "\"}"
			_, err = ws.Write([]byte(msgBody))
			fmt.Println("Msg in json:" + msgBody)
			if err != nil {
				//log.Fatal(err)
			}
			fmt.Printf("Send: %s\n", string(msg.Data()))
		}
		//fmt.Println("Topics:" + string(msg.Topics()))
		//fmt.Println("Seqno:" + string(msg.Seq()))
		if err := res.Emit(&pubsubMessage{
			Data:     msg.Data(),
			From:     []byte(msg.From()),
			Seqno:    msg.Seq(),
			TopicIDs: msg.Topics(),
		}); err != nil {
			return err
		}
	}
	ws.Close()
	return nil
}

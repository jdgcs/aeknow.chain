//getlanguage,read languagefiles
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type langFile struct {
	//register
	Register             string
	Register_description string
}

//return the language string for the page
func getPageString(clientLang string) langFile {
	langFilePath := "./data/lang/" + clientLang + ".json"

	_, err := os.Stat(langFilePath)

	if err != nil {
		langFilePath = "./data/lang/en.json"
	}

	JsonParse := NewJsonStruct()
	readLangfile := langFile{}
	JsonParse.Load(langFilePath, &readLangfile)

	return readLangfile

}

//Get the language of the browser
func getPageLang(r *http.Request) string {
	langStr := r.Header["Accept-Language"]
	browserLang := langStr[0]
	myLang := strings.Split(browserLang, ",")[0]
	retunLang := strings.Replace(myLang, "[", "", -1)
	retunLang = strings.Replace(retunLang, "]", "", -1)
	fmt.Println(retunLang)

	return retunLang
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

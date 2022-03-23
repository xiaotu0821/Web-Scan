package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	funcs "web_scan/Functions"
	Configs "web_scan/config"
)

func init() {
	flag.StringVar(&Configs.UserObject.KeyWord, "find", "", "fuzz to found *.json file")
	flag.StringVar(&Configs.UserObject.JsonFile, "json", "", "it will be used in exploit")
	flag.StringVar(&Configs.UserObject.OriAddr, "url", "", "this url will be attack")
	//flag.StringVar(&Configs.UserObject.Cmd, "cmd", "whoami", "will be execute command")
	flag.StringVar(&Configs.UserObject.File, "file", "", "will be urls list file")
	flag.BoolVar(&Configs.UserObject.AllJson, "alljson", false, "will be all json exec payload")
	flag.BoolVar(&Configs.UserObject.GetTitle, "gettitle", false, "Get url title")
	flag.IntVar(&Configs.UserObject.ThreadNum, "thread", 1, "use threadn number,default number 1")

	flag.Parse()

}
func LoadOneExpJson(filepath string, oneExpjson *ExpJsons) {
	file, err := os.Open(filepath)
	if err != nil {

		fmt.Println("LoadOneExpJson func open filepath err...")
	}
	fileValue, err := ioutil.ReadAll(file)
	if err != nil {

		fmt.Println("LoadOneExpJson func readAll err...")
	}
	err = json.Unmarshal(fileValue, oneExpjson)
	if err != nil {

		fmt.Println("LoadOneExpJson Json file to load failed")
	}
}

type ExpJsons struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Product     string `json:"Product"`
	Author      string `json:"author"`

	Request []struct {
		Method           string            `json:"Method"`
		Header           map[string]string `json:"Header"`
		Uri              string            `json:"Uri"`
		Port             string            `json:"Port"`
		Data             string            `json:"Data"`
		Follow_redirects string            `json:"Follow_redirects"`
		Upload           struct {
			Name     string `json:"Name"`
			FileName string `json:"fileName"`
			FilePath string `json:"FilePath"`
		} `json:"Upload"`
		Response struct {
			Check_Steps string `json:"Check_Steps"`
			Checks      []struct {
				Operation string `json:"Operation"`
				Key       string `json:"Key"`
				Value     string `json:"Value"`
			} `json:"Checks"`
		}
		Next_decide string `json:"Next_decide"`
	} `json:"Request"`
}

func main() {

	// url := "https://www.baidu.com"
	// resp, str := option.GetTitleRequest(url, 5)
	// if str != "" {
	// 	fmt.Println(str)
	// }
	// if resp != nil {
	// 	strs := string(resp.Header.Get(""))

	// 	//fmt.Println(strings.Contains(strs, "baidu"))
	// 	fmt.Println(strs)
	// 	//fmt.Println(strings.Contains(dataString, "BAIDU"))
	//	ceshi()
	funcs.Banner()
	funcs.Choose()

	fmt.Println("程序执行完毕，主程序退出")
	//}

}

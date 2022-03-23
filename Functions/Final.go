package Funcs

import (
	"encoding/json"
	"fmt"
	"os"
	option "web_scan/Functions/options"
	Configs "web_scan/config"
)

var file_success_log string = option.Get_Time() + "success.txt" //成功记录写入这
var file_fail_log string = option.Get_Time() + "fail.txt"       //失败记录写入这

type results struct {
	res bool
	url string
}

type all_result struct {
	res      []bool
	exp_name []string
	url      string
}

func GetAllJson() []Configs.ExpJson {

	FindResltAllJson := []string{}
	FindResltAllJson, _ = option.FindFileAllJson(Configs.ConfigJsonMap.Exploit.Path, FindResltAllJson)
	//所有的json文件

	for _, filename := range FindResltAllJson {
		fmt.Println(filename)
	}
	size := len(FindResltAllJson)
	var AllExpJsonContent []Configs.ExpJson = make([]Configs.ExpJson, size)
	for i := 0; i < size; i++ {
		filevalue := option.LoadExpJsonAll(FindResltAllJson[i]) //获取每个输入json文件的内容

		var expjson Configs.ExpJson
		err := json.Unmarshal(filevalue, &expjson) //将每个json文件内容放到结构体数组中
		AllExpJsonContent[i] = expjson
		//	fmt.Println(AllExpJsonContent[i])  详细每个json的内容
		if err != nil {
			option.MistakPrint("Json file to load failed")
			continue
		}
	}
	return AllExpJsonContent //用于存放 返回的所有json内容
}

func All_url_one_Json(oneExpjson Configs.ExpJson, urls <-chan string, result chan<- results, timeout_count map[string]int) {

	for url := range urls {
		var res results
		res.url = url
		res.res = option.JudgeMent_OneUrl_OneJson(url, oneExpjson, timeout_count)
		result <- res

	}

}

func final_One_url_allJson(timeout_count map[string]int) {

	AllExpJsonContent := GetAllJson()
	size := len(AllExpJsonContent)
	result := false
	//---------------
	fp_succ, err := os.OpenFile(file_success_log, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("open success-file  fail")
	}
	fp_fail, err := os.OpenFile(file_fail_log, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("open fail-file  fail")
	}
	for i := 0; i < size; i++ {

		result = option.JudgeMent_OneUrl_OneJson(Configs.UserObject.OriAddr, AllExpJsonContent[i], timeout_count)
		if result == true {
			fmt.Println("-----------------------------------------------------------------------------------------")
			success := Configs.UserObject.OriAddr + "\t" + AllExpJsonContent[i].Name + "\t" + "Exploit Success !" + "\n" //成功时字符串拼接
			option.SuccessPrint(success)

			fp_succ.WriteString(success)

		} else {
			fmt.Println("-----------------------------------------------------------------------------------------")
			Failed := Configs.UserObject.OriAddr + "\t" + AllExpJsonContent[i].Name + "\t" + "Exploit Failed !" + "\n" //失败时字符串拼接
			option.FailPrint(Failed)
			fp_fail.WriteString(Failed)

		}
	}
	defer fp_succ.Close()
	defer fp_fail.Close()
}

func final_Oneurl_OneJson(timeout_count map[string]int) {
	var oneExpjson Configs.ExpJson
	option.LoadOneExpJson(Configs.UserObject.JsonFile, &oneExpjson)

	status := option.JudgeMent_OneUrl_OneJson(Configs.UserObject.OriAddr, oneExpjson, timeout_count)

	fmt.Print("status=")
	fmt.Println(status)
	if status == true {
		fmt.Println(Configs.UserObject.OriAddr + "\t" + oneExpjson.Name + "\t" + "Exploit Success !" + "\n")

	} else {

		fmt.Println(Configs.UserObject.OriAddr + "\t" + oneExpjson.Name + "\t" + "Exploit Failed !" + "\n")
	}
}

func final_ALLurl_OneJson(timeout_count map[string]int) {

	filename := Configs.UserObject.File

	urllist := option.GetUrlFile(filename)
	size := len(urllist)
	result := make(chan results, size+1)

	//results := make(chan bool, size+1)
	jobs_url := make(chan string, size+1)
	var oneExpjson Configs.ExpJson

	fp_succ, err := os.OpenFile(file_success_log, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	fp_fail, err := os.OpenFile(file_fail_log, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)

	option.LoadOneExpJson(Configs.UserObject.JsonFile, &oneExpjson) //载入一个expjson

	for i := 0; i < Configs.UserObject.ThreadNum; i++ {
		go All_url_one_Json(oneExpjson, jobs_url, result, timeout_count)
	}

	for i := 0; i < size; i++ {
		jobs_url <- urllist[i]
	}

	for i := 0; i < size; i++ {

		res := <-result

		if res.res == true {
			fmt.Println("-----------------------------------------------------------------------------------------")
			success := res.url + "\t" + oneExpjson.Name + "\t" + "Exploit Success !" + "\n" //成功时字符串拼接   //两处的url可能有问题
			option.SuccessPrint(success)

			if err != nil {
				panic(err)
			}
			fp_succ.WriteString(success)

		} else {
			fmt.Println("-----------------------------------------------------------------------------------------")
			Failed := res.url + "\t" + oneExpjson.Name + "\t" + "Exploit Failed !" + "\n" //失败时字符串拼接
			option.FailPrint(Failed)

			fp_fail.WriteString(Failed)
			if err != nil {
				fmt.Println("file write err...:", err)
			}

		}
	}
	defer fp_succ.Close()
	defer fp_fail.Close()
}

//---------------------------------------------------------------------------------

func All_url_ALLjson(urls <-chan string, allresult chan<- all_result, timeout_count map[string]int) {
	tmp_result := false
	AllExpJsonContent := GetAllJson()

	fmt.Println("外层 timeout_count=", timeout_count)
	for url := range urls {
		var res all_result
		res.url = url
		for _, expjsoncontent := range AllExpJsonContent {
			tmp_result = option.JudgeMent_OneUrl_OneJson(url, expjsoncontent, timeout_count)
			res.res = append(res.res, tmp_result)
			res.exp_name = append(res.exp_name, expjsoncontent.Name)
		}
		allresult <- res
	}

}

func final_ALLurl_ALLJson(timeout_count map[string]int) {
	if Configs.UserObject.AllJson == true && Configs.UserObject.File != "" {
		filename := Configs.UserObject.File

		urllist := option.GetUrlFile(filename)
		size := len(urllist)
		allresult := make(chan all_result, size+1)

		//results := make(chan bool, size+1)
		jobs_url := make(chan string, size+1)
		for i := 0; i < Configs.UserObject.ThreadNum; i++ {
			go All_url_ALLjson(jobs_url, allresult, timeout_count)
		}
		for i := 0; i < size; i++ {
			jobs_url <- urllist[i]
		}
		for i := 0; i < size; i++ {
			res := <-allresult
			for key, res_tmp := range res.res {
				if res_tmp == true {
					fmt.Println("-----------------------------------------------------------------------------------------")
					success := res.url + "\t" + res.exp_name[key] + "\t" + "Exploit Success !" + "\n" //成功时字符串拼接   //两处的url可能有问题
					option.SuccessPrint(success)
					fp_succ, err := os.OpenFile(file_success_log, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
					if err != nil {
						fmt.Println("write err :", err)
					}
					fp_succ.WriteString(success)
					defer fp_succ.Close()
				} else {
					fmt.Println("-----------------------------------------------------------------------------------------")
					Failed := res.url + "\t" + res.exp_name[key] + "\t" + "Exploit Failed !" + "\n" //失败时字符串拼接
					option.FailPrint(Failed)
					fp_fail, err := os.OpenFile(file_fail_log, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
					fp_fail.WriteString(Failed)
					if err != nil {
						fmt.Println("write err :", err)
					}

					defer fp_fail.Close()
				}

			}
		}
	}
}

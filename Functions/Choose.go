package Funcs

import (
	option "web_scan/Functions/options"
	Configs "web_scan/config"
)

type extras struct {
	timeout_count map[string]int
	replace       string
}

var extra extras

func Choose() {
	//timeout_count := make(map[string]int)
	extra.timeout_count = make(map[string]int) //这个用来统计是否超过五次超时或失败，如果一个URL超过五次，则直接跳过该url

	if Configs.UserObject.KeyWord != "" {
		option.FindFile(Configs.ConfigJsonMap.Exploit.Path, Configs.UserObject.KeyWord)
	}
	if Configs.UserObject.AllJson == true && Configs.UserObject.OriAddr != "" {
		//-------------------------------
		final_One_url_allJson(extra.timeout_count)

	}
	if Configs.UserObject.GetTitle == true && Configs.UserObject.File != "" {
		option.GetUrlTitle()
	}
	if Configs.UserObject.JsonFile != "" && Configs.UserObject.OriAddr != "" { //单个已完成，调试
		final_Oneurl_OneJson(extra.timeout_count)
	}

	if Configs.UserObject.File != "" && Configs.UserObject.JsonFile != "" {
		final_ALLurl_OneJson(extra.timeout_count)
	}
	if Configs.UserObject.File != "" && Configs.UserObject.AllJson == true {
		final_ALLurl_ALLJson(extra.timeout_count)
	}
}

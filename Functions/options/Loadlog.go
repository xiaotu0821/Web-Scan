package option

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
	Configs "web_scan/config"
)

func init() {
	file, err := os.Open("./config.json")
	if err != nil {
		fmt.Println("Can't found config.json file")
		os.Exit(1)
	}
	fileValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Can't loaded config.json file")
		os.Exit(1)
	}
	err = json.Unmarshal(fileValue, &Configs.ConfigJsonMap)
	if err != nil {
		fmt.Println("Can't to read config.json file")
	}
	FileLog, err := os.OpenFile(Configs.ConfigJsonMap.Exploit.Logs, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Can't to build " + Configs.ConfigJsonMap.Exploit.Logs)
		os.Exit(1)
	}

	Configs.ColorInfo = log.New(os.Stdout, "[INFO]", log.Ldate|log.Ltime)
	Configs.ColorMistake = log.New(io.MultiWriter(FileLog, os.Stderr), "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)
	Configs.ColorSend = log.New(os.Stdout, "[MESSAGE-SEND]", log.Ldate|log.Ltime)
	Configs.ColorSuccess = log.New(os.Stdout, "[SUCCESS]", log.Ldate|log.Ltime)
	Configs.ColorFail = log.New(io.MultiWriter(FileLog, os.Stderr), "[FAILED]", log.Ldate|log.Ltime)

}

func Get_Time() string {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	now_time := fmt.Sprintf("%d年%d月%d日%d时%d分", year, month, day, hour, minute)
	return now_time
}

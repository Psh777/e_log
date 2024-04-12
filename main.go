package e_log

import (
	"fmt"
	"time"
)

const (
	SERVICE        = "billing"
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

var er E
var saveToRedis = false

type E struct {
	Url string `json:"url"`
}

func Init(url string) E {
	er = E{
		Url: url,
	}
	return er
}

func Log(method string, log string) {
	fmt.Println(time.Now().Format(DDMMYYYYhhmmss), "", "Method:", ""+method, "Log:", log, "")
	sent("log", method, log)
	return
}

func Info(method string, log string) {
	fmt.Println(time.Now().Format(DDMMYYYYhhmmss), "\u001B[1m\u001B[32m", "Method:", "\u001B[0m\u001B[32m"+method, "\u001B[1m\u001B[32mLog:\u001B[0m\u001B[32m", log, "\u001B[0m")
	sent("log", method, log)
	return
}

func Err(method string, err error) {
	fmt.Println(time.Now().Format(DDMMYYYYhhmmss), "\u001B[31m", "Method:", method, " Error:", err, "\u001B[0m")
	sent("error", method, err.Error())
	return
}

func Error(err error) {
	fmt.Println(time.Now().Format(DDMMYYYYhhmmss), "\u001B[31m", "Method:", err, "\u001B[0m")
	sent("error", "", err.Error())
	return
}

func ErrText(method string, err string) {
	fmt.Println(time.Now().Format(DDMMYYYYhhmmss), "\u001B[31m", "Method:", method, " Error:", err, "\u001B[0m")
	sent("error", method, err)
	return
}

func sent(types, method, log string) {
	//if saveToRedis {
	//	redis.SetLog(types, SERVICE, method, log)
	//}
}

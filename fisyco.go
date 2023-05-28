package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var total = 0.0
var start = time.Now()

func watchFlow(writer http.ResponseWriter, reader *http.Request) {
	err := reader.ParseForm() // 解析参数
	if err != nil {
		return
	}

	fmt.Println("path:", strings.Join(reader.Form["filepath"], ""))

	filesize, err := strconv.Atoi(strings.Join(reader.Form["filesize"], ""))
	fmt.Println("size:", filesize)
	total += float64(filesize)

	flowPerSec := (total / 1000) / time.Since(start).Seconds()

	_, err = fmt.Printf("flow: %f KB/s\n", flowPerSec)
	if err != nil {
		return
	}

	if time.Since(start).Seconds() >= 60 {
		total = 0.0
		start = time.Now()
	}
}

func main() {
	http.HandleFunc("/", watchFlow)          // 设置访问的路由
	err := http.ListenAndServe(":15752", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
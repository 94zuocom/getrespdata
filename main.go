package main

import (
	"fmt"
	"log"
	"net/http"
	"getrespdata/p" 
)

func main() {
	http.HandleFunc("/r", p.RequestHandler)
	http.HandleFunc("/time", p.TimeHandler)
	http.HandleFunc("/502", p.Status502)
	// 设置静态文件服务
	p.ServeStatic("/static/", "./static")

	// 设置根URL处理程序
	http.HandleFunc("/", p.IndexHandler("./static"))
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}


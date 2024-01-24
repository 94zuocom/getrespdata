package main

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/r", requestHandler)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应内容类型为HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// 创建一个HTML页面
	fmt.Fprintf(w, "<html><head><title>Request Information</title></head><body>")
	fmt.Fprintf(w, "<h1>Request Information</h1>")

	// 打印请求头信息
	fmt.Fprintf(w, "<h2>Request Headers:</h2><ul>")
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "<li>%v: %v</li>", html.EscapeString(name), html.EscapeString(h))
		}
	}
	fmt.Fprintf(w, "</ul>")

	// 打印请求URL和查询参数
	fmt.Fprintf(w, "<h2>Request URL:</h2><p>%v</p>", html.EscapeString(r.URL.String()))
	fmt.Fprintf(w, "<h2>Query Parameters:</h2><ul>")
	query := r.URL.Query()
	for param, values := range query {
		for _, value := range values {
			fmt.Fprintf(w, "<li>%v: %v</li>", html.EscapeString(param), html.EscapeString(value))
		}
	}
	fmt.Fprintf(w, "</ul>")

	// 处理POST或PUT请求的请求体
	if r.Method == "POST" || r.Method == "PUT" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		fmt.Fprintf(w, "<h2>Request Body:</h2><pre>%v</pre>", html.EscapeString(string(body)))
	}

	// 处理文件上传
	if err := r.ParseMultipartForm(32 << 20); err == nil {
		files := r.MultipartForm.File["file"]
		fmt.Fprintf(w, "<h2>Uploaded Files:</h2><ul>")
		for _, fileheader := range files {
			file, err := fileheader.Open()
			if err != nil {
				http.Error(w, "Error opening file", http.StatusInternalServerError)
				return
			}
			defer file.Close()

			fmt.Fprintf(w, "<li>File Name: %s</li>", html.EscapeString(fileheader.Filename))
			fmt.Fprintf(w, "<li>File Size: %s bytes</li>", strconv.FormatInt(fileheader.Size, 10))

			// 丢弃文件内容
			io.Copy(ioutil.Discard, file)
		}
		fmt.Fprintf(w, "</ul>")
	}

	fmt.Fprintf(w, "</body></html>")
}

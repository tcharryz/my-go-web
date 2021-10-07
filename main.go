package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// 接收客户端 request，并将 request 中带的 header 写入 response header
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			w.Header().Add(k, v[0])
			// fmt.Print("---客户端 Request Header开始--- \n")
			// fmt.Print(r.Header)
			// fmt.Print("\n---客户端 Request Header结束--- \n\n")
		}
		//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
		w.Header().Add("Version", os.Getenv("VERSION"))
		//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
		w.WriteHeader(http.StatusOK)
		fmt.Print("---服务器 Response Header开始--- \n")
		fmt.Print(w.Header())
		fmt.Print("\n---服务器 Response Header结束--- \n\n")
		log.Printf("路径: /   返回码:%d   客户端IP: %s", http.StatusOK, r.RemoteAddr)

	})
	//当访问 localhost/healthz 时，应返回200
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("200"))
		w.WriteHeader(http.StatusOK)
		//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
		log.Printf("路径: /healthz   返回码:%d   客户端IP: %s", http.StatusOK, r.RemoteAddr)

	})

	log.Println("Launching my Go HTTP server!")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Server error")
	}
}

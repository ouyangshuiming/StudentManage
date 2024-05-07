package main

import (
	"StudentManage/routers"
	"log"
	"net/http"
)

func main() {
	routers.Router()
	
	server := http.Server{ //server是一个结构体，因此里面有一些属性和方法
		Addr: "127.0.0.1:8080", //指明服务端的ip和端口号，即服务运行在哪个端口
	}
	err := server.ListenAndServe() //也就是说监听8080端口，看看有没有客户端发送请求到这个端口
	if err != nil {                //报错就打印错误信息
		log.Println(err)
	}
}

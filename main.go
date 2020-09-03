package main

import (
	"go-web/logger"
	"net/http"
)

func main() {

	//设置处理静态资源，如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	logger.Info.Println("server starting at http://localhost:8080")

	logger.Error.Fatal(http.ListenAndServe(":8080", nil))

}

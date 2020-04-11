// bootstrap of this project
package main

import (
	"log"
	"net/http"

	"maki.io/demo/shaver/handler"
)

// 注册服务
func registerHandlers() {
	// GET:get the real-url from short-url
	http.HandleFunc("/index", handler.IndexHandler)
	// POST:generate the short-url from the real-url
	http.HandleFunc("/url", handler.UrlHandler)
}

// start http server
func main() {
	registerHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

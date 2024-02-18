package main //let main.go to use this function

import (
	"fmt"
	"net/http"
)

// Hello is a handler function that responds with "hello, world".
// Hello 是一个function，它返回 "hello, world"。

func Hello(w http.ResponseWriter, r *http.Request){
	// w is the response writer used to write the HTTP response back to the client.
	// w 為將HTTP回應寫回client端的變數容器。
	// r is the HTTP request received from the client.
// r 是從client端收到的 HTTP 请求。
	fmt.Fprint(w,"hello ,world")
}
package main //let main.go to use this function

import (
	"fmt"
	"net/http"
)

// Hello is a handler function that responds with "hello, world".
// Hello 是一个function，它返回 "hello, world"。

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	// w is the response writer used to write the HTTP response back to the client.
	// w 為將HTTP回應寫回client端的變數容器。
	// r is the HTTP request received from the client.
	// r 是從client端收到的 HTTP 请求。
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
		//no need for "," in definition of struc

	}{Status: "active",
		Message: "Dennis's project is running.",
		Version: "1.0.1",
	}
	// initialize payload
	//fmt.Fprintf(w, "hello, world from %s", app.Domain)
	_=app.writeJSON(w,http.StatusOK,payload)
}

//

func (app *application) ALLMovies(w http.ResponseWriter, r *http.Request) { //change from response to request
	movies ,err :=app.DB.ALLMovies()
	if err != nil {
		fmt.Println(err)
        return
	}
	_=app.writeJSON(w,http.StatusOK, movies)
}

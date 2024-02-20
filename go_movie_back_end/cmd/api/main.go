package main

import (
	"fmt"
	"log"
	"net/http"
)


const port=8080

type application struct {
	Domain string

}

//
//func (app application) handlerFunc(w http.ResponseWriter, r *http.Request) {
  //  fmt.Fprintf(w, "Hello, %s!", app.Domain)
//}
//
func main(){
	//1.設置應用程式config/setting up application config
var app application//從application struc中提取資訊並做application設置
	//2.read from cmd line


//3.connect to a db
app.Domain = "example.com"
log.Println("starting application on port",port)
//3.1Define routes and handlers
//no need after adding app.routes function , http.HandleFunc("/", Hello)//using http start handle function->Hello function in handler.go on http://8080"/""

//4.starting a server
err :=http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
if err!=nil{
	log.Fatal(err)
}
}
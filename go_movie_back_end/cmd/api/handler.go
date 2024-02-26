package main //let main.go to use this function

import (
	"encoding/json"
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
	out, err := json.Marshal(payload)
	// 1.Marshal is one of the default object in json package
	// 2.定義out(寫在瀏覽器inspector裡面的物件)和錯誤訊息為Marshal(使用payload當作參數)
	if err != nil {
		fmt.Println(err)

	}
	w.Header().Set("Content-Type", "application/json") //此行從 http.ResponseWriter (w) 中擷取標頭資料，將 "Content-Type" 標頭設置為 "application/json"。這表示對客戶端回應的主體內容將是 JSON 格式。
	w.WriteHeader(http.StatusOK)                       //此行使用 http.ResponseWriter 介面 (w) 的 WriteHeader 方法將回應的 HTTP 狀態碼設置為 200 (OK)。這表示請求成功。
	w.Write(out)                                       //此行將 out 變數的內容寫入回應主體。假設 out 變數是一個包含 JSON 編碼回應資料的位元組切片。這將實際的 JSON 資料作為回應主體發送給客戶端。
	fmt.Fprintf(w, "hello, world from %s", app.Domain)
}

//

func (app *application) ALLMovies(w http.ResponseWriter, r *http.Request) { //change from response to request
	movies ,err :=app.DB.ALLMovies()
	if err != nil {
		fmt.Println(err)
        return
	}
	//
	out, err := json.Marshal(movies)
	// 1.Marshal is one of the default object in json package
	// 2.定義out(寫在瀏覽器inspector裡面的物件)和錯誤訊息為Marshal(使用payload當作參數)
	if err != nil {
		fmt.Println(err)

	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json") //此行從 http.ResponseWriter (w) 中擷取標頭資料，將 "Content-Type" 標頭設置為 "application/json"。這表示對客戶端回應的主體內容將是 JSON 格式。
	w.WriteHeader(http.StatusOK)                       //此行使用 http.ResponseWriter 介面 (w) 的 WriteHeader 方法將回應的 HTTP 狀態碼設置為 200 (OK)。這表示請求成功。
	w.Write(out)

	//
}

package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JSONResponse struct{
Error bool 	`json:"error"`
Message string `json:"message"`
Data interface{} `json:"data,omitempty"`
}

func (app *application) writeJSON(w http.ResponseWriter,status int,data interface{},headers ...http.Header) error{
out, err :=json.Marshal(data)// 1.Marshal is one of the default object in json package
// 2.定義out(寫在瀏覽器inspector裡面的物件)和錯誤訊息為Marshal(使用payload當作參數)
if err !=nil{
	return err
}	
if len(headers)>0{
	for key, value:=range headers[0]{
		w.Header()[key]=value
	}
}

w.Header().Set("Content-Type","application/json")//此行從 http.ResponseWriter (w) 中擷取標頭資料，將 "Content-Type" 標頭設置為 "application/json"。這表示對客戶端回應的主體內容將是 JSON 格式。

w.WriteHeader(status)//此行使用 http.ResponseWriter 介面 (w) 的 WriteHeader 方法將回應的 HTTP 狀態碼設置為 200 (OK)。這表示請求成功。

_, err =w.Write(out)//此行將 out 變數的內容寫入回應主體。假設 out 變數是一個包含 JSON 編碼回應資料的位元組切片。這將實際的 JSON 資料作為回應主體發送給客戶端。

if err!=nil{
	return err
}
return nil

}

func (app *application) readJSON(w http.ResponseWriter,r *http.Request, data interface{}) error{
	maxBytes:=1024*1024//1MB
	r.Body=http.MaxBytesReader(w,r.Body,int64(maxBytes))
dec:=json.NewDecoder(r.Body)
dec.DisallowUnknownFields()
err:=dec.Decode(data)
if err!=nil{
	return err
}
err=dec.Decode(&struct{}{})
if err!=io.EOF{
	return errors.New("body must only contains a single JSON value.")
}
return nil
}
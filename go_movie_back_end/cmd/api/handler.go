package main // 宣告這個文件屬於 main 套件 // Declare this file as part of the main package

import (
	"log"      // 匯入 log 套件以進行日誌記錄 // Import the log package for logging
	"net/http" // 匯入 net/http 套件以處理 HTTP 請求和回應 // Import the net/http package to handle HTTP requests and responses
)

// Hello is a handler function that responds with "hello, world".
// Hello 是一個處理函式，它回應 "hello, world"。
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	// w 是用於將 HTTP 回應寫回客戶端的響應寫入器。
	// w is the response writer used to write the HTTP response back to the client.
	// r 是從客戶端收到的 HTTP 請求。
	// r is the HTTP request received from the client.
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Dennis's project is running.",
		Version: "1.0.1",
	}
	// 初始化 payload 結構體
	// Initialize the payload struct

	// 使用 writeJSON 方法回應 JSON 格式的 payload
	// Use the writeJSON method to respond with the payload in JSON format
	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) ALLMovies(w http.ResponseWriter, r *http.Request) {
	// 從資料庫獲取所有電影
	// Retrieve all movies from the database
	movies, err := app.DB.ALLMovies()
	if err != nil {
		// 如果有錯誤，使用 errorJSON 方法回應錯誤訊息
		// If there is an error, respond with an error message using the errorJSON method
		app.errorJSON(w, err)
		return
	}
	// 使用 writeJSON 方法回應 JSON 格式的電影列表
	// Use the writeJSON method to respond with the list of movies in JSON format
	_ = app.writeJSON(w, http.StatusOK, movies)
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// 創建一個 jwtUser 實例
	// Create a jwtUser instance
	u := jwtUser{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
	}
	// 生成令牌對
	// Generate token pairs
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		// 如果有錯誤，使用 errorJSON 方法回應錯誤訊息
		// If there is an error, respond with an error message using the errorJSON method
		app.errorJSON(w, err)
		return
	}

	// 記錄令牌並將訪問令牌寫回客戶端
	// Log the tokens and write the access token back to the client
	log.Println(tokens.Token)
	w.Write([]byte(tokens.Token))
}

package main

import "net/http"

// enableCORS 為啟用跨來源資源共享（CORS）的函數。它接受一個 http.Handler 作為參數並返回一個 http.Handler。
// enableCORS is a function to enable Cross-Origin Resource Sharing (CORS). It takes an http.Handler as a parameter and returns an http.Handler.
func (app *application) enableCORS(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 設置 Access-Control-Allow-Origin 標頭，允許所有來源的跨來源請求。
        // Set the Access-Control-Allow-Origin header to allow cross-origin requests from all sources.
        w.Header().Set("Access-Control-Allow-Origin", "http://*")

        if r.Method == "OPTIONS" {
            // 如果是 OPTIONS 請求，則設置跨來源請求相關的標頭信息。
            // If it's an OPTIONS request, set the headers related to cross-origin requests.
            w.Header().Set("Access-Control-Allow-Credentials", "true")
            w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "ACCEPT, Content-Type, X-CSRF-Token, Authorization")
            return
        } else {
            // 如果不是 OPTIONS 請求，則將請求轉發給下一個處理程序。
            // If it's not an OPTIONS request, forward the request to the next handler.
            h.ServeHTTP(w, r)
        }
    })
}


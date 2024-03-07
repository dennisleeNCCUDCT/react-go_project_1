package main

import "time"
type Authenticator struct { // 認證結構 // Auth struct
	Issuer        string        // 發行者 // Issuer of the token
	Audience      string        // 觀眾 // Audience of the token
	SecretKey     string        // 秘密 // Secret key for signing the token
	TokenExpiry   time.Duration // 令牌到期時間 // Expiry duration for the token
	RefreshExpiry time.Duration // 刷新到期時間 // Expiry duration for the refresh token
	CookieDomain  string        // Cookie域 // Domain for the cookie
	CookiePath    string        // Cookie路徑 // Path for the cookie
	CookieName    string        // Cookie名稱 // Name of the cookie
}
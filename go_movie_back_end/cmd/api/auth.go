package main // 主套件 // Main package

import ( // 導入套件 // Import packages
	"fmt"  // 格式化套件 // Format package
	"time" // 時間套件 // Time package

	"github.com/golang-jwt/jwt/v4" // JWT套件 // JWT package
)

type Auth struct { // 認證結構 // Auth struct
	Issuer        string        // 發行者 // Issuer of the token
	Audience      string        // 觀眾 // Audience of the token
	Secret        string        // 秘密 // Secret key for signing the token
	TokenExpiry   time.Duration // 令牌到期時間 // Expiry duration for the token
	RefreshExpiry time.Duration // 刷新到期時間 // Expiry duration for the refresh token
	CookieDomain  string        // Cookie域 // Domain for the cookie
	CookiePath    string        // Cookie路徑 // Path for the cookie
	CookieName    string        // Cookie名稱 // Name of the cookie
}

type jwtUser struct { // JWT用戶結構 // JWT user struct
	ID        int    `json:"id"`         // 用戶ID // User ID
	FirstName string `json:"first_name"` // 名字 // First name
	LastName  string `json:"last_name"`  // 姓氏 // Last name
}

type TokenPairs struct { // 令牌對結構 // Token pairs struct
	Token        string `json:"access_token"`  // 存取令牌 // Access token
	RefreshToken string `json:"refresh_token"` // 刷新令牌 // Refresh token
}

type Claims struct { // 聲明結構 // Claims struct
	jwt.RegisteredClaims // 註冊聲明 // Registered claims
}

// 生成令牌對函數 // Generate token pair function
func (j *Auth) GenerateTokenPair(user *jwtUser) (TokenPairs, error) {
	// 創建一個令牌 // Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// 設置聲明 // Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	claims["sub"] = fmt.Sprint(user.ID)
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"

	// 設置JWT的到期時間 // Set the expiry for JWT
	claims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()

	// 創建一個簽名的令牌 // Create a signed token
	signedAccessToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err // 返回錯誤 // Return error
	}

	// 創建一個刷新令牌並設置聲明 // Create a refresh token and set claims
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprint(user.ID)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()

	// 設置刷新令牌的到期時間 // Set the expiry for the refresh token
	refreshTokenClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()

	// 創建簽名的刷新令牌 // Create a signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err // 返回錯誤 // Return error
	}

	// 創建TokenPairs並用簽名的令牌填充 // Create TokenPairs and populate with signed tokens
	var tokenPairs = TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken, 
	}

	// 返回TokenPairs // Return TokenPairs
	return tokenPairs, nil
}
type Authenticator struct {
    // Add fields related to authentication, such as secret keys, token expiration, etc.
    SecretKey string
    // Other fields...
}
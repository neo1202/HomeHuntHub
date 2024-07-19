package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

/*
Issuer: 類似xxx.com, 誰發行的token
Audience: 誰能用
Secret: 加密的
RefreshExpiry: cookie相關
*/
type Auth struct {
	Issuer        string
	Audience      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

// 最少需要的資訊讓我們能發token
type jwtUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Token: actual jwt
type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// 希望額外附的資訊，例如這個token只給這個audience
type Claims struct {
	jwt.RegisteredClaims
}

func (j *Auth) GenerateTokenPair(user *jwtUser) (TokenPairs, error) {
	// Create a token
	token := jwt.New(jwt.SigningMethodHS256) // return a pointer to jwt token
	// Set the claims
	claims := token.Claims.(jwt.MapClaims) //聲明我想要用Map這種形式的claims
	claims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	claims["sub"] = fmt.Sprint(user.ID)
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"
	// Set the expiry for JWT
	claims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()
	// Create a signed token
	signedAccessToken, err := token.SignedString([]byte(j.Secret)) // 使用密钥签名令牌
	if err != nil {
		return TokenPairs{}, err
	}
	// Create a refresh token and set claims
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprint(user.ID) // fmt.Sprint可以將ID轉成string來存
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()
	// Set the expiry for the refresh token
	refreshTokenClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()
	// Create signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}
	// Create TokenPairs and populate with signed tokens
	var tokenPairs = TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}
	// Return TokenPairs
	return tokenPairs, nil
}

// refresh token 的作用是在用户的访问令牌过期时，用于获取新的访问令牌，而无需用户再次登录。刷新令牌通常具有较长的有效期，保证用户的持续登录体验
func (j *Auth) GetRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    refreshToken, // 把refreshToken存在安全的cookie之後有需要會調用
		Expires:  time.Now().Add(j.RefreshExpiry),
		MaxAge:   int(j.RefreshExpiry.Seconds()),
		SameSite: http.SameSiteDefaultMode,
		Domain:   j.CookieDomain,
		HttpOnly: true, //javascript無法讀
		Secure:   true,
	}
}
func (j *Auth) GetExpiredRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteDefaultMode,
		Domain:   j.CookieDomain,
		HttpOnly: true, //javascript無法讀
		Secure:   true,
	}
}


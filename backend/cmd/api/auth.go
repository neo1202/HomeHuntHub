package main

import (
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

package jwt

import "github.com/golang-jwt/jwt/v5"

type ContextKey string

const (
	dataCtxKey ContextKey = "jwtData"
)

type Payload struct {
	UserId int
}

type Claims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

func (c Claims) Payload() *Payload {
	return &Payload{
		UserId: c.UserId,
	}
}

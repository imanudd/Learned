package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

var IsAuth = middleware.JWTConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

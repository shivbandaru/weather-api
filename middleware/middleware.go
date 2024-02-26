package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/labstack/echo"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	//can be filled in if needed
}

// RequestID adds request-id to http request
func (m *GoMiddleware) RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("RequestID", generateID())
		return next(c)
	}
}

// InitMiddleware intialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}

// Local method to generate a random ID to uniquole identify the request from the client
func generateID() string {
	r := make([]byte, 12)
	_, err := rand.Read(r)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(r)
}

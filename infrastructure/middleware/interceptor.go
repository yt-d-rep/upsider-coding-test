package middleware

import (
	"upsider-base/domain/auth"

	"github.com/gin-gonic/gin"
)

// Intetceptor は処理の前後に挟む処理を定義します
type Interceptor struct {
	tokenSvc auth.TokenService
}

func (i *Interceptor) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(401)
			return
		}
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.AbortWithStatus(401)
			return
		}
		token := authHeader[7:]
		ok, err := i.tokenSvc.Validate(auth.NewToken(token))
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		if !ok {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}

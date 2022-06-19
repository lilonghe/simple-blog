package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"simple-blog/pkg/utils"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id := session.Get("id")
		if id == nil {
			utils.GetMessageError("NO_LOGIN", "No login", c)
			c.Abort()
		}
		c.Set("id", id)
	}
}

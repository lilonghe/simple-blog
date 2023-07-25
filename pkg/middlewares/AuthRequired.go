package middlewares

import (
	"simple-blog/pkg/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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

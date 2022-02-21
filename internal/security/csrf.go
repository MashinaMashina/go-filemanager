package security

import (
	"file-manager/internal/errorpage"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func RegisterMiddleware(r *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions("mysession", store))
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			errorpage.Page("Неверный CSRF токен", c)
			c.Abort()
		},
	}))
}

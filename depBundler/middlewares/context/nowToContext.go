package context

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/session"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NowToContext(session session.SessionResolver) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := session.GetCurrentSession(c.Request.Context())
		if err != nil{
			fmt.Println(err)
		} else {
			fmt.Println(session.GetNow())
		}
		c.Next()
	}
}

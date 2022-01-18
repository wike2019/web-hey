package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"hey-back/ctl"
	"net/http"
)

//go:embed html
var dist  embed.FS





func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println(c.Request.URL)
		if method != "" {
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func main() {
	web:=goft.Ignite(Cors(), func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.AbortWithStatusJSON(400, gin.H{"error": err})
			}
		}()
		context.Next()
	}).
		Mount("/v1",
				ctl.NewCtl(),
		)
	web.Engine.GET("/", func(c *gin.Context) {
		c.FileFromFS("/html/"+c.Param("filepath"),http.FS(dist))
	})
	web.Engine.ForwardedByClientIP=true
	web.LaunchWithPort(7377)

}

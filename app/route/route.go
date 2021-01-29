package route

import (
	"net/http"
	"zhihu-analysis/app/controller/gif"
	"zhihu-analysis/app/controller/simapleHomeApp"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.Any("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	simapleHomeAppRoute := r.Group("/simapleHomeApp")
	simapleHomeAppRoute.GET("", simapleHomeApp.XlsxAnalysis)

	gifRoute := r.Group("/gif")
	gifRoute.GET("", gif.XlsxAnalysis)
	gifRoute.GET("/download", gif.DownloadGif)
}

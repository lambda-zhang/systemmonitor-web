package routes

import (
	ctr "github.com/lambda-zhang/systemmonitor-web/controllers"

	"net/http"
	"os"
	"time"

	//"github.com/DeanThompson/ginpprof"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	//ginpprof.Wrap(r)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.StaticFile("/favicon.png", "./webpage/dist/static/images/favicon.png")

	r.StaticFile("/", "./webpage/dist/")
	r.StaticFS("/static", http.Dir("./webpage/dist/static"))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"text": "Not Found.",
		})
	})

	corsObject := cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:    []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		// ExposeHeaders:    []string{"Content-Length"},
		// AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	})
	r.Use(corsObject)

	r.GET("/ping", ctr.Pong)
	r.GET("/ws", ctr.WsHandler)
	usedb := os.Getenv("USEDB")
	if usedb == "true" {
		r.GET("/sysusage", ctr.GetCpuState)
		r.GET("/osinfo", ctr.GetOsInfo)
		r.GET("/netifusage", ctr.GetNetState)
		r.GET("/tcpusage", ctr.GetTcpState)
		r.GET("/thermalstate", ctr.GetThermalState)
		r.GET("/diskusage", ctr.GetDiskState)
	}
	return r
}

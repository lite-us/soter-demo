package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(NoCache)
	g.Use(Options)
	g.Use(Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The URL Route NOT Found")
	})

	path, err := os.Getwd()
	if err != nil {
		path = "."
	}
	filepath := fmt.Sprintf("%s/public", path)
	g.Static("/index", filepath)

	return g
}

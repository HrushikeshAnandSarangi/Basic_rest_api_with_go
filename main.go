package main

import (
	"github.com/HrushikeshAnandSarangi/go-rest/config"
	"github.com/HrushikeshAnandSarangi/go-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	routes.AuthRoutes(r)
	r.Run(":8000")
}

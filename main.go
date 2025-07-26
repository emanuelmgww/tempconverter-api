package main

import (
	"github.com/emanuelmgww/conversorDeTemperaturas/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":3333")
}

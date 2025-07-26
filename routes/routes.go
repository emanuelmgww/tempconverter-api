package routes

import (
	"github.com/emanuelmgww/conversorDeTemperaturas/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/converter/celsius", controllers.ConversorCelsius)
	r.GET("/converter/fahrenheit", controllers.ConversorFahrenheit)
	r.GET("/converter/kelvin", controllers.ConversorKelvin)
}

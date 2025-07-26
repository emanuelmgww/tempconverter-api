package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/emanuelmgww/conversorDeTemperaturas/services"
	"github.com/gin-gonic/gin"
)

type RespostaCelsius struct {
	Celsius    string `json:"celsius"`
	Fahrenheit string `json:"fahrenheit"`
	Kelvin     string `json:"kelvin"`
}

func ConversorCelsius(c *gin.Context) {
	celsiusStr := c.Query("celsius")

	celsiusStr = strings.TrimSpace(celsiusStr)
	celsius, err := strconv.ParseFloat(celsiusStr, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "temperatura inválida"})
		return
	}

	if celsius < -273.15 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "celsius abaixo de zero absoluto (-273.15ºC) não é válido"})
		return
	}

	fahrenheit, kelvin := services.ConversorCelsius(celsius)

	resposta := RespostaCelsius{
		Celsius:    fmt.Sprintf("%.2f", celsius),
		Fahrenheit: fmt.Sprintf("%.2f", fahrenheit),
		Kelvin:     fmt.Sprintf("%.2f", kelvin),
	}

	c.JSON(http.StatusOK, resposta)
}

type RespostaFahrenheit struct {
	Fahrenheit string `json:"fahrenheit"`
	Celsius    string `json:"celsius"`
	Kelvin     string `json:"kelvin"`
}

func ConversorFahrenheit(c *gin.Context) {
	fahrenheitStr := c.Query("fahrenheit")

	fahrenheitStr = strings.TrimSpace(fahrenheitStr)
	fahrenheit, err := strconv.ParseFloat(fahrenheitStr, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "temperatura inválida"})
		return
	}

	if fahrenheit < -459.67 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fahrenheit abaixo de zero absoluto (-459.67ºF) não é válido"})
		return
	}

	celsius, kelvin := services.ConversorFahrenheit(fahrenheit)

	resposta := RespostaFahrenheit{
		Fahrenheit: fmt.Sprintf("%.2f", fahrenheit),
		Celsius:    fmt.Sprintf("%.2f", celsius),
		Kelvin:     fmt.Sprintf("%.2f", kelvin),
	}

	c.JSON(http.StatusOK, resposta)
}

type RespostaKelvin struct {
	Kelvin     string `json:"kelvin"`
	Celsius    string `json:"celsius"`
	Fahrenheit string `json:"fahrenheit"`
}

func ConversorKelvin(c *gin.Context) {
	kelvinStr := c.Query("kelvin")

	kelvinStr = strings.TrimSpace(kelvinStr)
	kelvin, err := strconv.ParseFloat(kelvinStr, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "temperatura inválida"})
		return
	}

	if kelvin < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kelvin não pode ser negativo"})
		return
	}

	celsius, fahrenheit := services.ConversorKelvin(kelvin)

	resposta := RespostaKelvin{
		Kelvin:     fmt.Sprintf("%.2f", kelvin),
		Celsius:    fmt.Sprintf("%.2f", celsius),
		Fahrenheit: fmt.Sprintf("%.2f", fahrenheit),
	}

	c.JSON(http.StatusOK, resposta)
}

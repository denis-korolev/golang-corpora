package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"parser/config"
)

func main() {
	fmt.Println("Настраиваем конфиг")
	config.CalculatetConfig()
	fmt.Println("Запускаем web приложение")

	gin.SetMode(viper.GetString("GIN_MODE"))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(viper.GetString("WEB_HOST") + ":" + viper.GetString("WEB_PORT"))
}

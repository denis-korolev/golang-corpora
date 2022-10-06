package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"parser/config"
	"parser/config/route"
)

func main() {
	fmt.Println("Настраиваем конфиг")
	config.CalculatetConfig()
	fmt.Println("Запускаем web приложение")

	gin.SetMode(viper.GetString("GIN_MODE"))
	router := route.SetupRoutes()
	router.Run(viper.GetString("WEB_HOST") + ":" + viper.GetString("WEB_PORT"))
}

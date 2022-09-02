package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
)

func CalculatetConfig() {

	// вычисление корня проекта
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filepath.Join(b, "/../"))
	viper.Set("ROOT_PATH", basepath)
	viper.Set("VAR_PATH", filepath.Join(basepath, "/var"))

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(basepath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

}

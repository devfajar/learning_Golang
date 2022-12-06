package main

import (
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"net/http"
	"fmt"
)

func main() {
	e := echo.New()

	// viper.SetConfigType("json")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")


	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/index", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, true)
	})

	// Watcher
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	e.Logger.Print("Starting", viper.GetString("appname"))
	e.Logger.Fatal(e.Start(":" + viper.GetString("server.port")))
}
package main

import (
	"net/http"

	"github.com/DyrisIT/AutoconfigAgent/cli"
	"github.com/DyrisIT/AutoconfigAgent/handler"
	"github.com/spf13/viper"
)

func main() {
	cli.Setup()

	http.HandleFunc("/mail/config-v1.1.xml", handler.ThunderbirdAutoconfig)
	http.HandleFunc("/autodiscover/autodiscover.xml", handler.OutlookAutodiscover)

	var addr = viper.GetString("bind") + ":" + viper.GetString("port")
	http.ListenAndServe(addr, nil)
}

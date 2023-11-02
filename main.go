package main

import (
	"log"
	"net/http"

	"github.com/DyrisIT/AutoconfigAgent/cli"
	"github.com/DyrisIT/AutoconfigAgent/handler"
	"github.com/spf13/viper"
)

func main() {
	cli.Setup()

	http.HandleFunc("/mail/config-v1.1.xml", handler.Autoconfig)
	http.HandleFunc("/autodiscover/autodiscover.xml", handler.Autodiscover)
	http.HandleFunc("/.well-known/mobileconfig", handler.Mobileconfig)

	var addr = viper.GetString("bind") + ":" + viper.GetString("port")
	log.Printf("Starting API on %s", addr)
	http.ListenAndServe(addr, nil)
}

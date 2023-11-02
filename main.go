package main

import (
	"net/http"

	"github.com/DyrisIT/AutoconfigAgent/handler"
)

func main() {
	http.HandleFunc("/mail/config-v1.1.xml", handler.ThunderbirdAutoconfig)
	http.HandleFunc("/autodiscover/autodiscover.xml", handler.OutlookAutodiscover)

	http.ListenAndServe(":8080", nil)
}

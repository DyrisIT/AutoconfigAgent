package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DyrisIT/AutoconfigAgent/cli"
)

// Spec: https://developer.apple.com/business/documentation/Configuration-Profile-Reference.pdf

// For iOS
func Mobileconfig(w http.ResponseWriter, r *http.Request) {
	email, domain, err := validateEmail(r.URL.Query().Get("emailaddress"))
	if err != nil {
		log.Printf("Request Error: %s", r.URL)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var imap_ssl = "true"
	if cli.Get("IMAP_SECURITY") == "plain" {
		imap_ssl = "false"
	}
	var imap_auth = "EmailAuthPassword"
	if cli.Get("IMAP_AUTH") == "none" {
		imap_auth = "EmailAuthNone"
	}
	var smtp_ssl = "true"
	if cli.Get("SMTP_SECURITY") == "plain" {
		smtp_ssl = "false"
	}
	var smtp_auth = "EmailAuthPassword"
	if cli.Get("SMTP_AUTH") == "none" {
		smtp_auth = "EmailAuthNone"
	}

	plist := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>EmailAddress</key>
    <string>%s</string>
    <key>IncomingMailServerAuthentication</key>
    <string>%s</string>
    <key>IncomingMailServerHostname</key>
    <string>%s</string>
    <key>IncomingMailServerPortNumber</key>
    <integer>%s</integer>
    <key>IncomingMailServerUseSSL</key>
    <%s/>
    <key>IncomingMailServerUsername</key>
    <string>%s</string>
    <key>OutgoingMailServerAuthentication</key>
    <string>%s</string>
    <key>OutgoingMailServerHostname</key>
    <string>%s</string>
    <key>OutgoingMailServerPortNumber</key>
    <integer>%s</integer>
    <key>OutgoingMailServerUseSSL</key>
    <%s/>
    <key>OutgoingMailServerUsername</key>
    <string>%s</string>
</dict>
</plist>`,
		email,
		imap_auth, cli.Get("IMAP_SUBDOMAIN")+"."+domain, cli.Get("IMAP_PORT"), imap_ssl, email,
		smtp_auth, cli.Get("SMTP_SUBDOMAIN")+"."+domain, cli.Get("SMTP_PORT"), smtp_ssl, email,
	)

	w.Header().Set("Content-Type", "application/x-apple-aspen-config; charset=utf-8")
	fmt.Fprint(w, plist)
}

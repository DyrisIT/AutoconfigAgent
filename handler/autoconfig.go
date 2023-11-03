package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DyrisIT/AutoconfigAgent/cli"
)

// Spec: https://wiki.mozilla.org/Thunderbird:Autoconfiguration:ConfigFileFormat

// For Thunderbird
func Autoconfig(w http.ResponseWriter, r *http.Request) {
	email, domain, err := validateEmail(r.URL.Query().Get("emailaddress"))
	if err != nil {
		log.Printf("Request Error: %s", r.RequestURI)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	xml := fmt.Sprintf(`<?xml version="1.0"?>
<clientConfig version="1.1">
    <emailProvider id="%s">
        <domain>%s</domain>
        <displayName>%s</displayName>
        <displayShortName>%s</displayShortName>
        <incomingServer type="imap">
            <hostname>%s</hostname>
            <port>%s</port>
            <socketType>%s</socketType>
            <username>%s</username>
            <authentication>%s</authentication>
        </incomingServer>
        <outgoingServer type="smtp">
            <hostname>%s</hostname>
            <port>%s</port>
            <socketType>%s</socketType>
            <username>%s</username>
            <authentication>%s</authentication>
        </outgoingServer>
    </emailProvider>
</clientConfig>`,
		domain, domain, domain, domain,
		cli.Get("IMAP_SUBDOMAIN")+"."+domain, cli.Get("IMAP_PORT"), cli.Get("IMAP_SECURITY"), email, cli.Get("IMAP_AUTH"),
		cli.Get("SMTP_SUBDOMAIN")+"."+domain, cli.Get("SMTP_PORT"), cli.Get("SMTP_SECURITY"), email, cli.Get("SMTP_AUTH"),
	)

	w.Header().Set("Content-Type", "application/xml")
	fmt.Fprint(w, xml)
}

package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DyrisIT/AutoconfigAgent/cli"
)

// Spec: https://github.com/gronke/email-autodiscover/blob/master/mail/autodiscover.xml

// For Outlook
func Autodiscover(w http.ResponseWriter, r *http.Request) {
	email, domain, err := validateEmail(r.URL.Query().Get("emailaddress"))
	if err != nil {
		log.Printf("Request Error: %s", r.URL)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var imap_ssl = "on"
	if cli.Get("IMAP_SECURITY") == "plain" {
		imap_ssl = "off"
	}
	var imap_auth = "on"
	if cli.Get("IMAP_AUTH") == "none" {
		imap_auth = "off"
	}
	var smtp_ssl = "on"
	if cli.Get("SMTP_SECURITY") == "plain" {
		smtp_ssl = "off"
	}
	var smtp_auth = "on"
	if cli.Get("SMTP_AUTH") == "none" {
		smtp_auth = "off"
	}

	xml := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<Autodiscover xmlns="http://schemas.microsoft.com/exchange/autodiscover/responseschema/2006">
    <Response xmlns="http://schemas.microsoft.com/exchange/autodiscover/outlook/responseschema/2006a">
        <Account>
            <AccountType>email</AccountType>
            <Action>settings</Action>
            <Protocol>
                <Type>IMAP</Type>
                <Server>%s</Server>
                <Port>%s</Port>
                <DomainRequired>on</DomainRequired>
                <LoginName>%s</LoginName>
                <SPA>off</SPA>
                <SSL>%s</SSL>
                <AuthRequired>%s</AuthRequired>
            </Protocol>
            <Protocol>
                <Type>SMTP</Type>
                <Server>%s</Server>
                <Port>%s</Port>
                <DomainRequired>on</DomainRequired>
                <LoginName>%s</LoginName>
                <SPA>off</SPA>
                <SSL>%s</SSL>
                <AuthRequired>%s</AuthRequired>
            </Protocol>
        </Account>
    </Response>
</Autodiscover>`,
		cli.Get("IMAP_SUBDOMAIN")+"."+domain, cli.Get("IMAP_PORT"), email, imap_ssl, imap_auth,
		cli.Get("SMTP_SUBDOMAIN")+"."+domain, cli.Get("SMTP_PORT"), email, smtp_ssl, smtp_auth,
	)

	w.Header().Set("Content-Type", "application/xml")
	fmt.Fprint(w, xml)
}

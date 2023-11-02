package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/mail/config-v1.1.xml", handleThunderbirdAutoconfig)
	http.HandleFunc("/autodiscover/autodiscover.xml", handleOutlookAutodiscover)

	http.ListenAndServe(":8080", nil)
}

func handleThunderbirdAutoconfig(w http.ResponseWriter, r *http.Request) {
	xml := `<?xml version="1.0"?>
<clientConfig version="1.1">
    <emailProvider id="example.com">
        <domain>example.com</domain>
        <displayName>Example Mail</displayName>
        <displayShortName>Example</displayShortName>
        <incomingServer type="imap">
            <hostname>imap.example.com</hostname>
            <port>993</port>
            <socketType>SSL</socketType>
            <username>%s</username>
            <authentication>OAuth2</authentication>
        </incomingServer>
        <outgoingServer type="smtp">
            <hostname>smtp.example.com</hostname>
            <port>465</port>
            <socketType>SSL</socketType>
            <username>%s</username>
            <authentication>OAuth2</authentication>
        </outgoingServer>
    </emailProvider>
</clientConfig>`

	email := r.URL.Query().Get("emailaddress")
	if email == "" {
		email = "user@example.com" // default email
	}

	w.Header().Set("Content-Type", "application/xml")
	fmt.Fprintf(w, xml, email, email)
}

func handleOutlookAutodiscover(w http.ResponseWriter, r *http.Request) {
	xml := `<?xml version="1.0" encoding="utf-8"?>
<Autodiscover xmlns="http://schemas.microsoft.com/exchange/autodiscover/responseschema/2006">
    <Response xmlns="http://schemas.microsoft.com/exchange/autodiscover/outlook/responseschema/2006a">
        <Account>
            <AccountType>email</AccountType>
            <Action>settings</Action>
            <Protocol>
                <Type>IMAP</Type>
                <Server>imap.example.com</Server>
                <Port>993</Port>
                <DomainRequired>on</DomainRequired>
                <LoginName>%s</LoginName>
                <SPA>off</SPA>
                <SSL>on</SSL>
                <AuthRequired>on</AuthRequired>
            </Protocol>
            <Protocol>
                <Type>SMTP</Type>
                <Server>smtp.example.com</Server>
                <Port>465</Port>
                <DomainRequired>on</DomainRequired>
                <LoginName>%s</LoginName>
                <SPA>off</SPA>
                <SSL>on</SSL>
                <AuthRequired>on</AuthRequired>
            </Protocol>
        </Account>
    </Response>
</Autodiscover>`

	email := r.URL.Query().Get("emailaddress")
	if email == "" {
		email = "user@example.com" // default email
	}

	w.Header().Set("Content-Type", "application/xml")
	fmt.Fprintf(w, xml, email, email)
}

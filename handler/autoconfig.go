package handler

import (
	"fmt"
	"net/http"
)

// Spec: https://wiki.mozilla.org/Thunderbird:Autoconfiguration:ConfigFileFormat

// For Thunderbird
func Autoconfig(w http.ResponseWriter, r *http.Request) {
	email, domain, err := validateEmail(r.URL.Query().Get("emailaddress"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	w.Header().Set("Content-Type", "application/xml")
	fmt.Fprintf(w, xml, email, email)
}

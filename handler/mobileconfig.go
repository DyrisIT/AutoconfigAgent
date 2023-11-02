package handler

import (
	"fmt"
	"net/http"
)

// For iOS
func Mobileconfig(w http.ResponseWriter, r *http.Request) {
	email, err := validateEmail(r.URL.Query().Get("emailaddress"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	plist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>EmailAccountDescription</key>
    <string>Example Mail</string>
    <key>EmailAccountName</key>
    <string>%s</string>
    <key>EmailAddress</key>
    <string>%s</string>
    <key>IncomingMailServerAuthentication</key>
    <string>OAuth2</string>
    <key>IncomingMailServerHostname</key>
    <string>imap.example.com</string>
    <key>IncomingMailServerPortNumber</key>
    <integer>993</integer>
    <key>IncomingMailServerUseSSL</key>
    <true/>
    <key>IncomingMailServerUsername</key>
    <string>%s</string>
    <key>OutgoingMailServerAuthentication</key>
    <string>OAuth2</string>
    <key>OutgoingMailServerHostname</key>
    <string>smtp.example.com</string>
    <key>OutgoingMailServerPortNumber</key>
    <integer>465</integer>
    <key>OutgoingMailServerUseSSL</key>
    <true/>
    <key>OutgoingMailServerUsername</key>
    <string>%s</string>
</dict>
</plist>`

	w.Header().Set("Content-Type", "application/x-apple-aspen-config; charset=utf-8")
	fmt.Fprintf(w, plist, email, email, email, email)
}

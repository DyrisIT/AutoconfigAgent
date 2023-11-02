package cli

var config = map[string]string{
	// E-Mail Provider Settings
	"PROVIDER_NAME": "KiloHost",

	// SMTP
	"IMAP_SUBDOMAIN": "imap",               // imap or pop3
	"IMAP_PORT":      "993",                // 993: IMAP 995: POP3
	"IMAP_SECURITY":  "SSL",                // plain, SSL or STARTTLS
	"IMAP_AUTH":      "password-cleartext", // password-cleartext, ..., OAuth2, none

	// IMAP (preferred) or POP3
	"SMTP_SUBDOMAIN": "smtp",               // smtp
	"SMTP_PORT":      "465",                // 465: implicit TLS 587: plain or STARTTLS
	"SMTP_SECURITY":  "SSL",                // plain, SSL or STARTTLS
	"SMTP_AUTH":      "password-cleartext", // password-cleartext, none (should never be used)
}

func Get(key string) string {
	return config[key]
}

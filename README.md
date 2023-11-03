# AutoconfigAgent
Simple API, parsing an email address into autoconfig XML files for Thunderbird, Outlook and iOS Mail

## Notes
This requires Postfix & Dovecot SNI if used on a single server with multiple different email domains, since the server hostname for imap and smtp will be different and encryption should always be used.

## Build
```bash
go mod download
go build
```

## Usage
```bash
./AutoconfigAgent -b 0.0.0.0 -p 1234
```
This starts the API bound to a specific IP address and port.
### Systemd Service
Replace `<user>` with a non-root user under which the API should be run. 
Make sure the executable has the correct owner/permissions and make sure you set the correct path under `ExecStart`. 
Save this file under `/etc/systemd/system/autoconfig-agent.service`, run `systemctl daemon-reload` and enable and run it via `systemctl enable autoconfig-agent; systemctl start autoconfig-agent`.
```ini
[Unit]
Description=AutoconfigAgent
After=network-online.target

[Service]
Type=simple
Restart=always
User=<user>
ExecStart=/path/to/AutoconfigAgent -b 0.0.0.0 -p 1234

[Install]
WantedBy=multi-user.target
```
### Nginx Config (SSL is optional, but recommended)
For every email domain in use, point the corresponding `@`, `autoconfig.` and `autodiscover.` records to the IP of the nginx server. 
The following is an example config, repalace any `<domain>` placeholder with your email domain.
```nginx
upstream autoconfig {
	# Backend Server where the AutoconfigAgent API is hosted
	server 127.0.0.1:1234;
}
server {
	listen 80;
	listen [::]:80;

	server_name <domain>;

	if ($scheme = http) {
		return 301 https://$host$request_uri;
	}
}
server {
	#listen 443 ssl http2;
	listen 443 ssl;
	listen [::]:443 ssl;

	server_name <domain>;

	ssl_certificate /etc/letsencrypt/live/<domain>/fullchain.pem;
	ssl_certificate_key /etc/letsencrypt/live/<domain>/privkey.pem;
	include /etc/letsencrypt/options-ssl-nginx.conf;
	ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;

	root /var/www/<domain>;
	index index.html index.htm index.nginx-debian.html;
	charset utf8;

	# iOS E-Mail Autoconfig
	location = /.well-known/mobileconfig {
		proxy_pass http://autoconfig;
		proxy_set_header Host $host;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header X-Forwarded-Proto $scheme;
	}

	location / {
		try_files $uri $uri/ =404;
	}
}

server {
	#listen 443 ssl http2;
	listen 443 ssl;
	listen [::]:443 ssl;

	server_name autoconfig.<domain> autodiscover.<domain>;

	ssl_certificate /etc/letsencrypt/live/<domain>/fullchain.pem;
	ssl_certificate_key /etc/letsencrypt/live/<domain>/privkey.pem;
	include /etc/letsencrypt/options-ssl-nginx.conf;
	ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;

	root /var/www/<domain>;
	index index.html index.htm index.nginx-debian.html;
	charset utf8;

	# Thunderbird E-Mail Autoconfig
	location = /mail/config-v1.1.xml {
		proxy_pass http://autoconfig;
		proxy_set_header Host $host;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header X-Forwarded-Proto $scheme;
	}

	# Outlook E-Mail Autoconfig
	location = /autodiscover/autodiscover.xml {
		proxy_pass http://autoconfig;
		proxy_set_header Host $host;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header X-Forwarded-Proto $scheme;
	}

	location / {
		try_files $uri $uri/ =404;
	}
}
```

## TODO:
- [ ] Add support for CalDAV/CardDAV autoconfig (Radical or Baikal Server required)
- [ ] Research into more autoconfig methods (e.g. K-9 Mail uses SRV Records, so no XML API is necessary)

## Contribution
Feel free to create a pull request with your improvements.

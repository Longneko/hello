server {
	listen 80 default_server;
	listen [::]:80 default_server;

	location / {
		proxy_pass http://HELLO_SERVER_HOST:HELLO_SERVER_PORT; # proxy to Hello GoLang server
	}

	# Statics:
	location ~ \.(gif|jpg|png)$ {
	# TODO: fill from template based on porject's actual directory
		root HELLO_DIR/static;
	}
}

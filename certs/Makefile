all:
	openssl genrsa -out server.key 2048
	openssl req -new -x509 -key server.key -out server.pem -days 3650

check:
	openssl x509 -in server.pem  --noout --text
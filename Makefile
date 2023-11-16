.PHONY:

cert:
	@echo Generating SSL certificates
	cd ./ssl && sh instructions.sh
	cd ..
	cd ./nginx && openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout nginx-selfsigned.key -out nginx-selfsigned.crt

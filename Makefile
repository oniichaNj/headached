all:
	go build .

install: 
	cp -f headached /usr/sbin/headached
	cp -f headached.json /etc/headached.json
	cp -f headached.json.example /etc/headached.json.example

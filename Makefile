all:
	go build .

install: 
	cp -f headached /usr/sbin/headached
	cp -f headached.json /etc/

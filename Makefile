all:
	@go build .

install: 
	@cp -f headached /usr/sbin/headached
	@cp -f headached.json /etc/headached.json
	@cp -f headached.json.example /etc/headached.json.example

systemd-install: install
	@cp -f init/headached.service /usr/lib/systemd/system/headached.service
	@echo "Unit file installed. Run `systemctl enable headached.service` to enable it. "

gentoo-install: install
	@cp -f init/gentoo-headached.sh /etc/init.d/headached
	@chmod +x /etc/init.d/headached
	@echo "Init script installed. Run `rc-update add headached <runlevel>` to enable it. "

debian-install: install
	@cp -f init/debian-headached.sh /etc/init.d/headached
	@chmod +x /etc/init.d/headached
	@echo "Init script installed. Run `update-rc.d headached defaults` to enable it. "

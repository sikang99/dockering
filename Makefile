#
# Docker listing
#
all: usage

info:
	@docker version

install:
	curl -s https://get.docker.io/ubuntu/ | sudo sh
	
group:
	sudo groupadd docker
	sudo gpasswd -a ${USER} docker
	sudo service docker restart

run r:
	@httpfile . &

kill k:
	@killall httpfile

view v:
	@chromium-browser http://localhost:8080/

readme md:
	vi README.md

template mt:
	vi Makefile.template

make m:
	vi Makefile

usage:
	@echo ""
	@echo "usage: make [install|info|group|run|kill|view]"
	@echo ""


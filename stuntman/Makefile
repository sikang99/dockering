#
# stuntman - stun client and server from stunprotocol.org
# by Stoney Kang, sikang99@gmail.com
#
IMAGE=stoney/stuntman:0.5
NAME=stuntman

OPT_P=-p 3478:3478/udp -p 3478:3478/tcp
OPTION=$(OPT_P)

#DOCKER=Dockerfile.alpine
DOCKER=Dockerfile.ubuntu
#-----------------------------------------------------------------------------
all: usage

edit e:
	vi $(DOCKER)

build b:
	docker build -f Dockerfile.$(LINUX) -t $(IMAGE) .

#-----------------------------------------------------------------------------
run r:
	@echo ""
	@echo "make (run) [rd(aemon)|rs(hell)|r(client)|rt(est)]"
	@echo ""

rd:
	docker run -d -it $(OPTION) --name=$(NAME) $(IMAGE)
rs:
	docker run -it --rm $(OPTION) --name=$(NAME) $(IMAGE) sh
rc:
	docker exec $(NAME) ./stunclient --verbosity 2 stun.stunprotocol.org
rt:
	docker exec -it $(NAME) /bin/bash

#-----------------------------------------------------------------------------
list l:
	@docker images $(IMAGE)
	@docker ps

kill k:
	@echo ""
	@echo "Stopping ... `docker stop $(NAME)`"
	@echo "Removing ... `docker rm $(NAME)`"
	@echo ""

clean:
	docker rmi -f $(IMAGE)

md:
	vi README.md

make m:
	vi Makefile

usage:
	@echo ""
	@echo "usage: make [e(dit)|b(uild)|r(un)|l(ist)|k(ill)]"
	@echo ""


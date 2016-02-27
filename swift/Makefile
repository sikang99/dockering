#
# Docker image for swiftlang
#
NAME=swiftfun
IMAGE=swiftdocker/swift

all: usage

edit e:

build b:

pull p:
	@docker pull $(IMAGE)

list l:
	@docker images $(IMAGE)
	@docker ps -f 'name=$(NAME)'

start:
	@docker start $(NAME)
	@docker attach $(NAME)

run r:
	@docker run -it --rm -v $(HOME)/coding/swift:/home --name $(NAME) $(IMAGE):latest /bin/bash

make m:
	@vi Makefile

usage:
	@echo ""
	@echo "Docker handling for '$(IMAGE)'"
	@echo "usage: make [edit|build|pull|list|run]"
	@echo ""

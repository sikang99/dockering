#
# Handling docker for simple go program
#
TARGET=app
#TARGET=hello
#TARGET=server
DOCKER=Dockerfile.$(TARGET)
IMAGE= stoney/go.$(TARGET)

all: usage

edit e:
	vi $(DOCKER)

es:
	vi $(TARGET).go

pull p:
	@docker pull golang

build b:
	@docker build -t $(IMAGE) -f $(DOCKER) .
	@docker images $(IMAGE)

list l:
	@docker images $(IMAGE)
	@docker ps -f 'image=$(IMAGE)'

run r:
	@docker run --rm -p 4444:4444 $(IMAGE) &

ri:
	@docker run -it --rm -p 4444:4444 $(IMAGE) /bin/bash

kill k:
	@for image in `docker ps -q -f 'image=$(IMAGE)'`; do \
		echo Stopping `docker stop $$image`; \
	done;
	@docker ps -f 'image=$(IMAGE)'

clean c:
	@for image in `docker images -q $(IMAGE)`; do \
		docker rmi -f $$image; \
	done;
	@docker images $(IMAGE)

make m:
	vi Makefile

usage:
	@echo ""
	@echo "Docker handling for '$(IMAGE)' by Stoney Kang"
	@echo "usage: make [e(dit)|p(ull)|b(uild)|l(ist)|r(un)|k(ill)|c(lean)]"
	@echo ""

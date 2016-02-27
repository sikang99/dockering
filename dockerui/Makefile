.PHONY: all edit run kill make usage

NAME=dockerui
IMAGE=$(NAME)/$(NAME)

all: usage
#---------------------------------------------------------------------
pull p:
	@docker pull $(IMAGE)

list l:
	@docker images $(IMAGE)
	@docker ps

run r:
	@docker run -d -p 9000:9000 --privileged -v /var/run/docker.sock:/var/run/docker.sock $(IMAGE)
	@docker ps

view v:
	@chromium-browser http://localhost:9000/
	
kill k:
	@for image in `docker ps -q -f 'image=$(IMAGE)'`; do \
		echo Stopping `docker stop $$image`; \
	done;
	@docker ps -f 'image=$(IMAGE)'

#---------------------------------------------------------------------
make m:
	vi Makefile

usage:
	@echo ""
	@echo "Docker handling for '$(IMAGE)' by Stoney Kang"
	@echo "make [pull|list|run|view|kill]"
	@echo ""

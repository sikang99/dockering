#
# https://github.com/docker-library/official-images/
# https://hub.docker.com/_/postgres/
# Offiial Image : 9.5 lastest
#
.PHONY: all pkg run kill make usage

NAME=postgres
IMAGE=$(NAME)
PORT=5432

all: usage
#---------------------------------------------------------------------
pkg p:
	@echo ""
	@echo "Select run an option for ($@) for '$(IMAGE)'"
	@echo "\tmake (pkg) [pi(nstall)|pc(lean)]"
	@echo ""

pi:
	@docker pull $(IMAGE)

pc:
	@docker rmi $(IMAGE)

list l:
	@docker images $(IMAGE)
	@docker ps -f 'name=$(IMAGE)'

run r:
	@echo "Select run an option for ($@)"
	@echo "\tmake (run) [rd|ri]"

rd:
	@docker run -d -p $(PORT):$(PORT) -e POSTGRES_PASSWORD=root --name $(NAME) $(IMAGE)
	@docker ps -f 'name=$(IMAGE)'

ri:
	@docker run -it --link $(NAME):pg --rm $(IMAGE) bash

view v:
	@chromium-browser http://localhost:$(PORT) 
	
kill k:
	@echo ""
	@echo "Select run an option for ($@) for '$(IMAGE)'"
	@echo "\t> make (kill) [kd]"
	@echo ""

kd:
	@docker stop $(NAME)
	@docker rm $(NAME)

#---------------------------------------------------------------------
make m:
	vi Makefile

usage:
	@echo ""
	@echo "Docker handling for '$(IMAGE)' by Stoney Kang"
	@echo "\tmake [pkg|list|run|view|kill]"
	@echo ""

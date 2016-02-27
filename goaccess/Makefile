#
#
#
all: usage

NAME=goaccess
IMAGE=diyan/$(NAME)
#---------------------------------------------------------------------------------
docker d:
	@echo ""
	@echo "make (docker) [dp|ds|di]"
	@echo ""

ds:
	@docker search $(IMAGE)
dp:
	@docker pull $(IMAGE)
di:
	@docker inspect $(IMAGE)

#---------------------------------------------------------------------------------
md:
	vi README.md

make m:
	vi Makefile

usage:
	@echo ""
	@echo "usage: make [docker|run]"
	@echo ""

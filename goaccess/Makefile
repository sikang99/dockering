#
# Makefile for goaccess docker
#
all: usage

NAME=goaccess
IMAGE=diyan/$(NAME)
#---------------------------------------------------------------------------------
docker d:
	@echo ""
	@echo "make (docker) [dp(ull)|ds(earch)|di(nspect)|dr(un)|dt(est)]"
	@echo ""

ds:
	@docker search $(IMAGE)
dp:
	@docker pull $(IMAGE)
di:
	@docker inspect $(IMAGE)
dr:
	@docker run --rm $(IMAGE) --help

dt:
	cat access.log | docker run --rm -i diyan/goaccess \
	--time-format='%H:%M:%S' \
	--date-format='%d/%b/%Y' \
	--log-format='%h %^[%d:%t %^] "%r" %s %b "%R" "%u"' > access.html

#---------------------------------------------------------------------------------
md:
	vi README.md

make m:
	vi Makefile

usage:
	@echo ""
	@echo "usage: make [docker|run]"
	@echo ""

#
# Makefile template for docker handling
#
all: usage

NAME=nifi
IMAGE=mkobit/nifi
TAG=0.1
#---------------------------------------------------------------------------------
docker d:
	@echo ""
	@echo "make (docker) [db(uild)|dp(ull)|ds(earch)|di(nspect)|dc(ommit)|dr(un)]"
	@echo ""

db:
	@docker build -f Dockerfile -t $(IMAGE):$(TAG) .
ds:
	@docker search $(NAME)
dp:
	@docker pull $(IMAGE)
dc:
	@docker commit $(IMAGE):$(TAG)
dr:
	@docker run --rm -it --name $(NAME) $(IMAGE) /bin/bash

#---------------------------------------------------------------------------------
run r:
	@echo ""
	@echo "make (run) [rd(aemon)|rt(erm)|rs(hell)|re(xec)|rw(eb)]"
	@echo ""

NIFI_HOME=/opt/nifi
P_OPT=-p 8080-8081:8080-8081
V_OPT=-v ${PWD}/database:$(NIFI_HOME)/database_repository -v ${PWD}/content:$(NIFI_HOME)/content_repository \
-v ${PWD}/flowfile:$(NIFI_HOME)/flowfile_repository -v ${PWD}/provenance:$(NIFI_HOME)/provenance_repository
E_OPT=
OPTION=$(P_OPT) $(V_OPT) $(E_OPT)

rd:
	docker run -d $(OPTION) --name $(NAME) $(IMAGE)
rt:
	@docker run -it $(OPTION) --name $(NAME) $(IMAGE)
rs:
	@docker run --rm -it $(OPTION) --name $(NAME) /bin/bash
re:
	@docker exec -it $(NAME) /bin/bash
rw:
	@chromium-browser http://localhost:8080/nifi/ &	

#---------------------------------------------------------------------------------
kill k:
	@echo ""
	@echo "make (kill) [kc(ontainer)]"
	@echo ""

kc:
	@docker stop $(NAME)
	@docker rm $(NAME)

#---------------------------------------------------------------------------------
command c:
	@echo ""
	@echo "make (command) [inspect|search|logs|info|ps]"
	@echo ""

search:
	@docker search $(NAME)
inspect:
	@docker inspect $(IMAGE)
logs:
	@docker logs $(NAME)
exec:
	@docker exec -it $(NAME) /bin/bash
info:
	@docker info
ps:
	@docker ps

#---------------------------------------------------------------------------------
md:
	vi README.md
make m:
	vi Makefile

usage:
	@echo ""
	@echo "Docker Handling for '$(IMAGE)' with '$(NAME)'"
	@echo "usage: make [docker|run|kill]"
	@echo ""


#
# Apache Drill
#
all: usage

IMAGE=sequenceiq/drill
NAME=drill

OPT_P=-p 8047:8047
OPT_V=-v /home/stoney/coding/docker/drill/data:/data
OPT_E=
OPTION=$(OPT_P) $(OPT_V) $(OPT_E)
#---------------------------------------------------------------------------------
docker d:
	@echo ""
	@echo "make ($@) [ds(earch)|dp(ull)|di(nfo)|dl(list)]"
	@echo ""

ds:
	@docker search $(NAME)
dp:
	@docker pull $(IMAGE)
di:
	@docker inspect $(IMAGE) | grep IPAddr
dl:
	@docker ps -f name=$(NAME)
dr:
	@docker run -it $(OPTION) --name=$(NAME) $(IMAGE) /etc/bootstrap.sh
dh:
	@docker run --rm -it $(OPTION) $(IMAGE) sh
dq:
	@docker stop $(NAME)
	@docker rm $(NAME)

#---------------------------------------------------------------------------------
api a:
	@echo ""
	@echo "make ($@) [as(tatus)|aq(uery)]"
	@echo ""

as:
	@chromium-browser http://localhost:8047/status
aq:
	@chromium-browser http://localhost:8047/queries
#---------------------------------------------------------------------------------
md:
	vi README.md

make m:
	vi Makefile

usage:
	@echo ""
	@echo "Docker Handling for '$(IMAGE)'"
	@echo "usage: make [docker|run]"
	@echo ""

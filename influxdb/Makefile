# Microservice in Go
# Handling docker for influxdb
#
NAME=influxdb
DOCKER=Dockerfile.$(NAME)
IMAGE=tutum/influxdb
PORT=8083

all: usage

edit e:
	vi $(DOCKER)

#-------------------------------------------------------------------------------
install i:
	@echo ""
	@echo "\t> make ($@) [ip(ackage)|iu(tils)|id(ocker)]"
	@echo ""

ip:
	wget https://s3.amazonaws.com/influxdb/influxdb_0.10.0-1_amd64.deb
	sudo dpkg -i influxdb_0.10.0-1_amd64.deb
	sudo service influxdb start

iu:
	go get github.com/rakyll/boom
	go get github.com/tsenart/vegeta

id:
	@docker pull $(IMAGE)

#------------------------------------------------------------------------------
list l:
	@echo ""
	@echo "\t> make ($@) [li(mage)|ll(og)|ls|ln]"
	@echo ""

li:
	@docker images $(IMAGE)
	@docker ps -f 'name=$(NAME)'

ll:
	@docker logs $(NAME)

ls:
	@docker inspect $(NAME)

ln:
	@ifconfig docker0

#------------------------------------------------------------------------------
run r:
	@echo ""
	@echo "\t> make ($@) [rd(aemon)|rs(hell)|re(xec)]"
	@echo ""

OPT_P=-p 8083:8083 -p 8086:8086 
OPT_X=--expose 8090 --expose 8099
OPT_V=-v /home/stoney/coding/docker/influxdb/data:/data 
OPT_E=-e ADMIN_USER="root" -e FORCE_HOSTNAME="localhost" -e PRE_CREATE_DB="metric"
OPTION=$(OPT_P) $(OPT_V) $(OPT_E)

rd:
	docker run -d $(OPTION) --name $(NAME) $(IMAGE)

rs:
	@docker run -it --rm $(IMAGE) /bin/bash # influx shell

rf:
	@docker exec -it $(NAME) influx

re:
	@docker exec -it $(NAME) /bin/bash

rt:
	curl -G http://localhost:8086/query --data-urlencode "q=CREATE DATABASE mydb"
	curl -i -XPOST 'http://localhost:8086/write?db=mydb' --data-binary 'cpu_load_short,host=server01,region=us-west value=0.64 1434055562000000000'

#------------------------------------------------------------------------------
start:
	@docker start $(NAME)

stop s:
	@for image in `docker ps -q -f 'image=$(IMAGE)'`; do \
		echo Stopping `docker stop $$image`; \
	done;
	@docker ps -f 'image=$(IMAGE)'

clean c:
	@for image in `docker images -q $(IMAGE)`; do \
		docker rmi -f $$image; \
	done;
	@docker images $(IMAGE)


#------------------------------------------------------------------------------
kill k:
	@echo ""
	@echo "make (kill) [kd(aemon)|ki(mage)]"
	@echo ""

kd:
	@docker stop $(NAME)
	@docker rm $(NAME)

ki:
	@docker rmi $(IMAGE)

# ----------------------------------------------------------------------------------------------------
go g:
	@echo ""
	@echo "\t> make ($@) [ge(dit)|gr(run)]"
	@echo ""

SOURCE=metric.go

ge:
	vi $(SOURCE)

gr:
	go run $(SOURCE)

# ----------------------------------------------------------------------------------------------------
md:
	vi README.md

make m:
	vi Makefile

usage:
	@echo ""
	@echo "Docker handling for $(IMAGE) by Stoney Kang"
	@echo "usage: make [i(nstall)|l(ist)|r(un)|k(ill)]"
	@echo ""

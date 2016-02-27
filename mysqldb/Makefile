#
# MYSQL docker
#
.PHONY: all pkg list run kill make usage

VERSION=:5.7#latest
IMAGE=mysql
NAME=mysql
PORT=3306

OPT_P=-p 3306:3306
OPT_V=-v ${PWD}/data:/var/lib/mysql
OPT_E=-e MYSQL_ROOT_PASSWORD=root 
OPTION=$(OPT_P) $(OPT_V) $(OPT_E)

all: usage
#---------------------------------------------------------------------
list l:
	@echo ""
	@echo "Select an option of ($@) for '$(IMAGE)'"
	@echo "\t> make (list) [li(mage)|ll(og)|lp(s)]"
	@echo ""

li:; @docker images $(IMAGE)
lp:; @docker ps -f 'name=$(NAME)'
ll:; @docker logs $(NAME)

view v:
	@chromium-browser http://localhost:$(PORT) 
	
#---------------------------------------------------------------------
pkg p:
	@echo ""
	@echo "Select an option of ($@) for '$(IMAGE)'"
	@echo "\t> make (pkg) [ps(earch)|pp(ull)|pc(lean)]"
	@echo ""

ps:; @docker search $(IMAGE)
pp:; @docker pull $(IMAGE)$(VERSION)
pc:; @docker rmi -f $(IMAGE)
pi:; @docker images $(IMAGE)

#---------------------------------------------------------------------
run r:
	@echo ""
	@echo "Select an option of ($@) for '$(IMAGE)$(VERSION)'"
	@echo "\t> make (run) [rd(aemon)|rs(hell)|re(xec)|ri(nspect)]"
	@echo ""

rd:; @docker run -d $(OPTION) --name $(NAME) $(IMAGE)$(VERSION)
rs:; @docker run -it --rm --link $(NAME):mysql $(IMAGE) bash
re:; @docker exec -it $(NAME) bash
ri:; @docker inspect $(NAME) | grep IP 

#---------------------------------------------------------------------
kill k:
	@echo ""
	@echo "Select an option of ($@) for '$(IMAGE)'"
	@echo "\t> make (kill) [kd(own)|ki(mage)]"
	@echo ""

kd:
	@docker stop $(NAME)
	@docker rm $(NAME)
kk:
	@docker stop $(NAME)
ki:
	@docker rmi -f $(IMAGE)$(VERSION)

#---------------------------------------------------------------------
go g:
	@echo ""
	@echo "Select run an option of ($@) for '$(IMAGE)'"
	@echo "\t> make ($@) [ge(dit)|gr(un)]"
	@echo ""

ge:; vi main.go
gr:; go run main.go

#---------------------------------------------------------------------
readme md:
	vi README.md

make m:
	vi Makefile

usage:
	@echo ""
	@echo "Docker handling for '$(IMAGE)$(VERSION)' by Stoney Kang"
	@echo "\t> make [list|pkg|run|view|kill]"
	@echo ""


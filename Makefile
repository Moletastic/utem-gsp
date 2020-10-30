# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=bin/gsp

DBNAME=utem-gsp.db
DBNAMEBAK=utem-gsp-test.db

MYSQLDB=utem_gsp

all: run

run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)

log:
	grc tail -f $(logfile)

watch:
	fresh

test:
	$(GOTEST) -v test/*.go -count=1

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

copy:
	cp $(DBNAME) $(DBNAMEBAK)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

front-sqlite:
	sqlitebrowser $(DBNAME)

front-mysql:
	mysql-workbench

refresh-mysql:
	mysql -u root -p -e "drop database $(MYSQLDB); create database $(MYSQLDB);"

.PHONY: test

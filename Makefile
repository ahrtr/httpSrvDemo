.PHONY : container 

TAG ?= `git rev-parse --short HEAD`

all: container

container:
	echo "Building httpsrvdemo"
	docker build -t httpsrvdemo:${TAG} .

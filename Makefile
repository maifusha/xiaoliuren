.PHONY: all gotool build test run clean image container

APP=xiaoliuren

all: gotool build test run

gotool:
	go mod tidy
	go fmt ./
	go vet ./

build:
	CGO_ENABLED=1 go build -ldflags "-w -s" -o ${APP} -v

test:
	go test -v

run:
	./${APP}

clean:
	@if [ -f ${APP} ] ; then rm -rf ${APP} ; fi

image:
	docker login -u maifusha hub.docker.com
	docker build -t maifusha/xiaoliuren -f docker/Dockerfile .
	docker push maifusha/xiaoliuren

container:
	docker run -d --rm -p 8000:8000 maifusha/xiaoliuren

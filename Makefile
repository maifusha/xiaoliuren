.PHONY: default build run gotool image container clean
APP=xiaoliuren

default: build run

build:
	CGO_ENABLED=1 go build -ldflags "-w -s" -o ${APP} -v

run:
	./${APP}

gotool:
	go mod tidy
	go vet ./...
	go fmt ./...

image:
	docker login --username $$(head -n 1 credential.docker) --password $$(tail -n 1 credential.docker)
	version=$$(date +%Y%m%d%H%M%S)-$$(git rev-parse --short HEAD);\
	docker build -t maifusha/xiaoliuren:$$version -f docker/Dockerfile .;\
	docker push maifusha/xiaoliuren:$$version;\
	docker tag maifusha/xiaoliuren:$$version maifusha/xiaoliuren:latest;\
	docker push maifusha/xiaoliuren:latest

container:
	docker run -d --rm -p 8000:8000 maifusha/xiaoliuren:latest

clean:
	go clean -i -x
	docker ps -qf ancestor=maifusha/xiaoliuren | xargs docker rm -fv
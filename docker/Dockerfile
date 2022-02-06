FROM golang:1.17-alpine3.15 as builder
LABEL maintainer="lixin <1045909037@qq.com>"

ENV CGO_ENABLED=1
ENV GOPROXY=https://goproxy.cn

COPY . /go/src
WORKDIR /go/src

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g" /etc/apk/repositories && apk update
RUN apk add --no-cache gcc libc-dev

RUN go mod download
RUN go build -ldflags "-w -s" -o xiaoliuren -v

########################################################################################################################

FROM alpine:3.15
LABEL maintainer="lixin <1045909037@qq.com>"

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g" /etc/apk/repositories && apk update
RUN apk add --no-cache tzdata && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=builder /go/src/xiaoliuren /svc/xiaoliuren
COPY sqlite.db /svc/sqlite.db
WORKDIR /svc

ENV GOPATH=/var/www
VOLUME $GOPATH/log

EXPOSE 8000

COPY docker/entrypoint.sh /svc/entrypoint.sh
COPY docker/scripts /svc/scripts
RUN chmod -R +x /svc/entrypoint.sh /svc/scripts && rm /svc/scripts/.gitkeep
ENTRYPOINT ["./entrypoint.sh"]
CMD ["./xiaoliuren"]

HEALTHCHECK --interval=5s --timeout=3s \
  CMD curl --fail http://localhost/ || exit 1
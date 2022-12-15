FROM golang:1.17-alpine

RUN apk add --no-cache gcc musl-dev

ENV SRC_DIR=/go/src/app
ENV GOBIN=/go/bin

WORKDIR ${SRC_DIR}
ADD ./gin-app ${SRC_DIR}
RUN cd ${SRC_DIR};

# バージョンを上げるには `golang:1.18` 以降が必要
RUN go install github.com/go-delve/delve/cmd/dlv@v1.9.1

RUN go install github.com/cosmtrek/air@latest

# goパッケージのインストール
RUN ["go", "mod", "tidy"]

ENTRYPOINT ["air", "-c", ".air.toml"]
FROM golang:1.17-alpine

ENV SRC_DIR=/go/src/app

ENV GOBIN=/go/bin

WORKDIR ${SRC_DIR}

ADD ./gin-app ${SRC_DIR}

RUN cd ${SRC_DIR};

# goパッケージのインストール
RUN ["go", "mod", "tidy"]

ENTRYPOINT ["go", "run", "main.go"]
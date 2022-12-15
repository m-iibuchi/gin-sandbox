FROM golang:1.19-alpine

ENV SRC_DIR=/go/src/app

ENV GOBIN=/go/bin

WORKDIR ${SRC_DIR}

RUN cd ${SRC_DIR};

# RUN go get github.com/go-sql-driver/mysql \
#     && go get -u github.com/gin-gonic/gin \
#     && go get github.com/gorilla/mux \
#     && go get -u github.com/jinzhu/gorm \
#     && go get github.com/gin-contrib/cors \
#     && go get gopkg.in/ini.v1

# ENTRYPOINT ["go", "run", "main.go"]
FROM golang:1.11.2-alpine
WORKDIR /go/src/api
ADD . /go/src/api
RUN 	apk add git \
	&& cd /go/src/api \
	&& go get github.com/gorilla/mux  \
	&& go get go.mongodb.org/mongo-driver/mongo \
	&& go get go.mongodb.org/mongo-driver/mongo/options \
	&& go build
EXPOSE 3000
ENTRYPOINT ./api
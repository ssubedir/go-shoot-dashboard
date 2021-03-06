FROM golang:latest

EXPOSE 9000 9001

LABEL maintainer="ssubedir <ssubedir@gmail.com>"

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

CMD ["./go-shoot-dashboard"]
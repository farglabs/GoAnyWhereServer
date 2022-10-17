#build stage
FROM golang:bullseye AS builder
RUN apt update \
    && apt upgrade \
    && apt install -y sqlite3
WORKDIR /go/src/
COPY ./fileserver/src .
RUN echo pwd
RUN go get -d -v ./...
RUN  go build -o /go/bin -v ./...
RUN mkdir /home/fileserver
RUN touch /home/fileserver/index.txt
ENTRYPOINT /go/bin/fileserver
LABEL Name=fileserver Version=0.0.1
EXPOSE 55433
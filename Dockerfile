FROM golang:onbuild
MAINTAINER Vojtech Bartos <hi@vojtech.me>

RUN go get github.com/zenazn/goji

EXPOSE 8000

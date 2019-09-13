# API
FROM golang:1.12.5-stretch AS development

RUN go get github.com/pilu/fresh
WORKDIR /source/api

EXPOSE 80
ENTRYPOINT ["fresh"]

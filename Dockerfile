FROM golang:1.18-alpine as buildbase

ENV KV_VIPER_FILE=./config.yaml

RUN apk add git build-base

WORKDIR /go/src/blob
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/testService /go/src/blob


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/testService /usr/local/bin/testService
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["testService"]

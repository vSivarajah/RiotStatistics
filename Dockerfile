


FROM golang:1.14.2-alpine as builder
LABEL maintainer="Vignesh Sivarajah <vigi_03_6@hotmail.com>"

RUN apk add alpine-sdk
ENV RIOTAPI_KEY=RGAPI-e1301563-214b-4f94-bda5-c77d3d110e9c \
               KAFKA_BOOTSTRAP_SERVERS=localhost:9092
WORKDIR /go/app
COPY . /go/app
COPY go.mod /go/app
COPY go.sum /go/app
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o main -tags musl

FROM alpine:latest as runner
WORKDIR /root/
COPY --from=builder /go/app/ .
ENTRYPOINT /root/main
FROM golang:1.8-alpine3.5 as builder

# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# setting working directory
WORKDIR /go/src/app

# installing dependencies
RUN go get github.com/leonelquinteros/gorand
RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/sirupsen/logrus

COPY / /go/src/app/
RUN go build -o alerter

FROM alpine:3.5

RUN apk update && apk upgrade && \
    apk add --no-cache bash openssh curl

WORKDIR /go/src/app
COPY --from=builder /go/src/app/alerter /go/src/app/alerter

EXPOSE 10010

CMD ["./alerter"]
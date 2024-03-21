FROM golang:1.22-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/rarimo/rarime-passport-verifier
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/rarime-passport-verifier /go/src/github.com/rarimo/rarime-passport-verifier


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/rarime-passport-verifier /usr/local/bin/rarime-passport-verifier
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["rarime-passport-verifier"]

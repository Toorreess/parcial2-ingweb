FROM golang:alpine AS build
WORKDIR /go/src/app
COPY . .
RUN export CGO_ENABLED=0 && go build -o /go/bin/app ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /go/bin/app /go/bin/app
WORKDIR /go/src/app

COPY config/config.yml config/config.yml
COPY creds.json creds.json

ENV GOPATH=/go
ENV GOOGLE_APPLICATION_CREDENTIALS=./creds.json

ENTRYPOINT ["/go/bin/app"]
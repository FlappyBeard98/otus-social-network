# First stage: Build the binary
FROM golang:alpine AS build-env

ARG service_dir

WORKDIR /go/src/app
COPY go.mod go.sum ./
COPY lib/ lib/
COPY vendor/ vendor/
COPY $service_dir $service_dir

RUN go build -o service $service_dir/main.go

# Second stage: Use a smaller image
FROM alpine:latest

COPY --from=build-env /go/src/app/service /usr/local/bin/

CMD ["service"]

#example for build profile service: `docker build --build-arg service_dir=profile -t profile .`
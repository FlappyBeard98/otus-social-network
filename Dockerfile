# First stage: Build the binary
FROM golang:alpine AS build-env

ARG service_dir

WORKDIR /go/src/app
COPY go.mod go.sum ./
COPY lib/ lib/
COPY vendor/ vendor/
COPY services/$service_dir services/$service_dir
COPY services/$service_dir/config.json ./

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseVendor --parseInternal --parseDependency -g services/$service_dir/main.go
RUN go build -o service services/$service_dir/main.go

# Second stage: Use a smaller image
FROM scratch

COPY --from=build-env /go/src/app/service /usr/local/bin/
COPY --from=build-env /go/src/app/config.json /usr/local/bin/

CMD ["service"]

#example for build profile service: `docker build --build-arg service_dir=profile -t profile -f . /services`


FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

# Export necessary port
EXPOSE 1323


CMD  ["./social-network"]
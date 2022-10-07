FROM golang:1.18.3-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go install gotest.tools/gotestsum@latest
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/swag/cmd/swag

RUN go env -w CGO_ENABLED=0
RUN go env -w GOOS=linux
RUN go env -w GOARCH=amd64
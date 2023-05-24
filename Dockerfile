FROM golang:1.20.3-alpine3.17
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
CMD ["go","run","server.go"]

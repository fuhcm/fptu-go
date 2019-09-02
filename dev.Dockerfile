FROM golang:alpine
RUN apk add ca-certificates git

RUN mkdir -p /root/src/go
WORKDIR /root/src/go
COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 5000

ENTRYPOINT ["go","run","cmd/fptugo/main.go"]
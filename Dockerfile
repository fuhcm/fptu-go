FROM golang:alpine as builder
RUN apk add ca-certificates git

RUN mkdir -p /root/src/go
WORKDIR /root/src/go

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

FROM alpine

RUN apk add ca-certificates rsync openssh

WORKDIR /root/src/go

COPY --from=builder /root/src/go/main /root/src/go/main

EXPOSE 5000

ENTRYPOINT ["./main"]
FROM golang:1.20-alpine as builder

WORKDIR /app
COPY go.* ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

COPY pkg ./pkg
COPY main.go ./
RUN go build -o server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server ./

ENTRYPOINT ["./server"]
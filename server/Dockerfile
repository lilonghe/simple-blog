FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .  

RUN go mod tidy && go build -o server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app ./

ENTRYPOINT ["./server"]
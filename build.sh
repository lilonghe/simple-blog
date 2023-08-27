# mac build, you need install upx
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o simple-blog -ldflags '-w -s' main.go && upx simple-blog
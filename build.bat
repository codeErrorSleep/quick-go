echo windows下的执行脚本
go env -w GOOS=linux
go env -w GOARCH=amd64
go build  -o ./bin/ ./...

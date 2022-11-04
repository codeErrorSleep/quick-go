echo windows下的执行脚本
go env -w GOOS=linux
go env -w GOARCH=amd64
protoc --go_out=./ --go-grpc_out=./ app/rpc/*.proto
go env -w GOOS=windows
@REM go build  -o ./bin/ ./...

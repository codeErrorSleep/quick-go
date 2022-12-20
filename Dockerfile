# 构建镜像使用

FROM golang as builder

WORKDIR /app

COPY . ./

# 加速！！！
ENV  GOPROXY=https://goproxy.cn,direct GO111MODULE=on

# 拉取依赖
RUN  go mod tidy

# Build
RUN  go build -o quick-go
# # 为了缩小镜像体积，做分层处理
FROM centos:7

WORKDIR /app

COPY --from=builder /app/ ./

# 启动命令，多行参数使用,隔开/data1/wwwroot/dev/logistics_go/logistics_go_dev/bin/logistics_go_server run -b /data1/wwwroot/dev/logistics_go/logistics_go_dev/ -e  .env.develop -p 26002
ENTRYPOINT ["./quick-go","run", "-e", "config-docker","-b","./"]

# 生成命令
# docker run -d --name quick-test --network host quick
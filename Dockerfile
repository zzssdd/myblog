FROM golang:1.18-alpine
#为镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0

#工作目录
WORKDIR /project/go-docker/

#下载依赖
COPY go.* ./
RUN go mod download
#编译
COPY . .
RUN go build -o /project/go-docker/build/myapp .

EXPOSE 8000
ENTRYPOINT ["/project/go-docker/build/myapp"]
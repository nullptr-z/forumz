# 第三方 golang 运行环境的镜像
FROM golang:alpine
# 定义环境变量
ENV GO111MODULE=on \
  GOPROXY=https://goproxy.cn \
  GOOS=linux \
  GOARCH=amd64
# 设置工作目录为 /build
WORKDIR /build
# 将当前目录代码内容复制到容器中
COPY . .
# SQL

# 构建
RUN go mod download
RUN go build -o app .
# 切换到目录,并把可执行文件复制过来
# WORKDIR /dist
# RUN cp /build/app .
# 安装 requirements.txt 中指定的任何所需包
# RUN go get package
# 申明暴露到外面的端口，没有实际作用
EXPOSE 8899
# 在容器启动时运行 app
CMD ["/build/app"]

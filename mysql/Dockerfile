# 使用官方MySQL 8.0镜像作为基础镜像
FROM mysql:latest

# 设置环境变量，以便在容器启动时初始化MySQL数据库
# 注意: 这些是示例值，请根据实际需要修改
ENV MYSQL_DATABASE=learn
ENV MYSQL_USER=zhou
ENV MYSQL_PASSWORD=zheng
ENV MYSQL_ROOT_PASSWORD=zhengmr

# 可选：如果你有自定义的SQL脚本需要在容器启动时执行，可以将它们添加到/docker-entrypoint-initdb.d目录
# COPY ./your-custom-scripts.sql /docker-entrypoint-initdb.d/

# 暴露MySQL端口
EXPOSE 3306

# 设置默认启动命令（实际上这一行不是必需的，因为基础镜像已经设置了默认的启动命令）
CMD ["mysqld"]

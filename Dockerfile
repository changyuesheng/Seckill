# FROM golang:alpine as builder

# ENV GOPROXY https://goproxy.cn
# ENV GO111MODULE on

# WORKDIR /go/cache
# ADD go.mod .
# ADD go.sum .
# RUN go mod download

# WORKDIR /go/release
# ADD . .
# RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o gin_demo main.go

# FROM scratch as prod
# COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# COPY --from=builder /go/release/gin_demo /
# COPY --from=builder /go/release/conf ./conf
# CMD ["/gin_demo"]

FROM golang:alpine as builder
RUN apk add build-base
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1

#设置时区参数
ENV TZ=Asia/Shanghai
RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories
RUN apk --no-cache add tzdata zeromq \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo '$TZ' > /etc/timezone

WORKDIR /app
COPY . /app

RUN go build .

EXPOSE 3000

ENTRYPOINT ["./test"]
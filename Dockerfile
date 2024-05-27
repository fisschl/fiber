FROM registry.cn-shanghai.aliyuncs.com/fisschl/golang:latest AS builder
WORKDIR /root
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main

FROM registry.cn-shanghai.aliyuncs.com/fisschl/golang:latest
WORKDIR /root
COPY --from=builder /root/main .
CMD ./main

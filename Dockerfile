FROM golang:1.18 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /ji
COPY . .
RUN go mod tidy \
    &&  cd /ji/cmd \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags="-w -s" -o main
RUN cd /ji \
    && mkdir publish -p  \
    && cd publish \
    && mkdir config -p  \
    && mkdir cmd -p  \
    && cp /ji/cmd/main /ji/publish/cmd  \
    && cp /ji/config/ji.conf /ji/publish/config

FROM busybox:1.28.4

WORKDIR /ji
COPY --from=builder /ji/publish .
ENV GIN_MODE=release
EXPOSE 3000

WORKDIR /ji/cmd
ENTRYPOINT ["./main"]
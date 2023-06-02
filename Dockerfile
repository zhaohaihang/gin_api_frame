FROM golang:1.18 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /gin_api_frame
COPY . .
RUN go mod tidy \
    &&  cd /gin_api_frame/cmd \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags="-w -s" -o main
RUN cd /gin_api_frame \
    && mkdir publish -p  \
    && cd publish \
    && mkdir config -p  \
    && mkdir cmd -p  \
    && cp /gin_api_frame/cmd/main /gin_api_frame/publish/cmd  \
    && cp /gin_api_frame/config/gin_api_frame.conf /gin_api_frame/publish/config

FROM busybox:1.28.4

WORKDIR /gin_api_frame
COPY --from=builder /gin_api_frame/publish .
ENV GIN_MODE=release
EXPOSE 3000

WORKDIR /gin_api_frame/cmd
ENTRYPOINT ["./main"]
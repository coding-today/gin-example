from golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/coder-stars/gin-example
COPY . $GOPATH/src/github.com/coder-stars/gin-example
RUN  go build .

EXPOSE 8000
ENTRYPOINT ["./gin-example"]
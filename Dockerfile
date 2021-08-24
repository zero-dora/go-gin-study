FROM scratch
ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/zero-dora/go-gin-example
COPY . $GOPATH/src/github.com/zero-dora/go-gin-example


EXPOSE 9090
ENTRYPOINT ["./go-gin-example"]
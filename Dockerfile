FROM golang:1.13 AS builder

RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /go/src/github.com/hublabs/auth-api
ADD go.mod go.sum ./
RUN go mod download
ADD . /go/src/github.com/hublabs/auth-api
ENV CGO_ENABLED=0
RUN go build -o auth-api

FROM pangpanglabs/alpine-ssl
WORKDIR /go/src/github.com/hublabs/auth-api
COPY --from=builder /go/src/github.com/hublabs/auth-api ./

EXPOSE 8003

CMD ["./auth-api", "api-server"]
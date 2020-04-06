FROM golang:1.13 as builder

ENV REPOSITORY github.dip-net.co.jp/ku-takeyama/grpc_test_api
ARG WORKDIR=$GOPATH/src/${REPOSITORY}
WORKDIR ${WORKDIR}

ENV GOSUMDB off
ENV GO111MODULE on
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

COPY . .
RUN cp -p build/package/golang/.gitconfig ~/.gitconfig
RUN go build -o /main $WORKDIR/cmd/api/grpc.go

# FROM alpine:3.11
FROM golang:1.13

ENV GOPATH /
ENV GOENV dev
COPY --from=builder /main /main
COPY  configs /configs

RUN chmod +x /main
ENTRYPOINT [ "/main" ]
EXPOSE 80
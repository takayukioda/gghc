FROM golang:1.8-alpine

# GOPATH configured in golang image is `/go`
ENV PROJPATH=/go/src/go-ghc

RUN set -eux; \
	apk update && \
	apk add git make
RUN go get github.com/golang/lint/golint

RUN mkdir -p $PROJPATH
WORKDIR $PROJPATH

COPY . $PROJPATH

CMD ["printenv"]

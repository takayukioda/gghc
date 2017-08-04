FROM golang:1.8-alpine

RUN go get github.com/golang/lint/golint

RUN mkdir -p /project/app
WORKDIR /project/app

COPY . .

CMD ["printenv"]

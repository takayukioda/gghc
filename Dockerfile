FROM golang:1.8-alpine

RUN mkdir -p /project/app
WORKDIR /project/app

COPY . .

CMD ["printenv"]

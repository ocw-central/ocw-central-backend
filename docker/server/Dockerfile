FROM golang:alpine3.16 as builder
WORKDIR /workspace

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main server.go wire_gen.go && chmod +x ./main

FROM alpine:3.15
WORKDIR /app

RUN wget -q -O - https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz \
    && ln -s /app/migrate.linux-amd64 /usr/bin/migrate

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

COPY --from=builder /workspace/main ./

COPY ./docker/server/deploy.sh /usr/local/bin
RUN chmod +x /usr/local/bin/deploy.sh

CMD ["deploy.sh"]

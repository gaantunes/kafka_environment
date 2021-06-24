FROM golang:alpine as builder

WORKDIR /app 

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o kafka-client . 

FROM scratch

WORKDIR /app

COPY --from=builder /app/kafka-client /app/kafka-client

ENTRYPOINT ["./kafka-client"]
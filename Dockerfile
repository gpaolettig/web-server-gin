FROM golang:1.24.1 AS builder
WORKDIR /app
COPY go.mod ./
COPY . ./
RUN go build -o web-serv .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/web-serv .
EXPOSE 8080
ENTRYPOINT ["./web-serv"]

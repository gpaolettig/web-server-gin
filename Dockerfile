FROM alpine:latest
WORKDIR /app
COPY web-serv .
EXPOSE 8080
ENTRYPOINT ["./web-serv"]



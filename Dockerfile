FROM golang:latest AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/myapp .


FROM alpine:latest AS runtime
COPY --from=build /app/myapp /myapp
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/myapp"]



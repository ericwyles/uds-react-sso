FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /go-jwt-app main.go


FROM gcr.io/distroless/base-debian12
WORKDIR /root/
COPY --from=builder /go-jwt-app /tmp
EXPOSE 8080
CMD ["/tmp/go-jwt-app"]
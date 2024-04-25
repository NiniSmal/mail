FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/mail ./main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/mail /app/mail
EXPOSE 8090
ENTRYPOINT ["/app/mail"]
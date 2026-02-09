FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wallet ./cmd/crm/main.go

FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/wallet /app/wallet

EXPOSE 8080

ENTRYPOINT [ "/app/wallet" ]
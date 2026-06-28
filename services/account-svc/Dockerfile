FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o account-svc ./cmd/main.go

FROM gcr.io/distroless/static-debian11
WORKDIR /app
COPY --from=builder /app/account-svc .
EXPOSE 8081
CMD ["./account-svc"]

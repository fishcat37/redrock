FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o myapp .
FROM alipine:latest
RUN apk add --no-cache libc6-compat
WORKDIR /root/
COPY --from=builder /app/myapp .
EXPOSE 8080
CMD ["./myapp"]
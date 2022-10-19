FROM golang:1.18-alpine as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./auth-banking ./auth-banking
WORKDIR /app/auth-banking
RUN go build -o /auth-banking-service
FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /auth-banking-service /app/auth-banking/auth-banking-service
EXPOSE 3001
CMD ["./auth-banking/auth-banking-service"]
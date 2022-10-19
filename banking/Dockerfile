FROM golang:1.18-alpine as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./banking ./banking
WORKDIR /app/banking
RUN go build -o /banking-service
FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /banking-service /app/banking/banking-service
EXPOSE 3000
CMD ["./banking/banking-service"]
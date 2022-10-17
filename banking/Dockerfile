# syntax=docker/dockerfile:1

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.18-alpine as builder

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./
COPY go.sum ./
# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY ./banking ./banking

WORKDIR /app/banking

# compile application
RUN go build -o /banking-service



FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /banking-service /app/banking/banking-service
# command to be used to execute when the image is used to start a container
# tells Docker that the container listens on specified network ports at runtime
EXPOSE 3001
CMD ["./banking/banking-service"]
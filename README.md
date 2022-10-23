# Banking App golang Microservices

## Introduction

### This repo contains two Microservices

1. **Banking**
2. **Auth-Banking**

#### Banking service is responsible for below functionality provided via API endpoints.

- Open new account for customer
- Making a Deposit or Withdrawal Transaction in the accounts.

#### Auth-Banking service is use to apply security & authorization to above banking application.

- Authorization is applied to the APIs
- Role based access control (RBAC)

#### Setup Instructions

#### This project uses docker containers please ensure docker is installed on your system.Open a terminal in this directory and run the following command `docker-compose up -d`

This creates the following docker containers

- mysql
- auth-banking-service
- banking-service

Once the containers are up and running, you can import postman collection and environment file then start using the postman and refer to collection documentation to know more about API usage.

#### API Usage

Use of each and every API is mentioned in a dedicated documentation page visit [API Documentation Link](https://documenter.getpostman.com/view/18769640/2s84DvpyXY) visit now.

#### Technical learning points along with libraries used

1. Hexagonal Architecture implemented across all microservices
2. Banking-Auth microservice based on OAuth
3. Decoupling the domain objects and Data transfer objects
4. Dependency inversion in Go and working with stubs
5. Routing capabilities of gorilla/mux, JWT Tokens for authentication & SQLX for database interaction
6. Designed own custom error and logging [library](https://github.com/gaurav-js-dev/go-banking-lib).

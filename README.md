# Banking App golang Microservices

## Introduction

### This repo contains 2 Microservices **Banking** and **Auth-Banking**

#### Banking service is responsible for below functionality provided via API endpoints.

- Open new account for customer
- Making a Deposit or Withdrawal Transaction in the accounts.
- Auth-Banking service is use to apply security & authorization to above banking application.
- Authorization is applied to the APIs
- Role based access control (RBAC)

#### Setup Instructions

#### This project uses docker containers please ensure docker is installed on your system.Open a terminal in this directory and run the following command `docker-compose up -d`

- This creates the following docker containers

  - mysql
  - auth-banking-service
  - banking-service

- Once the containers are up and running, you can start using the postman.

#### Technical learning points along with libraries used

1. Hexagonal Architecture implemented across all microservices
2. Banking-Auth microservice based on OAuth
3. Decoupling the domain objects and DTOs
4. Dependency inversion in Go and working with stubs
5. Routing capabilities of gorilla/mux, JWT Tokens for authentication & SQLX for database interaction
6. Designed own custom error and logging [library](https://github.com/gaurav-js-dev/go-banking-lib).

### Hexagonal Architecture

![hex1](https://user-images.githubusercontent.com/23628920/195387386-083bc876-4367-4621-bb4c-7178969cd193.png)
![hex2](https://user-images.githubusercontent.com/23628920/195387443-5f551cbf-f497-4794-8d41-f9d9eade53b3.png)

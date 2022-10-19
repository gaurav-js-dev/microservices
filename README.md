# Banking App golang Microservices

## Introduction

This repo contains 2 Microservices **Banking** and **Auth-Banking**

- Banking service is responsible for below functionality provided via API endpoints.

● Open new account for customer
● Making a Deposit or Withdrawal Transaction in the accounts.

- Auth-Banking service is use to apply security & authorization to above banking application.

● Authorization is applied to the APIs
● Role based access control (RBAC)

## Setup Instructions

## Libraries used

● Hexagonal Architecture implemented across all microservices
● Banking-Auth microservice based on OAuth
● Decoupling the domain objects and DTOs
● Dependency inversion in Go and working with stubs
● JWT Tokens
● Designed own custom error library

![hex1](https://user-images.githubusercontent.com/23628920/195387386-083bc876-4367-4621-bb4c-7178969cd193.png)
![hex2](https://user-images.githubusercontent.com/23628920/195387443-5f551cbf-f497-4794-8d41-f9d9eade53b3.png)

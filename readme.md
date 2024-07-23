# API Documentation

## Overview

This project provides a set of RESTful APIs for user management. Below are the endpoints available, along with examples of how to use them.

## API Endpoints

### 1. Login

Authenticate a user and obtain a session token.

**Endpoint:**



**Request:**
Endpoint : POST /user/login

```bash

curl --location '127.0.0.1:8080/user/login' \
--form 'user_name="thangth1"' \
--form 'password="1234"'

2. Get List of Sessions
Endpoint: GET /user/list

curl --location '127.0.0.1:8080/user/list' \
--header 'Authorization: token-your'

3. Logout
Terminate a session. You need to provide an authorization token obtained from the login step.

Endpoint:POST /user/logout

curl --location --request POST '127.0.0.1:8080/user/logout' \
--header 'Authorization: token-your'


Endpoint: POST /user/register
Request:

curl --location '127.0.0.1:8080/user/register' \
--header 'Authorization: token-your' \
--form 'user_name=""' \
--form 'display_name=""' \
--form 'password=""' \
--form 'avatar=""'


Project Setup
1. Initialize Go Modules
Initialize the Go module for the project.

go mod init ap_sell_products


2. Set Go Vendor Mode
Set the Go flags to use the vendor directory.

export GOFLAGS=-mod=vendor


3. Run the Project
Start the application.

go run cmd/main.go

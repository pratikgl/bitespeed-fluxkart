# Bitespeed - Identity Reconciliation

## Overview

This would provide a way to identify and keep track of a customer's identity across multiple purchases.

It would link different orders made with different contact information to the same person.

## Prerequisites

Please ensure that you have the following installed:\
[golang](https://golang.org/doc/install)\
[mysql](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/)

## Installation

```bash
git clone https://github.com/pratikgl/bitespeed-fluxkart.git
cd bitespeed-fluxkart
go mod tidy
```

## Setup

1. open the `pkg/config/app.go` file and replace the following values with your own

```go
const (
	DBUserName = "your_username"
	DBPassword = "your_password"
	DBName     = "your_db_name"
)
```

2. Create a database with the name you provided in the previous step

3. Run the following command to start the server

```
go run main.go
```

4. Open Postman and navigate to `http://localhost:8080/`

5. You can now use the API

## Usage

There is only one API endpoint for this project.\
It is a POST request to `http://localhost:8080/identify`

The request body should be a JSON object with the following structure:

```typescript
{
	"email"?: string,
	"phoneNumber"?: number
}
```

The response will be a JSON object with the following structure:

```typescript
{
  "contact":{
    "primaryContatctId": number,
    "emails": string[], // first element being email of primary contact
    "phoneNumbers": string[], // first element being phoneNumber of primary contact
    "secondaryContactIds": number[] // Array of all Contact IDs that are "secondary" to the primary contact
  }
}
```

## References

https://bitespeed.notion.site/Bitespeed-Backend-Task-Identity-Reconciliation-53392ab01fe149fab989422300423199

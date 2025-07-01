# Loyalty Cash App Backend

This is a backend service built using the [Gin](https://github.com/gin-gonic/gin) web framework in Go. It provides a RESTful API with secure and efficient request handling.

## 📁 Project Structure

├── main.go
├── go.mod / go.sum
├── config/ # Environment configuration and database connection
├── controllers/ # API controllers
├── models/ # Database models
├── services/ # Business logic
├── repositories/ # DB interaction layer
├── routes/ # Route grouping
|── utils/ # Helpers/utilities
├── dtos/ # Dtos for mapping responses

## Setup Instructions

### 1. Clone the repository

```bash
git clone <https://github.com/SasinduNanayakkara/loyalty-backend>
```

## Envs

DB_URL=""
PORT=8081
LOYALTY_APP_ID=""
LOYALTY_ACCESS_TOKEN=""
LOYALTY_API_URL="https://connect.squareupsandbox.com/v2"
SQUARE_VERSION="2025-06-18"
LOYALTY_PROGRAM_ID=""
LOYALTY_LOCATION_ID=""
JWT_SECRET=""

## Run the server
```bash
go run main.go
```

## Technology Used

Go

Gin

GORM (for ORM)

MySQL or (as database)

JWT for authentication


# KreditPlus

## Installation

To run this project need:

1. [Taskfile](https://taskfile.dev/installation/)
2. [Go](https://go.dev/doc/install)
3. [Postman](https://www.postman.com/downloads/)
4. [Docker](https://docs.docker.com/engine/install/)

run this command on terminal it will download dependencies

```bash
task setup
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`GOOSE_DRIVER`

`GOOSE_DBSTRING`

`GOOSE_MIGRATION_DIR`

`JWT_SECRET`

`PORT`

## Run Locally

Run with docker

```bash
docker-compose up -d
```

Run without docker

```bash
task run
```

## Documentation

import using postman this json

```json
{
  "info": {
    "_postman_id": "22735c77-bfd9-4b25-a50e-c41ede25efba",
    "name": "Kredit Plus",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
    "_exporter_id": "14623263"
  },
  "item": [
    {
      "name": "Register User",
      "request": {
        "method": "POST",
        "header": [],
        "body": {
          "mode": "raw",
          "raw": "{\r\n    \"email\": \"test@mail.com\",\r\n    \"password\": \"password\",\r\n    \"name\": \"test\"\r\n}",
          "options": {
            "raw": {
              "language": "json"
            }
          }
        },
        "url": {
          "raw": "http://localhost:6565/api/v1/users/register",
          "protocol": "http",
          "host": ["localhost"],
          "port": "6565",
          "path": ["api", "v1", "users", "register"]
        }
      },
      "response": []
    },
    {
      "name": "Customers",
      "request": {
        "auth": {
          "type": "bearer",
          "bearer": [
            {
              "key": "token",
              "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY1ODU3NjMsInN1YiI6IjAxOTQ0ZjZkLTcxNTctN2QzNS04MmZiLWQzZGE1OGQzZDVlZSJ9.RgiRSVzXup7eiYkDW1rMkSlu16P4GrCuwwFPq8v0AXQ",
              "type": "string"
            }
          ]
        },
        "method": "POST",
        "header": [],
        "body": {
          "mode": "raw",
          "raw": "{\r\n    \"identification_number\": \"123456\",\r\n    \"full_name\": \"name full\",\r\n    \"legal_name\": \"legal name\",\r\n    \"place_of_birth\": \"tempat lahir\",\r\n    \"date_of_birth\": \"1989-08-07\",\r\n    \"salary\": \"12000000\",\r\n    \"photo_ktp\": \"ktp url\",\r\n    \"photo_selfie\": \"selfie url\",\r\n    \"customer_limits\": [\r\n        {\r\n            \"tenor\": 1,\r\n            \"limit_amount\": 100000\r\n        },\r\n        {\r\n            \"tenor\": 2,\r\n            \"limit_amount\": 200000\r\n        },\r\n        {\r\n            \"tenor\": 3,\r\n            \"limit_amount\": 500000\r\n        },\r\n        {\r\n            \"tenor\": 6,\r\n            \"limit_amount\": 700000\r\n        }\r\n    ]\r\n}",
          "options": {
            "raw": {
              "language": "json"
            }
          }
        },
        "url": {
          "raw": "http://localhost:6565/api/v1/customers/register",
          "protocol": "http",
          "host": ["localhost"],
          "port": "6565",
          "path": ["api", "v1", "customers", "register"]
        }
      },
      "response": []
    },
    {
      "name": "Login",
      "request": {
        "method": "POST",
        "header": [],
        "body": {
          "mode": "raw",
          "raw": "{\r\n    \"email\": \"test@mail.com\",\r\n    \"password\": \"password\"\r\n}",
          "options": {
            "raw": {
              "language": "json"
            }
          }
        },
        "url": {
          "raw": "http://localhost:6565/api/v1/users/login",
          "protocol": "http",
          "host": ["localhost"],
          "port": "6565",
          "path": ["api", "v1", "users", "login"]
        }
      },
      "response": []
    },
    {
      "name": "Transaction",
      "request": {
        "method": "POST",
        "header": [],
        "body": {
          "mode": "raw",
          "raw": "{\r\n    \"customer_id\": \"01944f6d-b4bd-71b5-9ca3-175ffcae3260\",\r\n    \"otr\": 700000,\r\n    \"admin_fee\": 20000,\r\n    \"installment\": 100000,\r\n    \"interest\": 1000,\r\n    \"asset_name\": \"CAR\",\r\n    \"tenor\": 6\r\n}",
          "options": {
            "raw": {
              "language": "json"
            }
          }
        },
        "url": {
          "raw": "http://localhost:6565/api/v1/transactions/create",
          "protocol": "http",
          "host": ["localhost"],
          "port": "6565",
          "path": ["api", "v1", "transactions", "create"]
        }
      },
      "response": []
    }
  ]
}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Authors

- [@sutantodadang](https://www.github.com/sutantodadang)

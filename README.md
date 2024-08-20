# Task Management App

## How to run

### Docker

Prerequisites:

1. Docker
2. Internet

Steps to run:

1. Run command `make run-docker`

### Local

Prerequisites:

1. Go
2. Postgresql
3. Install swag

Steps to run:

1. Clone the project
2. Go to the directory
3. Run command `go mod tidy`
4. copy .env.example to .env and fill the POSTGRE_* environment value according to your postgresql
5. Run command `make run-local`

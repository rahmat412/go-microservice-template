# go-microservice-template

A template for building microservices in Go. This repository provides a starting point for creating new services with a standardized structure and tools.

## Tech Stack

- **Go**: Programming language.
- **Chi**: Lightweight router for building HTTP services.
- **Go Jett**: Database schema management.
- **Goose**: Database migrations.
- **Postgres**: Relational database.

## Features

- Pre-configured project structure.
- Database migration and schema management tools.
- Example API server using Chi.
- Docker support for containerization.
- Makefile for common tasks.

## Getting Started

### Prerequisites

- Go 1.20+
- Docker & Docker Compose
- PostgreSQL
- `jet` CLI for schema management
- `goose` CLI for migrations

### Setup

1. Clone this repository:
   ```bash
   git clone https://github.com/rahmat412/go-microservice-template.git
   cd go-microservice-template
   ```
2. Install dependencies:
   ```bash
    go mod tidy
   ```
3. Make sure to install `jet` and `goose`:

   ```bash
   go install github.com/bitfield/script@latest
   go install github.com/pressly/goose/v3/cmd/goose@latest
   ```

   and ensure they are in your `PATH`.

4. Create a `.env` file in the root directory and configure your database connection:
   ```env
   DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
   ```

### Handling Database Migrations

1. Create a new migration:
   ```bash
   make new-migrations
   ```
   This will create a new migration file in the `database/migrations` directory based on the name you provide after `make new-migrations`.
2. Run the service:
   ```bash
   make run-api
   ```
   After we create the migrations, you need to run the service to apply the migrations to the database. The service will automatically run the migrations on startup.
3. Sync Go-Jett schema:
   ```bash
   make sync-jet
   ```
   You need to run this command to sync the Go-Jett schema with the database whenever you make changes to the database schema. This will generate the Go code for the database schema in the `database/{database_name}` directory.

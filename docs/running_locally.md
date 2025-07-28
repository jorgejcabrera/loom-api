# 📦 Running Locally

## Prerequisites

Before starting, ensure you have the following installed:

- Docker (latest stable version)
- Docker Compose (latest stable version)
- Go 1.24 or later

## Setup Steps

### 1. Start Required Services

Launch the required services using Docker Compose:

```shell
docker-compose up -d
```

This command will start:

- Temporal server
- Temporal UI (accessible at http://localhost:8082)
- PostgreSQL database
- Elasticsearch
- Other required services

### 2. Build the Application

Follow these steps to build the application:

```shell
go mod tidy
go mod vendor
go build -o loom-api cmd/main.go
```

### 3. Run the Application

Execute the built binary:

```shell
./loom-api
```

The application will start and listen on the default port (8081).

## Verification

To verify that everything is running correctly:

1. Check that all Docker containers are running:

```shell
docker  ps
```

2. Ensure the application is responsive:

```shell
curl localhost:8081/health
```

3. Access Temporal UI at: http://localhost:8082


# Description

A go boilerplate for building scalable and maintainable web applications in Go.

## Setup
Start Infra

```bash
docker compose up -d postgres redis rustfs
```

## Running in development
Running server
```bash
go run cmd/main.go server
```
Running worker
```bash
go run cmd/main.go worker
```
Or Running server
```bash
air
```

## Production
Start app production
```bash
docker compose build myapp

docker compose up -d --force-recreate myapp
docker compose up -d --force-recreate worker
```

Run migrations command
```bash
go run cmd/main.go migrate up
```

Run migrations using docker compose
```bash
docker compose run --rm myapp migrate up
```

or from source code
Run migrations command
```bash
go run cmd/main.go migrate
```

## fieldalignment
```bash
# Check fieldalignment
fieldalignment ./...
# Fix fieldalignment
fieldalignment -fix ./...
```


## Technology used

- Gin – High-performance HTTP web framework
- GORM – Powerful ORM for database operations
- Goose - Migration management
- Uber Fx – Dependency injection and application lifecycle management
- Cobra – Command-line application framework
- Zap – Logging framework
- Asynq – Asynchronous task queue
- Postgres – Relational database management system
- Docker – Containerization platform

Monitoring and logging
- Prometheus – Metrics collection and monitoring system
- Loki – Log aggregation and storage system
- Grafana – Visualization and dashboarding tool

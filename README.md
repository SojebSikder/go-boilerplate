# Description

A go boilerplate for building scalable and maintainable web applications in Go.

## Setup
Run migrations command
```
go run cmd/main.go migrate up
```

## Running

Running server
```bash
go run cmd/main.go server
```

## Running using Docker (for development)

```
docker compose -f docker-compose.yml up -d redis postgres prometheus grafana loki promtail node_exporter rustfs --build
```
## Running using Docker (for production)

```
docker compose -f docker-compose.yml --profile prod up --remove-orphans --force-recreate --build -d
```

## Technology used

- Gin – High-performance HTTP web framework
- GORM – Powerful ORM for database operations
- Goose - Migration management
- Uber Fx – Dependency injection and application lifecycle management
- Cobra – Command-line application framework
- Zap – Logging framework
- Postgres – Relational database management system
- Docker – Containerization platform

Monitoring and logging
- Prometheus – Metrics collection and monitoring system
- Loki – Log aggregation and storage system
- Grafana – Visualization and dashboarding tool

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

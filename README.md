# Clean user service

[![Generic badge](https://img.shields.io/badge/LICENSE-MIT-orange.svg)](LICENSE)
[![Generic badge](https://img.shields.io/badge/DOCKER-HUB-blue.svg)](https://hub.docker.com/repository/docker/dangdancheg/discord_alerts)
[![Generic badge](https://img.shields.io/badge/SWAGGER-API-green.svg)](https://app.swaggerhub.com/apis/Dancheg97/clean_svc/1)

Service for operations with users.

## API description

- [Proto](users.proto)
- [Generated code](pb)
- [Swagger json](users.swagger.json)

## Adresses:

- http local adress:

```
localhost:8080
```

- grpc local adress:

```
localhost:9080
```

## Monitoring

Grafana link: [grafana_logs](nan)

## Local start

Required env file to launch service locally:

```ruby
POSTGRES_USER="user"
POSTGRES_PASSWORD="password"
POSTGRES_HOST="host.docker.internal"
POSTGRES_PORT="5432"
POSTGRES_DB="db"
APP_GRPC_PORT="9080"
APP_HTTP_PORT="8080"
JSON_LOGS="false"
```

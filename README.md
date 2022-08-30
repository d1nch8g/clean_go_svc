# <p  align="center" style="font-family:courier;font-size:180%" size=212px> Clean go service template </p>

<p align="center">
<img align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="238px" height="238px" src="docs/logo.png" /> 
</p>

[![Generic badge](https://img.shields.io/badge/LICENSE-MIT-orange.svg)](LICENSE)
[![Generic badge](https://img.shields.io/badge/DOCKER-HUB-blue.svg)](https://hub.docker.com/repository/docker/dangdancheg/clean_svc)

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

## Compose to launch service

```yaml
version: "3.8"

services:
  app:
    image: dangdancheg/clean_svc
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_HOST: host.docker.internal
      POSTGRES_PORT: 5432
      POSTGRES_DB: db
      JSON_LOGS: false
    ports:
      - "9080:9080"
      - "8080:8080"
    depends_on:
      - "db"

  db:
    image: postgres
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: db
    ports:
      - "5432:5432"
    extra_hosts:
      - "host.docker.internal:host-gateway"
```

## DB schema

![](schema.png)

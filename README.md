# Clean user service

Service for operations with users.

## API description

- [Proto](users.proto)
- [Generated code](pb)
- [Swagger](users.swagger.json)

## Adresses:

- http local adress example after `docker-comose`:

```
localhost:8080
```

- пкзс local adress example after `docker-comose`:

```
localhost:9080
```

## Monitoring

Grafana link: [grafana_modelrepo](nan)

## Local start

Required env file to launch service locally:

```ruby
POSTGRES_USER="user"
POSTGRES_PASSWORD="password"
POSTGRES_HOST="localhost"
POSTGRES_PORT="5432"
POSTGRES_DB="db"
APP_GRPC_PORT="9080"
APP_HTTP_PORT="8080"
```

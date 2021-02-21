# ISSPAY

## Outline

- [Install](#install)
- [Run Local Application](#run-local-application)

### Install

1. install migration tool ([goose](https://github.com/pressly/goose))

```shell
go get -u github.com/pressly/goose/cmd/goose
```

2. run postgreSQL server and execute db schema migration

```shell
make dev.db
migrate.dev.up
```

3. initialize configuration and secrets file

```shell
make app.yaml
touch ./configs/secrets.yaml
```

4. run server

```shell
go module vendor
make all
```

### Run Local Application

### References

- [linebot message type](https://developers.line.biz/en/docs/messaging-api/message-types/#template-messages)
- [action object](https://developers.line.biz/en/reference/messaging-api/#action-objects)
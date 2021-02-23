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

requirements:

1. [docker](https://docs.docker.com/engine/install/ubuntu/)
2. [docker-compose](https://docs.docker.com/compose/install/)
3. [linebot token](https://developers.line.biz/zh-hant/services/bot-designer/)

steps:

1. create docker.env file and copy line bot token and secret

```shell
  touch ./build/docker/.docker.env
```

2. run docker-compose file

```shell
docker-compose -f ./build/docker/docker-compose.yaml up
```

### References

- [linebot message type](https://developers.line.biz/en/docs/messaging-api/message-types/#template-messages)
- [action object](https://developers.line.biz/en/reference/messaging-api/#action-objects)
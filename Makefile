

.PHONY: migrate.file

migrate.file:
	goose -dir $(CURDIR)/deployments/migrations create $(f) sql

.PHONY: migrate.test.up

migrate.dev.up:
	goose -dir=$(CURDIR)/deployments/migrations postgres "user=dev dbname=isspay_dev password=dev host=127.0.0.1 port=5432 sslmode=disable" up

.PHONY: migrate.test.down

migrate.dev.down:
	goose -dir=$(CURDIR)/deployments/migrations postgres "user=dev dbname=isspay_dev password=dev host=127.0.0.1 port=5432 sslmode=disable" down

pg:
	docker run -d --name pg -e POSTGRES_USER=dev \
						-e POSTGRES_DB=isspay_dev \
						-e POSTGRES_PASSWORD=dev \
						-p 5432:5432 \
						postgres:12.3-alpine

build.image:
	docker build -f $(CURDIR)/build/docker/isspay.dockerfile -t isspay .


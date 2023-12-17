.PHONY: help database-up flyway-migration flyway-info

database-up:
	docker run --name postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_DB=project000 \
		-p 5432:5432 \
		-d postgres:14.8-alpine

flyway-migration:
	docker run --rm \
		--network host \
		-v $(shell pwd)/script/database/sql:/flyway/sql \
		-v $(shell pwd)/script/database/config/flyway-local.conf:/flyway/conf/flyway.conf \
		flyway/flyway:latest migrate

flyway-info:
	docker run --rm \
		--network host \
		-v $(shell pwd)/script/database/sql:/flyway/sql \
		-v $(shell pwd)/script/database/config/flyway-local.conf:/flyway/conf/flyway.conf \
		flyway/flyway:latest info
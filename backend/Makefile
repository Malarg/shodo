localConfigName ?= local

run:
	CONFIG_NAME=$(localConfigName) docker-compose up --build -d
	docker-compose logs -f backend
swagger:
	swag init -d ./cmd/app,./internal/app,./internal/transport,./models

testConfigName ?= test

test.integration:
	CONFIG_NAME=$(testConfigName) docker-compose up --build -d
	go test -v ./tests/; TEST_EXIT_CODE=$$?; \
	docker-compose logs backend; \
	docker-compose down; \
	exit $$TEST_EXIT_CODE
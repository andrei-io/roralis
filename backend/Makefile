.PHONY: build
build: clean
	go build -o build/server main.go

.PHONY: run
run: build
	./build/server

.PHONY: watch
watch:
	air -c .air.toml

.PHONY: clean
clean:
	rm -rf build/

.PHONY: migrate-up
migrate-up:
	goose -dir infrastructure/migrations postgres "user=admin password=admin dbname=asdf sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	goose -dir infrastructure/migrations postgres "user=admin password=admin dbname=asdf sslmode=disable" down

.PHONY: migrate-reset
migrate-reset:
	goose -dir infrastructure/migrations postgres "user=admin password=admin dbname=asdf sslmode=disable" reset

.PHONY: migrate-status
migrate-status:
	goose -dir infrastructure/migrations postgres "user=admin password=admin dbname=asdf sslmode=disable" status

.PHONY: swagger
swagger:
	swagger generate spec -o doc/swagger.yml

.PHONY: start-containers
start-containers:
	docker-compose -f docker-compose.dev.yml up 

.PHONY: teardown-containers
teardown-containers:
	docker-compose -f docker-compose.dev.yml down 


GOLANGCI_LINT_CACHE?=./golangci-lint/.cache/golangci-lint/v1.57.2

.PHONY: golangci-lint-run
golangci-lint-run: _golangci-lint-rm-unformatted-report

.PHONY: _golangci-lint-reports-mkdir
_golangci-lint-reports-mkdir:
	mkdir -p ./golangci-lint

.PHONY: _golangci-lint-run
_golangci-lint-run: _golangci-lint-reports-mkdir
	-docker run --rm \
    -v $(shell pwd):/app \
    -v $(GOLANGCI_LINT_CACHE):/root/.cache \
    -w /app \
    golangci/golangci-lint:v1.57.2 \
        golangci-lint run \
            -c .golangci.yml \
	> ./golangci-lint/report-unformatted.json

.PHONY: _golangci-lint-format-report
_golangci-lint-format-report: _golangci-lint-run
	cat ./golangci-lint/report-unformatted.json | jq > ./golangci-lint/report.json

.PHONY: _golangci-lint-rm-unformatted-report
_golangci-lint-rm-unformatted-report: _golangci-lint-format-report
	rm ./golangci-lint/report-unformatted.json

.PHONY: golangci-lint-clean
golangci-lint-clean:
	sudo rm -rf ./golangci-lint 

.PHONY: create-migration
create-migration:
	-docker run --rm \
    -v $(realpath ./internal/db/migrations):/migrations \
    migrate/migrate:v4.18.1 \
        create \
        -dir /migrations \
        -ext .sql \
        -seq -digits 5 \
        $(n)

.PHONY: migrations-up
migrations-up:
	-docker run --rm \
    --network gophermart_default \
    -v $(realpath ./internal/db/migrations):/migrations \
    migrate/migrate:v4.18.1 \
        -path=/migrations \
        -database postgres://gophermart:gophermart@db:5432/gophermart?sslmode=disable \
        up
    
.PHONY: migrations-down
migrations-down:
	-docker run --rm \
    --network gophermart_default \
    -v $(realpath ./internal/db/migrations):/migrations \
    migrate/migrate:v4.18.1 \
        -path=/migrations \
        -database postgres://gophermart:gophermart@db:5432/gophermart?sslmode=disable \
        down --all
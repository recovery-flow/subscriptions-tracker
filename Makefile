DB_URL=postgresql://postgres:postgres@localhost:7003/postgres?sslmode=disable
OPENAPI_GENERATOR := java -jar ~/openapi-generator-cli.jar
CONFIG_FILE := ./config_local.yaml
API_SRC := ./docs/api.yaml
API_BUNDLED := ./docs/api-bundled.yaml
OUTPUT_DIR := ./docs/web
RESOURCES_DIR := ./resources

generate-models:
	find $(RESOURCES_DIR) -type f ! \( -name "resources_types.go" -o -name "links.go" \) -delete
	swagger-cli bundle $(API_SRC) --outfile $(API_BUNDLED) --type yaml

	$(OPENAPI_GENERATOR) generate \
		-i $(API_BUNDLED) -g go \
		-o $(OUTPUT_DIR) \
		--additional-properties=packageName=resources

	mkdir -p $(RESOURCES_DIR)
	find $(OUTPUT_DIR) -name '*.go' -exec mv {} $(RESOURCES_DIR)/ \;
	find $(RESOURCES_DIR) -type f -name "*_test.go" -delete

generate-sqlc:
	sqlc generate

migrate-up:
	KV_VIPER_FILE=$(CONFIG_FILE) go build -o main ./cmd/subscriptions-tracker/main.go
	migrate -path internal/service/infra/data/sqldb/migrations -database $(DB_URL) -verbose up

migrate-down:
	KV_VIPER_FILE=$(CONFIG_FILE) go build -o main ./cmd/subscriptions-tracker/main.go
	migrate -path internal/service/infra/data/sqldb/migrations -database $(DB_URL) -verbose down

run-server:
	KV_VIPER_FILE=$(CONFIG_FILE) go build -o main ./cmd/subscriptions-tracker/main.go
	KV_VIPER_FILE=$(CONFIG_FILE) ./main run service
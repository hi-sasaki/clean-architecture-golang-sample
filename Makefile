CLI_DIR=cmd/cli
COMMAND=
API_DIR=cmd/api
DIST_DIR=bin

build-api: wire-api
	go build -o $(DIST_DIR)/api $(API_DIR)/main.go $(API_DIR)/wire_gen.go

run-api: wire-api
	go run $(API_DIR)/main.go $(API_DIR)/wire_gen.go

wire-api:
	cd $(API_DIR) && \
	wire gen


build-cli: wire-cli
	go build -o $(DIST_DIR)/$(COMMAND) $(CLI_DIR)/$(COMMAND)/main.go $(CLI_DIR)/$(COMMAND)/wire_gen.g

run-cli: wire-cli
	go run $(CLI_DIR)/$(COMMAND)/main.go $(CLI_DIR)/$(COMMAND)/wire_gen.go

wire-cli:
	cd $(CLI_DIR)/$(COMMAND) && \
	wire gen
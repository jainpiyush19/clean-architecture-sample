BUILD_ENVPARMS:=GOGC=off CGO_ENABLED=0
LOCAL_BIN:=$(CURDIR)/bin
BIN?=$(LOCAL_BIN)/wallet
GOPACKAGES=$(shell go list ./... | grep -v 'cryptowallet/gen' | grep -v 'cryptowallet/design' | grep -v 'cryptowallet/pkg')

help: ## Display this help
	@ echo "Please use \`make <target>' where <target> is one of:"
	@ echo
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-10s\033[0m - %s\n", $$1, $$2}'
	@ echo

.PHONY: deps
deps: ## Install dependencies
	$(info #Install dependencies...)
	go mod tidy

.PHONY: build
build: deps ## build executable
	$(info #Building...)
	$(BUILD_ENVPARMS) go build -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/wallet

.PHONY: run
run: ## run the service with docker compose
	$(info #Running...)
	docker-compose up --build

.PHONY: run_background
run_background:
	$(info #Running...)
	docker-compose up --build -d

.PHONY: stop
stop:
	$(info #Running...)
	docker-compose down

.PHONY: test
test: ## run all test for service
	@$(eval COVERFILE := "wallet.coverage")
	@go get -u github.com/rakyll/gotest
	@go get -u github.com/vektra/mockery/.../
	@mockery -all -inpkg
	gotest -coverprofile=$(COVERFILE).tmp -v -coverpkg=$(shell echo $(GOPACKAGES) | tr " " ",") ./...
	@echo Code Coverage
	@cat $(COVERFILE).tmp | grep -v "mock_" > $(COVERFILE)
	@rm $(COVERFILE).tmp
	@go tool cover -func=$(COVERFILE)

.PHONY: load_test
load_test: run_background ## run load test for the service
	npx wait-on -t 30000 http://localhost:3000/health
	k6 run -e BASE_URL=http://localhost:3000 scripts/loadtest/k6.js
	make stop

VERSION := 0

VERSION_PATH := "v$(VERSION)"
TOOL_PATH := "$(VERSION_PATH)/tools/gen"

.PHONY: default
default:
	@echo "Pick a card"

.PHONY: gen-maps
gen-maps:
	@gomp-gen "$(VERSION_PATH)/extras/gomp-conf.yml"
	@gofmt -s -w .

.PHONY: gen-examples
gen-examples:
	@go run "$(TOOL_PATH)/example/main.go"
	@gofmt -s -w .

.PHONY: gen-types
gen-types:
	@go run "$(TOOL_PATH)/type/main.go"
	@gofmt -s -w .

.PHONE: gen-doc
gen-doc:
	@go run "$(TOOL_PATH)/doc/main.go"
	@gofmt -s -w .

.PHONY: gen-all
gen-all: gen-doc gen-maps gen-examples gen-types

.PHONY: install-dev
install-dev:
	@which gomp-gen > /dev/null || go get -u github.com/Foxcapades/gomp/v1/cmd/gomp-gen
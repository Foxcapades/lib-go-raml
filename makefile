

default:
	@echo "Pick a card"

.PHONY: gen-maps
gen-maps:
	@go run v0/tools/gen/store/main.go
	@gofmt -s -w .

.PHONY: gen-examples
gen-examples:
	@go run v0/tools/gen/example/main.go
	@gofmt -s -w .

.PHONY: gen-types
gen-types:
	@go run v0/tools/gen/type/main.go
	@gofmt -s -w .

.PHONY: gen-all
gen-all: gen-maps gen-examples gen-types
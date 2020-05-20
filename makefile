

default:
	@echo "Pick a card"

.PHONY: gen-maps
gen-maps:
	@go run v0/tools/gen/store/main.go

.PHONY: gen-examples
gen-examples:
	@go run v0/tools/gen/example/main.go

.PHONY: gen-types
gen-types:
	@go run v0/tools/gen/type/main.go

.PHONY: gen-all
gen-all: gen-maps gen-examples gen-types


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

.PHONE: gen-doc
gen-doc:
	@go run v0/tools/gen/doc/main.go
	@gofmt -s -w .

.PHONY: gen-all
gen-all: gen-doc gen-maps gen-examples gen-types

install-dev:
	@which gomp-gen > /dev/null || go get -u github.com/Foxcapades/gomp/v1/cmd/gomp-gen
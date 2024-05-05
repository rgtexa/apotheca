
#
# BUILD
#

## build the apotheca binary
.PHONY: build/apotheca
build/apotheca:
	@echo 'Building apotheca binary...'
	go build -o=./bin/apotheca ./cmd/main.go
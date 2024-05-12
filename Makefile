
#
# BUILD
#

## build the apotheca binary
.PHONY: build/apotheca
build/apotheca:
	@echo 'Building apotheca binary...'
	go build -o=./bin/apotheca ./cmd/main.go

## run the dev server locally
.PHONY: run/apotheca
run/apotheca: build/apotheca
	./bin/apotheca run

## run tailwind in watch mode
.PHONY: build/tailwind
build/tailwind:
	@echo 'Running Tailwind in watch mode...'
	npx tailwindcss -i ./ui/static/css/dev.css -o ./ui/static/css/apotheca.css --watch
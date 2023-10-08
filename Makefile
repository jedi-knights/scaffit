.PHONE: lint lint-js lint-go

all: lint

lint-js:
	npx eslint .releaserc.js commitlint.config.js

lint-go:
	golangci-lint run ./...

lint: lint-js lint-go
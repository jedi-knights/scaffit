.PHONE: lint lint-js lint-go

all: lint

fix-js:
	npx eslint .releaserc.js commitlint.config.js --fix

lint-js: fix-js
	npx eslint .releaserc.js commitlint.config.js

lint-go:
	golangci-lint run ./...

lint: lint-js lint-go

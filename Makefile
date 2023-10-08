.PHONE: lint lint-js lint-go

all: lint

clean:
	rm -f junit.xml

fix-js:
	npx eslint .releaserc.js commitlint.config.js --fix

lint-js: fix-js
	npx eslint .releaserc.js commitlint.config.js

lint-go:
	golangci-lint run ./...

lint: lint-js lint-go

test:
	ginkgo --junit-report junit.xml ./...

build:
	go build -o scaffit main.go

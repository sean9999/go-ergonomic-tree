REPO := $$(go mod why | tail -n 1)
SEMVER := $$(git tag --sort=-version:refname | head -n 1)
GOPROXY=proxy.golang.org

.PHONY: test

build:
	GOEXPERIMENT=aliastypeparams go build -o ergo .

tidy:
	go mod tidy

clean:
	go clean

publish:
	GOPROXY=${GOPROXY} list -m ${REPO}@${SEMVER}

test:
	go test ./...

benchmark:
	go test -bench=. -count 5 -run=^#
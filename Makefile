BIN_FOLDER := $$(go env GOPATH)/bin
REPO := $$(go mod why | tail -n 1)
SEMVER := $$(git tag --sort=-version:refname | head -n 1)

.PHONY: test

tidy:
	go mod tidy

clean:
	go clean
	rm ${BUILD_FOLDER}/${BINARY_NAME}

publish:
	GOPROXY=proxy.golang.org go list -m ${REPO}@${SEMVER}

test:
	go test .

benchmark:
	go test -bench=. -count 5 -run=^#

VERSION := $(shell git describe --tags --always)
GIT_COMMIT := $(shell git rev-parse HEAD)
BINARY_NAME=k8sgen
SOURCE=./cmd/k8sgen
DESTINATION=./bin/local/${BINARY_NAME}
LDFLAGS := "-s -w -X github.com/wingkwong/k8sgen/cmd.Version=$(VERSION) -X github.com/wingkwong/k8sgen/cmd.GitCommit=$(GIT_COMMIT)"

.PHONY: all

.PHONY: build
build:
	go build -ldflags ${LDFLAGS} -o ${DESTINATION} ${SOURCE}

.PHONY: test
test:
	CGO_ENABLED=0 go test $(shell go list ./... | grep -v /vendor/|xargs echo) -cover

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: dist
dist:
	mkdir -p bin
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/k8sgen ${SOURCE}
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/k8sgen-darwin ${SOURCE}
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/k8sgen-armhf ${SOURCE}
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/k8sgen-arm64 ${SOURCE}
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/k8sgen.exe ${SOURCE}
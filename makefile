BINARY_NAME = transfeed
BUILD_DIR = build
VERSION = 0.0.1-alpha
OPTIONS = CGO_ENABLED=0
COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S)
ENV = -X main.Commit=$(COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.Version=$(VERSION)
DARWIN_AMD = $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_darwin_amd64
DARWIN_ARM = $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_darwin_arm64
LINUX_AMD = $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_linux_amd64
LINUX_ARM = $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_linux_arm64
WIN_AMD = $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_windows_amd64
WIN_ARM = $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_windows_arm64

default: build

build:
	# 编译为 macOS 平台 amd64
	GOOS=darwin GOARCH=amd64 $(OPTIONS) go build -trimpath -ldflags="-s -w $(ENV)" -o $(DARWIN_AMD)/transfeed main.go

	# 编译为 macOS 平台 arm64
	GOOS=darwin GOARCH=arm64 $(OPTIONS) go build -trimpath -ldflags="-s -w -w $(ENV)" -o $(DARWIN_ARM)/transfeed main.go

	# 编译为 linux 平台 amd64
	GOOS=linux GOARCH=amd64 $(OPTIONS) go build -trimpath -ldflags="-s -w $(ENV)" -o $(LINUX_AMD)/transfeed main.go

	# 编译为 linux 平台 arm64
	GOOS=linux GOARCH=arm64 $(OPTIONS)  go build -trimpath -ldflags="-s -w $(ENV)" -o $(LINUX_ARM)/transfeed  main.go

	# 编译为 Windows 平台 amd64
	GOOS=windows GOARCH=amd64 $(OPTIONS) go build -trimpath -ldflags="-s -w $(ENV)" -o $(WIN_AMD)/transfeed.exe main.go

	# 编译为 Windows 平台 arm64
	GOOS=windows GOARCH=amd64 $(OPTIONS) go build -trimpath -ldflags="-s -w $(ENV)" -o $(WIN_ARM)/transfeed.exe main.go

clean:
	rm -rf $(BUILD_DIR)/$(BINARY_NAME)_*/transfeed

.PHONY: build clean

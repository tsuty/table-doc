GOFLAGS := -tags netgo -installsuffix netgo -ldflags '-w -s --extldflags "-static"'
GOVERSION=$(shell go version)
GOOS=$(word 1,$(subst /, ,$(lastword $(GOVERSION))))
GOARCH=$(word 2,$(subst /, ,$(lastword $(GOVERSION))))
BUILD_DIR=build/$(GOOS)-$(GOARCH)

.PHONY: all build clean deps package package-targz

all: build

build:
	mkdir -p $(BUILD_DIR)
	cd cmd/table-doc && GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(GOFLAGS) -o ../../$(BUILD_DIR)/table-doc

clean:
	rm -rf build package

test:
	docker-compose run table-doc

package:
	$(MAKE) package-targz GOOS=linux GOARCH=amd64
	$(MAKE) package-targz GOOS=darwin GOARCH=amd64

package-targz: build
	mkdir -p package
	cd $(BUILD_DIR) && tar czvf ../../package/table-doc_$(GOOS)_$(GOARCH).tar.gz table-doc

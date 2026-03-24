BINARY=git-who
BUILD_DIR=bin

.PHONY: build install clean

build:
	go build -o $(BUILD_DIR)/$(BINARY) ./cmd/git-who

install: build
	cp $(BUILD_DIR)/$(BINARY) ~/.local/bin/$(BINARY)

clean:
	rm -rf $(BUILD_DIR)

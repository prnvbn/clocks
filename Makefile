.PHONY : dist windows darwin linux-amd64 linux-arm64

DIST_DIR="dist"
ZIP="zip -m"

LDFLAGS="-s -w"
BUILD_FLAGS=-ldflags=$(LDFLAGS)


dist:
	mkdir -p $(DIST_DIR)
	$(MAKE) windows darwin linux-amd64 linux-arm64

.PHONY= build-windows
windows:
	GOOS=windows \
	go build $(BUILD_FLAGS) -o clocks-windows.exe && \
	zip -m $(DIST_DIR)/clocks-windows.zip clocks-windows.exe

darwin:
	GOOS=darwin \
	go build $(BUILD_FLAGS) -o clocks-darwin && \
	zip -m $(DIST_DIR)/clocks-darwin.zip clocks-darwin

linux-amd64:
	GOOS=linux && GOARCH=amd64 \
	go build $(BUILD_FLAGS) -o clocks-linux-amd64 && \
	zip -m $(DIST_DIR)/clocks-linux-amd64.zip clocks-linux-amd64

linux-arm64:
	GOOS=linux && GOARCH=arm64 \
	go build $(BUILD_FLAGS) -o clocks-linux-arm64 && \
	zip -m $(DIST_DIR)/clocks-linux-arm64.zip clocks-linux-arm64

clean:
	rm -r $(DIST_DIR)/

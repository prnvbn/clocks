.PHONY : dist windows darwin linux-arm64 linux-amd64 clean

DIST_DIR="dist"
ZIP="zip -m"

LDFLAGS="-s -w"
BUILD_FLAGS=-ldflags=$(LDFLAGS)


dist:
	mkdir -p $(DIST_DIR)
	$(MAKE) windows darwin linux-arm64 linux-amd64

.PHONY= build-windows
windows:
	GOOS=windows \
	go build $(BUILD_FLAGS) -o $(DIST_DIR)/clocks-windows.exe

darwin:
	GOOS=darwin \
	go build $(BUILD_FLAGS) -o $(DIST_DIR)/clocks-darwin

linux-arm64:
	GOOS=linux GOARCH=arm64 \
	go build $(BUILD_FLAGS) -o  $(DIST_DIR)/clocks-linux-arm64

linux-amd64:
	GOOS=linux GOARCH=amd64 \
	go build $(BUILD_FLAGS) -o  $(DIST_DIR)/clocks-linux-amd64

clean:
	rm -r $(DIST_DIR)/

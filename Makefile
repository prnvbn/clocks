.PHONY : dist windows darwin linux

DIST_DIR="dist"
ZIP="zip -m"

LDFLAGS="-s -w"
BUILD_FLAGS=-ldflags=$(LDFLAGS)


dist:
	mkdir -p $(DIST_DIR)
	$(MAKE) windows darwin linux

.PHONY= build-windows
windows:
	GOOS=windows \
	go build $(BUILD_FLAGS) -o $(DIST_DIR)/clocks-windows.exe

darwin:
	GOOS=darwin \
	go build $(BUILD_FLAGS) -o $(DIST_DIR)/clocks-darwin
	chmod +x $(DIST_DIR)/clocks-darwin

linux:
	GOOS=linux \
	go build $(BUILD_FLAGS) -o  $(DIST_DIR)/clocks-linux
	chmod +x $(DIST_DIR)/clocks-linux

clean:
	rm -r $(DIST_DIR)/

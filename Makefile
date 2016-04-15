include ../includes.mk

BINARY_DEST_DIR=.
INSTALL_DIR=$$GOPATH/bin

build: 
	@echo -e '\033[0;32m' "Building cde client" '\033[0m'
	@CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o $(BINARY_DEST_DIR)/cde cde.go || exit 1
	@$(call check-static-binary,$(BINARY_DEST_DIR)/cde)
	@echo -e '\033[0;32m' "Building cde complete" '\033[0m'

install: build
	@cp $(BINARY_DEST_DIR)/cde $(INSTALL_DIR)
	@echo -e '\033[0;32m' "Install cde complete" '\033[0m'  
test: test-unit

test-unit:
	$(GOTEST) . .

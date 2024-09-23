all: build

SKIP_CONTRACTS=n
SKIP_EXAMPLES=n
SKIP_GENESIS=n
.PHONY: build
build:
	clear

.PHONY: run_proxy
run_proxy:
	cd fuel-proxy && go run main.go
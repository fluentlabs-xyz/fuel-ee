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

.PHONY: update_contracts_and_genesis_locally
update_contracts_and_genesis_locally:
	cd crates/contracts && make
	cp crates/contracts/assets/* ../fluentbase/crates/contracts/assets/
	cd ../fluentbase/crates/genesis && make
	notify-send "done"
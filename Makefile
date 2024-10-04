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

.PHONY: build_contracts_and_update_local_genesis_and_run_fluent_node
build_contracts_and_update_local_genesis_and_run_fluent_node:
	clear
	cd crates/contracts && make fvm
	cp crates/contracts/assets/*.wasm ../fluentbase/crates/contracts/assets/
	cp crates/contracts/assets/*.rwasm ../fluentbase/crates/contracts/assets/
	cp crates/contracts/assets/*.wat ../fluentbase/crates/contracts/assets/
	cd ../fluentbase/crates/genesis && make
	cd ../fluent; rm -rf datadir
	cd ../fluent; make fluent_run
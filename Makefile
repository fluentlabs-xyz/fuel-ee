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

.PHONY: build_contracts_and_update_local_genesis
build_contracts_and_update_local_genesis:
	cd crates/contracts && make fvm
	cp crates/contracts/assets/* ../fluentbase/crates/contracts/assets/
	cd ../fluentbase/crates/genesis && make
	notify-send "done"
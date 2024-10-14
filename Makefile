all: build_contracts_and_update_genesis_locally

.PHONY: prepare
prepare:
	cd examples; $(MAKE) prepare

.PHONY: run_fuel_proxy
run_fuel_proxy:
	cd fuel-proxy && go run main.go

.PHONY: build_contracts
build_contracts:
	cd crates/contracts && $(MAKE)

.PHONY: copy_contracts_locally
copy_contracts_locally:
	cp crates/contracts/assets/*.wasm ../fluentbase/crates/contracts/assets/
	cp crates/contracts/assets/*.rwasm ../fluentbase/crates/contracts/assets/
	cp crates/contracts/assets/*.wat ../fluentbase/crates/contracts/assets/

.PHONY: update_local_genesis
update_local_genesis:
	cd ../fluentbase/crates/genesis && $(MAKE)

.PHONY: build_contracts_and_update_genesis_locally
build_contracts_and_update_genesis_locally:
	$(MAKE) build_contracts
	$(MAKE) copy_contracts_locally
	$(MAKE) update_local_genesis

.PHONY: run_fluent_node
run_fluent_node:
	cd ../fluent; $(MAKE) fluent_build
	cd ../fluent; $(MAKE) fluent_clear_datadir
	cd ../fluent; RUST_LOG=debug make fluent_run

.PHONY: build_contracts_and_update_genesis_locally_and_run_fluent_node
build_contracts_and_update_genesis_locally_and_run_fluent_node:
	$(MAKE) build_contracts_and_update_genesis_locally
	$(MAKE) run_fluent_node

.PHONY: send_example_fuel_tx
send_example_fuel_tx:
	cd examples; $(MAKE) send_example_fuel_tx

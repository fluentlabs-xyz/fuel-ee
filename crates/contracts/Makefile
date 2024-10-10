PRECOMPILES=fvm
CUR_DIR=$(shell pwd)
CUR_DIR_TOP2=$(shell realpath $(CUR_DIR)/../..)
OUT_DIR="${CUR_DIR}/assets"
RUSTFLAGS='-C link-arg=-zstack-size=131072 -C target-feature=+bulk-memory'

all: abi $(PRECOMPILES)

.PHONY: $(PRECOMPILES)
$(PRECOMPILES):
	clear
	mkdir -p $(OUT_DIR)
	RUSTFLAGS=$(RUSTFLAGS) cargo b --release --target=wasm32-unknown-unknown --no-default-features --features=$@
	cp ../../target/wasm32-unknown-unknown/release/fuel_ee_contracts.wasm ./assets/precompile_$@.wasm
	wasm2wat $(OUT_DIR)/precompile_$@.wasm > $(OUT_DIR)/precompile_$@.wat || true
	cd ../../../fluentbase/bin; $(MAKE) custom_file FILE_IN="$(OUT_DIR)/precompile_$@.wasm" FILE_OUT="$(OUT_DIR)/precompile_$@.rwasm"

.PHONY: abi_definitions
abi_definitions:
	solc --abi assets/solidity/Interfaces.sol -o assets/solidity/generated --overwrite
	cat assets/solidity/generated/FuelEE.abi | jq > assets/solidity/generated/FuelEE.pretty.abi.json
	cat assets/solidity/generated/IFuelEE.abi | jq > assets/solidity/generated/IFuelEE.pretty.abi.json

.PHONY: abi_rs
abi_rs:
	cd ../abi_rs_gen; cargo run -p fuel-ee-abi-rs-gen --profile dev --bin fuel-ee-contracts-abi-rs-gen

.PHONY: abi
abi: abi_definitions abi_rs
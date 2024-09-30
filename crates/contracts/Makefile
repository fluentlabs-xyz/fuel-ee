PRECOMPILES=fvm fvm_codec
CUR_DIR="$(shell pwd)"
OUT_DIR="${CUR_DIR}/assets"
RUSTFLAGS='-C link-arg=-zstack-size=131072 -C target-feature=+bulk-memory'

all: $(PRECOMPILES)

.PHONY: $(PRECOMPILES)
$(PRECOMPILES):
	clear
	mkdir -p $(OUT_DIR)
	RUSTFLAGS=$(RUSTFLAGS) cargo b --release --target=wasm32-unknown-unknown --no-default-features --features=$@
	cp ../../target/wasm32-unknown-unknown/release/fuel_ee_contracts.wasm ./assets/precompile_$@.wasm
	wasm2wat $(OUT_DIR)/precompile_$@.wasm > $(OUT_DIR)/precompile_$@.wat || true
	cd ../../../fluentbase/bin; $(MAKE) custom_file FILE_IN="$(OUT_DIR)/precompile_$@.wasm" FILE_OUT="$(OUT_DIR)/precompile_$@.rwasm"

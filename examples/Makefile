FILES = greeting erc20 hashing panic router rwasm shakmaty storage

all: $(FILES) $(MANUAL)

.PHONY: $(FILES)
$(FILES):
	cd $@ && $(MAKE)

.PHONY: prepare
prepare:
	yarn

.PHONY: clean
clean:
	clear

.PHONE: send_wrapped_fuel_tx
send_wrapped_fuel_tx:
	clear
	node ./send-fuel-tx.js --local

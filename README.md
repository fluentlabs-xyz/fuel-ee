# Example of Fuel smart contract implementation and RPC proxy adapter for Fuel SDK for the Fluent environment

## Semantic units
- Fuel smart contract implementation: crates/contracts
- Fuel adoption for Fluent-EE: crates/contracts
- E2E tests: e2e
- Transaction examples: examples
- GraphQL Proxy adapter implementation for Fuel SDK

## Requirements (not limited to, just tested with)
- Ubuntu 23.10
- GNU Make (tested on version 4.3)
- Rust installation (check 'rust-toolchain' for the version of rustc)
- Golang installation (check 'fuel-proxy/go.mod' file for Golang version)
- NodeJS version v20.12.1
- Yarn version 1.22.22
- Running empty Redis instance (check for dev creds in 'fuel-proxy/config.yaml')
- This project in some local directory
- Fluent node in 'fluent' directory on the same directory level as current project
- 50 GB of free memory

## How to run components and send example Fuel transaction
1. Prepare project: `make prepare`
2. Run Fluent node: `make run_fluent_node`
3. Run Fuel proxy: `make run_fuel_proxy`
4. Send example fuel transaction: `make send_example_fuel_tx`

## How to rebuild Fuel-EE smart contract and update Fluent node with it
use make command ``
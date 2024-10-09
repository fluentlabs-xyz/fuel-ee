use std::path::PathBuf;
use ethers::{
    prelude::{abigen, Abigen},
    providers::{Http, Provider},
    types::Address,
};
use eyre::Result;
use std::sync::Arc;
use std::str::FromStr;

fn main() -> Result<()> {
    rust_file_generation()?;
    Ok(())
}

fn rust_file_generation() -> Result<()> {
    let abi_source = "../contracts/assets/solidity/generated/IFuelEE.abi";
    let out_file = PathBuf::from_str("../../e2e/src/generated/i_fuel_ee.rs").unwrap();
    if out_file.exists() {
        std::fs::remove_file(&out_file)?;
    }
    Abigen::new("IFuelEE", abi_source)?.generate()?.write_to_file(out_file)?;

    Ok(())
}

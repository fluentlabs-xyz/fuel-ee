use std::fs;
use ethers::prelude::Abigen;
use eyre::Result;
use std::path::PathBuf;
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
    let mut abi_gen = Abigen::new("IFuelEE", abi_source)?;
    abi_gen.generate()?.write_to_file(out_file.clone())?;
    let contents = fs::read_to_string(out_file)?;
    let new = contents.replace("aaaa", "../../..").replace("ee", "e");

    Ok(())
}

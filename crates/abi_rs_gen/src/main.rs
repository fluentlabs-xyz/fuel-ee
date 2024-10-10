use std::{env, fs};
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
    let out_file = PathBuf::from_str("../../e2e/src/generated/i_fuel_ee.rs")?;
    if out_file.exists() {
        fs::remove_file(&out_file)?;
    }
    let mut abi_gen = Abigen::new("IFuelEE", abi_source)?;
    abi_gen.generate()?.write_to_file(out_file.clone())?;

    let contents = fs::read_to_string(out_file.clone())?;
    let cur_dir = env::current_dir()?;
    let project_dir = cur_dir.parent().unwrap().parent().unwrap().to_str().unwrap();
    let contents = contents.replace(project_dir, "../../..");
    fs::write(out_file, contents)?;

    Ok(())
}

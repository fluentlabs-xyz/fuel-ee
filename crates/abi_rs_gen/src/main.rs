use std::path::PathBuf;
use ethers::{
    prelude::{abigen, Abigen},
    providers::{Http, Provider},
    types::Address,
};
use eyre::Result;
use std::sync::Arc;
use std::str::FromStr;

// #[tokio::main]
fn main() -> Result<()> {
    rust_file_generation()?;
    // rust_inline_generation().await?;
    // rust_inline_generation_from_abi();
    Ok(())
}

fn rust_file_generation() -> Result<()> {
    let abi_source = "../contracts/assets/solidity/generated/FuelEE.abi";
    // let out_file = std::env::temp_dir().join("FuelEE.abi.rs");
    let out_file = PathBuf::from_str("../contracts/src/generated/fuel_ee.rs").unwrap();
    if out_file.exists() {
        std::fs::remove_file(&out_file)?;
    }
    Abigen::new("FuelEE", abi_source)?.generate()?.write_to_file(out_file)?;
    Ok(())
}

// fn rust_inline_generation_from_abi() {
//     abigen!(IERC20, "./assets/solidity/generated/FuelEE.abi");
// }

// async fn rust_inline_generation() -> Result<()> {
//     // The abigen! macro expands the contract's code in the current scope
//     // so that you can interface your Rust program with the blockchain
//     // counterpart of the contract.
//     abigen!(
//         IERC20,
//         r#"[
//             function totalSupply() external view returns (uint256)
//             function balanceOf(address account) external view returns (uint256)
//             function transfer(address recipient, uint256 amount) external returns (bool)
//             function allowance(address owner, address spender) external view returns (uint256)
//             function approve(address spender, uint256 amount) external returns (bool)
//             function transferFrom( address sender, address recipient, uint256 amount) external returns (bool)
//             event Transfer(address indexed from, address indexed to, uint256 value)
//             event Approval(address indexed owner, address indexed spender, uint256 value)
//         ]"#,
//     );
//
//     const RPC_URL: &str = "https://eth.llamarpc.com";
//     const WETH_ADDRESS: &str = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2";
//
//     let provider = Provider::<Http>::try_from(RPC_URL)?;
//     let client = Arc::new(provider);
//     let address: Address = WETH_ADDRESS.parse()?;
//     let contract = IERC20::new(address, client);
//
//     if let Ok(total_supply) = contract.total_supply().call().await {
//         println!("WETH total supply is {total_supply:?}");
//     }
//
//     Ok(())
// }

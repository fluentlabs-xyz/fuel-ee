#![cfg_attr(target_arch = "wasm32", no_std)]
extern crate alloc;
extern crate core;

// #[cfg(feature = "blended")]
// mod blended;
#[cfg(feature = "fvm")]
mod fvm;

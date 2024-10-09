pub use i_fuel_ee::*;
/// This module was auto-generated with ethers-rs Abigen.
/// More information at: <https://github.com/gakonst/ethers-rs>
#[allow(
    clippy::enum_variant_names,
    clippy::too_many_arguments,
    clippy::upper_case_acronyms,
    clippy::type_complexity,
    dead_code,
    non_camel_case_types,
)]
pub mod i_fuel_ee {
    const _: () = {
        ::core::include_bytes!(
            "/home/bfday/github/fluentlabs-xyz/fuel-ee/crates/contracts/assets/solidity/generated/IFuelEE.abi",
        );
    };
    #[allow(deprecated)]
    fn __abi() -> ::ethers::core::abi::Abi {
        ::ethers::core::abi::ethabi::Contract {
            constructor: ::core::option::Option::None,
            functions: ::core::convert::From::from([
                (
                    ::std::borrow::ToOwned::to_owned("fvmDeposit"),
                    ::std::vec![
                        ::ethers::core::abi::ethabi::Function {
                            name: ::std::borrow::ToOwned::to_owned("fvmDeposit"),
                            inputs: ::std::vec![
                                ::ethers::core::abi::ethabi::Param {
                                    name: ::std::borrow::ToOwned::to_owned("address32"),
                                    kind: ::ethers::core::abi::ethabi::ParamType::FixedArray(
                                        ::std::boxed::Box::new(
                                            ::ethers::core::abi::ethabi::ParamType::Uint(8usize),
                                        ),
                                        32usize,
                                    ),
                                    internal_type: ::core::option::Option::Some(
                                        ::std::borrow::ToOwned::to_owned("uint8[32]"),
                                    ),
                                },
                            ],
                            outputs: ::std::vec![],
                            constant: ::core::option::Option::None,
                            state_mutability: ::ethers::core::abi::ethabi::StateMutability::Payable,
                        },
                    ],
                ),
                (
                    ::std::borrow::ToOwned::to_owned("fvmDryRun"),
                    ::std::vec![
                        ::ethers::core::abi::ethabi::Function {
                            name: ::std::borrow::ToOwned::to_owned("fvmDryRun"),
                            inputs: ::std::vec![
                                ::ethers::core::abi::ethabi::Param {
                                    name: ::std::borrow::ToOwned::to_owned("msg"),
                                    kind: ::ethers::core::abi::ethabi::ParamType::Bytes,
                                    internal_type: ::core::option::Option::Some(
                                        ::std::borrow::ToOwned::to_owned("bytes"),
                                    ),
                                },
                            ],
                            outputs: ::std::vec![],
                            constant: ::core::option::Option::None,
                            state_mutability: ::ethers::core::abi::ethabi::StateMutability::NonPayable,
                        },
                    ],
                ),
                (
                    ::std::borrow::ToOwned::to_owned("fvmExec"),
                    ::std::vec![
                        ::ethers::core::abi::ethabi::Function {
                            name: ::std::borrow::ToOwned::to_owned("fvmExec"),
                            inputs: ::std::vec![
                                ::ethers::core::abi::ethabi::Param {
                                    name: ::std::borrow::ToOwned::to_owned("msg"),
                                    kind: ::ethers::core::abi::ethabi::ParamType::Bytes,
                                    internal_type: ::core::option::Option::Some(
                                        ::std::borrow::ToOwned::to_owned("bytes"),
                                    ),
                                },
                            ],
                            outputs: ::std::vec![],
                            constant: ::core::option::Option::None,
                            state_mutability: ::ethers::core::abi::ethabi::StateMutability::NonPayable,
                        },
                    ],
                ),
                (
                    ::std::borrow::ToOwned::to_owned("fvmWithdraw"),
                    ::std::vec![
                        ::ethers::core::abi::ethabi::Function {
                            name: ::std::borrow::ToOwned::to_owned("fvmWithdraw"),
                            inputs: ::std::vec![
                                ::ethers::core::abi::ethabi::Param {
                                    name: ::std::borrow::ToOwned::to_owned("msg"),
                                    kind: ::ethers::core::abi::ethabi::ParamType::Bytes,
                                    internal_type: ::core::option::Option::Some(
                                        ::std::borrow::ToOwned::to_owned("bytes"),
                                    ),
                                },
                            ],
                            outputs: ::std::vec![],
                            constant: ::core::option::Option::None,
                            state_mutability: ::ethers::core::abi::ethabi::StateMutability::NonPayable,
                        },
                    ],
                ),
            ]),
            events: ::std::collections::BTreeMap::new(),
            errors: ::std::collections::BTreeMap::new(),
            receive: false,
            fallback: false,
        }
    }
    ///The parsed JSON ABI of the contract.
    pub static IFUELEE_ABI: ::ethers::contract::Lazy<::ethers::core::abi::Abi> = ::ethers::contract::Lazy::new(
        __abi,
    );
    pub struct IFuelEE<M>(::ethers::contract::Contract<M>);
    impl<M> ::core::clone::Clone for IFuelEE<M> {
        fn clone(&self) -> Self {
            Self(::core::clone::Clone::clone(&self.0))
        }
    }
    impl<M> ::core::ops::Deref for IFuelEE<M> {
        type Target = ::ethers::contract::Contract<M>;
        fn deref(&self) -> &Self::Target {
            &self.0
        }
    }
    impl<M> ::core::ops::DerefMut for IFuelEE<M> {
        fn deref_mut(&mut self) -> &mut Self::Target {
            &mut self.0
        }
    }
    impl<M> ::core::fmt::Debug for IFuelEE<M> {
        fn fmt(&self, f: &mut ::core::fmt::Formatter<'_>) -> ::core::fmt::Result {
            f.debug_tuple(::core::stringify!(IFuelEE)).field(&self.address()).finish()
        }
    }
    impl<M: ::ethers::providers::Middleware> IFuelEE<M> {
        /// Creates a new contract instance with the specified `ethers` client at
        /// `address`. The contract derefs to a `ethers::Contract` object.
        pub fn new<T: Into<::ethers::core::types::Address>>(
            address: T,
            client: ::std::sync::Arc<M>,
        ) -> Self {
            Self(
                ::ethers::contract::Contract::new(
                    address.into(),
                    IFUELEE_ABI.clone(),
                    client,
                ),
            )
        }
        ///Calls the contract's `fvmDeposit` (0xfcf623ca) function
        pub fn fvm_deposit(
            &self,
            address_32: [u8; 32],
        ) -> ::ethers::contract::builders::ContractCall<M, ()> {
            self.0
                .method_hash([252, 246, 35, 202], address_32)
                .expect("method not found (this should never happen)")
        }
        ///Calls the contract's `fvmDryRun` (0xb8225987) function
        pub fn fvm_dry_run(
            &self,
            msg: ::ethers::core::types::Bytes,
        ) -> ::ethers::contract::builders::ContractCall<M, ()> {
            self.0
                .method_hash([184, 34, 89, 135], msg)
                .expect("method not found (this should never happen)")
        }
        ///Calls the contract's `fvmExec` (0x2ee9d397) function
        pub fn fvm_exec(
            &self,
            msg: ::ethers::core::types::Bytes,
        ) -> ::ethers::contract::builders::ContractCall<M, ()> {
            self.0
                .method_hash([46, 233, 211, 151], msg)
                .expect("method not found (this should never happen)")
        }
        ///Calls the contract's `fvmWithdraw` (0x9429ef5d) function
        pub fn fvm_withdraw(
            &self,
            msg: ::ethers::core::types::Bytes,
        ) -> ::ethers::contract::builders::ContractCall<M, ()> {
            self.0
                .method_hash([148, 41, 239, 93], msg)
                .expect("method not found (this should never happen)")
        }
    }
    impl<M: ::ethers::providers::Middleware> From<::ethers::contract::Contract<M>>
    for IFuelEE<M> {
        fn from(contract: ::ethers::contract::Contract<M>) -> Self {
            Self::new(contract.address(), contract.client())
        }
    }
    ///Container type for all input parameters for the `fvmDeposit` function with signature `fvmDeposit(uint8[32])` and selector `0xfcf623ca`
    #[derive(
        Clone,
        ::ethers::contract::EthCall,
        ::ethers::contract::EthDisplay,
        Default,
        Debug,
        PartialEq,
        Eq,
        Hash
    )]
    #[ethcall(name = "fvmDeposit", abi = "fvmDeposit(uint8[32])")]
    pub struct FvmDepositCall {
        pub address_32: [u8; 32],
    }
    ///Container type for all input parameters for the `fvmDryRun` function with signature `fvmDryRun(bytes)` and selector `0xb8225987`
    #[derive(
        Clone,
        ::ethers::contract::EthCall,
        ::ethers::contract::EthDisplay,
        Default,
        Debug,
        PartialEq,
        Eq,
        Hash
    )]
    #[ethcall(name = "fvmDryRun", abi = "fvmDryRun(bytes)")]
    pub struct FvmDryRunCall {
        pub msg: ::ethers::core::types::Bytes,
    }
    ///Container type for all input parameters for the `fvmExec` function with signature `fvmExec(bytes)` and selector `0x2ee9d397`
    #[derive(
        Clone,
        ::ethers::contract::EthCall,
        ::ethers::contract::EthDisplay,
        Default,
        Debug,
        PartialEq,
        Eq,
        Hash
    )]
    #[ethcall(name = "fvmExec", abi = "fvmExec(bytes)")]
    pub struct FvmExecCall {
        pub msg: ::ethers::core::types::Bytes,
    }
    ///Container type for all input parameters for the `fvmWithdraw` function with signature `fvmWithdraw(bytes)` and selector `0x9429ef5d`
    #[derive(
        Clone,
        ::ethers::contract::EthCall,
        ::ethers::contract::EthDisplay,
        Default,
        Debug,
        PartialEq,
        Eq,
        Hash
    )]
    #[ethcall(name = "fvmWithdraw", abi = "fvmWithdraw(bytes)")]
    pub struct FvmWithdrawCall {
        pub msg: ::ethers::core::types::Bytes,
    }
    ///Container type for all of the contract's call
    #[derive(Clone, ::ethers::contract::EthAbiType, Debug, PartialEq, Eq, Hash)]
    pub enum IFuelEECalls {
        FvmDeposit(FvmDepositCall),
        FvmDryRun(FvmDryRunCall),
        FvmExec(FvmExecCall),
        FvmWithdraw(FvmWithdrawCall),
    }
    impl ::ethers::core::abi::AbiDecode for IFuelEECalls {
        fn decode(
            data: impl AsRef<[u8]>,
        ) -> ::core::result::Result<Self, ::ethers::core::abi::AbiError> {
            let data = data.as_ref();
            if let Ok(decoded) = <FvmDepositCall as ::ethers::core::abi::AbiDecode>::decode(
                data,
            ) {
                return Ok(Self::FvmDeposit(decoded));
            }
            if let Ok(decoded) = <FvmDryRunCall as ::ethers::core::abi::AbiDecode>::decode(
                data,
            ) {
                return Ok(Self::FvmDryRun(decoded));
            }
            if let Ok(decoded) = <FvmExecCall as ::ethers::core::abi::AbiDecode>::decode(
                data,
            ) {
                return Ok(Self::FvmExec(decoded));
            }
            if let Ok(decoded) = <FvmWithdrawCall as ::ethers::core::abi::AbiDecode>::decode(
                data,
            ) {
                return Ok(Self::FvmWithdraw(decoded));
            }
            Err(::ethers::core::abi::Error::InvalidData.into())
        }
    }
    impl ::ethers::core::abi::AbiEncode for IFuelEECalls {
        fn encode(self) -> Vec<u8> {
            match self {
                Self::FvmDeposit(element) => {
                    ::ethers::core::abi::AbiEncode::encode(element)
                }
                Self::FvmDryRun(element) => {
                    ::ethers::core::abi::AbiEncode::encode(element)
                }
                Self::FvmExec(element) => ::ethers::core::abi::AbiEncode::encode(element),
                Self::FvmWithdraw(element) => {
                    ::ethers::core::abi::AbiEncode::encode(element)
                }
            }
        }
    }
    impl ::core::fmt::Display for IFuelEECalls {
        fn fmt(&self, f: &mut ::core::fmt::Formatter<'_>) -> ::core::fmt::Result {
            match self {
                Self::FvmDeposit(element) => ::core::fmt::Display::fmt(element, f),
                Self::FvmDryRun(element) => ::core::fmt::Display::fmt(element, f),
                Self::FvmExec(element) => ::core::fmt::Display::fmt(element, f),
                Self::FvmWithdraw(element) => ::core::fmt::Display::fmt(element, f),
            }
        }
    }
    impl ::core::convert::From<FvmDepositCall> for IFuelEECalls {
        fn from(value: FvmDepositCall) -> Self {
            Self::FvmDeposit(value)
        }
    }
    impl ::core::convert::From<FvmDryRunCall> for IFuelEECalls {
        fn from(value: FvmDryRunCall) -> Self {
            Self::FvmDryRun(value)
        }
    }
    impl ::core::convert::From<FvmExecCall> for IFuelEECalls {
        fn from(value: FvmExecCall) -> Self {
            Self::FvmExec(value)
        }
    }
    impl ::core::convert::From<FvmWithdrawCall> for IFuelEECalls {
        fn from(value: FvmWithdrawCall) -> Self {
            Self::FvmWithdraw(value)
        }
    }
}

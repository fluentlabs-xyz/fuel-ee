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
            "../../../crates/contracts/assets/solidity/generated/IFuelEE.abi",
        );
    };
    #[allow(deprecated)]
    fn __abi() -> ::ethers::core::abi::Abi {
        ::ethers::core::abi::ethabi::Contract {
            constructor: ::core::option::Option::None,
            functions: ::core::convert::From::from([
                (
                    ::std::borrow::ToOwned::to_owned("_stub_1"),
                    ::std::vec![
                        ::ethers::core::abi::ethabi::Function {
                            name: ::std::borrow::ToOwned::to_owned("_stub_1"),
                            inputs: ::std::vec![
                                ::ethers::core::abi::ethabi::Param {
                                    name: ::std::borrow::ToOwned::to_owned("data"),
                                    kind: ::ethers::core::abi::ethabi::ParamType::Tuple(
                                        ::std::vec![
                                            ::ethers::core::abi::ethabi::ParamType::Uint(64usize),
                                            ::ethers::core::abi::ethabi::ParamType::Array(
                                                ::std::boxed::Box::new(
                                                    ::ethers::core::abi::ethabi::ParamType::FixedArray(
                                                        ::std::boxed::Box::new(
                                                            ::ethers::core::abi::ethabi::ParamType::Uint(8usize),
                                                        ),
                                                        34usize,
                                                    ),
                                                ),
                                            ),
                                        ],
                                    ),
                                    internal_type: ::core::option::Option::Some(
                                        ::std::borrow::ToOwned::to_owned("struct FvmWithdrawSol"),
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
                    ::std::borrow::ToOwned::to_owned("fvm_deposit"),
                    ::std::vec![
                        ::ethers::core::abi::ethabi::Function {
                            name: ::std::borrow::ToOwned::to_owned("fvm_deposit"),
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
                    ::std::borrow::ToOwned::to_owned("fvm_dry_run"),
                    ::std::vec![
                        ::ethers::core::abi::ethabi::Function {
                            name: ::std::borrow::ToOwned::to_owned("fvm_dry_run"),
                            inputs: ::std::vec![
                                ::ethers::core::abi::ethabi::Param {
                                    name: ::std::borrow::ToOwned::to_owned("data"),
                                    kind: ::ethers::core::abi::ethabi::ParamType::Array(
                                        ::std::boxed::Box::new(
                                            ::ethers::core::abi::ethabi::ParamType::Uint(8usize),
                                        ),
                                    ),
                                    internal_type: ::core::option::Option::Some(
                                        ::std::borrow::ToOwned::to_owned("uint8[]"),
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
                    ::std::borrow::ToOwned::to_owned("fvm_exec"),
                    ::std::vec![
                        ::ethers::core::abi::ethabi::Function {
                            name: ::std::borrow::ToOwned::to_owned("fvm_exec"),
                            inputs: ::std::vec![
                                ::ethers::core::abi::ethabi::Param {
                                    name: ::std::borrow::ToOwned::to_owned("data"),
                                    kind: ::ethers::core::abi::ethabi::ParamType::Array(
                                        ::std::boxed::Box::new(
                                            ::ethers::core::abi::ethabi::ParamType::Uint(8usize),
                                        ),
                                    ),
                                    internal_type: ::core::option::Option::Some(
                                        ::std::borrow::ToOwned::to_owned("uint8[]"),
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
                    ::std::borrow::ToOwned::to_owned("fvm_withdraw"),
                    ::std::vec![
                        ::ethers::core::abi::ethabi::Function {
                            name: ::std::borrow::ToOwned::to_owned("fvm_withdraw"),
                            inputs: ::std::vec![
                                ::ethers::core::abi::ethabi::Param {
                                    name: ::std::borrow::ToOwned::to_owned("data"),
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
        ///Calls the contract's `_stub_1` (0x90022173) function
        pub fn stub_1(
            &self,
            data: FvmWithdrawSol,
        ) -> ::ethers::contract::builders::ContractCall<M, ()> {
            self.0
                .method_hash([144, 2, 33, 115], (data,))
                .expect("method not found (this should never happen)")
        }
        ///Calls the contract's `fvm_deposit` (0xbb861dbe) function
        pub fn fvm_deposit(
            &self,
            address_32: [u8; 32],
        ) -> ::ethers::contract::builders::ContractCall<M, ()> {
            self.0
                .method_hash([187, 134, 29, 190], address_32)
                .expect("method not found (this should never happen)")
        }
        ///Calls the contract's `fvm_dry_run` (0xf857b299) function
        pub fn fvm_dry_run(
            &self,
            data: ::std::vec::Vec<u8>,
        ) -> ::ethers::contract::builders::ContractCall<M, ()> {
            self.0
                .method_hash([248, 87, 178, 153], data)
                .expect("method not found (this should never happen)")
        }
        ///Calls the contract's `fvm_exec` (0x1da1c731) function
        pub fn fvm_exec(
            &self,
            data: ::std::vec::Vec<u8>,
        ) -> ::ethers::contract::builders::ContractCall<M, ()> {
            self.0
                .method_hash([29, 161, 199, 49], data)
                .expect("method not found (this should never happen)")
        }
        ///Calls the contract's `fvm_withdraw` (0x2f9838af) function
        pub fn fvm_withdraw(
            &self,
            data: ::ethers::core::types::Bytes,
        ) -> ::ethers::contract::builders::ContractCall<M, ()> {
            self.0
                .method_hash([47, 152, 56, 175], data)
                .expect("method not found (this should never happen)")
        }
    }
    impl<M: ::ethers::providers::Middleware> From<::ethers::contract::Contract<M>>
    for IFuelEE<M> {
        fn from(contract: ::ethers::contract::Contract<M>) -> Self {
            Self::new(contract.address(), contract.client())
        }
    }
    ///Container type for all input parameters for the `_stub_1` function with signature `_stub_1((uint64,uint8[34][]))` and selector `0x90022173`
    #[derive(
        Clone,
        ::ethers::contract::EthCall,
        ::ethers::contract::EthDisplay,
        Debug,
        PartialEq,
        Eq,
        Hash
    )]
    #[ethcall(name = "_stub_1", abi = "_stub_1((uint64,uint8[34][]))")]
    pub struct Stub1Call {
        pub data: FvmWithdrawSol,
    }
    ///Container type for all input parameters for the `fvm_deposit` function with signature `fvm_deposit(uint8[32])` and selector `0xbb861dbe`
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
    #[ethcall(name = "fvm_deposit", abi = "fvm_deposit(uint8[32])")]
    pub struct FvmDepositCall {
        pub address_32: [u8; 32],
    }
    ///Container type for all input parameters for the `fvm_dry_run` function with signature `fvm_dry_run(uint8[])` and selector `0xf857b299`
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
    #[ethcall(name = "fvm_dry_run", abi = "fvm_dry_run(uint8[])")]
    pub struct FvmDryRunCall {
        pub data: ::std::vec::Vec<u8>,
    }
    ///Container type for all input parameters for the `fvm_exec` function with signature `fvm_exec(uint8[])` and selector `0x1da1c731`
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
    #[ethcall(name = "fvm_exec", abi = "fvm_exec(uint8[])")]
    pub struct FvmExecCall {
        pub data: ::std::vec::Vec<u8>,
    }
    ///Container type for all input parameters for the `fvm_withdraw` function with signature `fvm_withdraw(bytes)` and selector `0x2f9838af`
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
    #[ethcall(name = "fvm_withdraw", abi = "fvm_withdraw(bytes)")]
    pub struct FvmWithdrawCall {
        pub data: ::ethers::core::types::Bytes,
    }
    ///Container type for all of the contract's call
    #[derive(Clone, ::ethers::contract::EthAbiType, Debug, PartialEq, Eq, Hash)]
    pub enum IFuelEECalls {
        Stub1(Stub1Call),
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
            if let Ok(decoded) = <Stub1Call as ::ethers::core::abi::AbiDecode>::decode(
                data,
            ) {
                return Ok(Self::Stub1(decoded));
            }
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
                Self::Stub1(element) => ::ethers::core::abi::AbiEncode::encode(element),
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
                Self::Stub1(element) => ::core::fmt::Display::fmt(element, f),
                Self::FvmDeposit(element) => ::core::fmt::Display::fmt(element, f),
                Self::FvmDryRun(element) => ::core::fmt::Display::fmt(element, f),
                Self::FvmExec(element) => ::core::fmt::Display::fmt(element, f),
                Self::FvmWithdraw(element) => ::core::fmt::Display::fmt(element, f),
            }
        }
    }
    impl ::core::convert::From<Stub1Call> for IFuelEECalls {
        fn from(value: Stub1Call) -> Self {
            Self::Stub1(value)
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
    ///`FvmWithdrawSol(uint64,uint8[34][])`
    #[derive(
        Clone,
        ::ethers::contract::EthAbiType,
        ::ethers::contract::EthAbiCodec,
        Debug,
        PartialEq,
        Eq,
        Hash
    )]
    pub struct FvmWithdrawSol {
        pub withdraw_amount: u64,
        pub utxo_ids: ::std::vec::Vec<[::ethers::core::types::Uint8; 34]>,
    }
}

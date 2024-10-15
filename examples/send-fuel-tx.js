const {Web3, ETH_DATA_FORMAT} = require('web3');
const {Wallet, Provider, Signer} = require('fuels');
const {BN} = require('@fuel-ts/math');
const {hexToBytes} = require("web3-utils");

const DEPLOYER_PRIVATE_KEY = 'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80';
const FVM_PRECOMPILE_ADDRESS = '0x0000000000000000000000000000000000005250';

function dec2hex(n) {
    let res = n ? [n % 256].concat(dec2hex(~~(n / 256))) : [];
    return res
}

function dec2hexReverse(n) {
    return dec2hex(n).reverse()
}

const main = async () => {
    if (process.argv.length < 2) {
        console.log(`You must specify local or remote flag`);
        console.log(`Example: node send-blended.js --local`);
        process.exit(-1);
    }
    let args = process.argv.slice(2);
    const LOCAL_FUEL_NETWORK_PROXY = 'http://127.0.0.1:8080/v1/graphql'; // proxy
    let fuelProvider = await Provider.create(LOCAL_FUEL_NETWORK_PROXY);
    const checkFlag = (param) => {
        let indexOf = args.indexOf(param)
        if (indexOf < 0) {
            return false
        }
        args.splice(indexOf, 1)
        return true
    };
    let isLocal = checkFlag('--local');
    let isDev = checkFlag('--dev');

    let web3Url = '';
    if (isLocal) {
        web3Url = 'http://127.0.0.1:8545';
    } else if (isDev) {
        web3Url = 'https://rpc.dev.thefluent.xyz/';
    } else {
        console.log(`You must specify --dev or --local flag!`);
        console.log(`Example: node deploy-contract.js --local`);
        process.exit(-1);
    }

    const web3 = new Web3(web3Url);

    const doDepositOnFuel = 1;
    const doSendFuelTx = 1;
    const doWithdrawFromFuel = 1;

    const FvmDepositSig = 3146128830
    let FvmDepositSigBytes = dec2hexReverse(FvmDepositSig)
    const FvmWithdrawSig = 798505135
    let FvmWithdrawSigBytes = dec2hexReverse(FvmWithdrawSig)

    let privateKey = process.env.DEPLOYER_PRIVATE_KEY || DEPLOYER_PRIVATE_KEY;
    let account = web3.eth.accounts.privateKeyToAccount('0x' + privateKey);

    let ethChainId = await web3.eth.getChainId();
    console.log(`ethChainId=${ethChainId}`);
    let fuelChainId = BigInt(fuelProvider.getChainId());
    console.log(`fuelChainId=${fuelChainId}`);
    if (ethChainId !== fuelChainId) {
        throw new Error(`ethChainId(${ethChainId}) !== fuelChainId(${fuelChainId})`)
    }

    let fuelSecretOfficial = "de97d8624a438121b86a1956544bd72ed68cd69f2c99555b08b1e8c51ffd511c";
    let fuelWalletOfficial = Wallet.fromPrivateKey(fuelSecretOfficial, fuelProvider);
    console.log(`fuelWalletOfficial.address`, fuelWalletOfficial.address.toHexString());
    let fuelSecret1 = "0x99e87b0e9158531eeeb503ff15266e2b23c2a2507b138c9d1b1f2ab458df2d61";
    let fuelWallet1 = Wallet.fromPrivateKey(fuelSecret1, fuelProvider);
    console.log(`fuelWallet1.address:`, fuelWallet1.signer().address.toHexString());

    const gasPrice = await web3.eth.getGasPrice(ETH_DATA_FORMAT);

    if (doDepositOnFuel) {
        let ethAmountToDeposit = web3.utils.toWei(1, "ether");
        console.log(`Depositing ${ethAmountToDeposit} Wei from Fluent account ${account.address} to Fuel account ${fuelWalletOfficial.address.toHexString()}`);
        let data = [];
        data = data.concat(FvmDepositSigBytes)
        data = data.concat(...fuelWalletOfficial.address.toBytes())
        let rawTransaction = {
            from: account.address,
            gasPrice: gasPrice,
            gas: 30_000_000,
            to: FVM_PRECOMPILE_ADDRESS,
            value: ethAmountToDeposit,
            data: Buffer.from(data),
        };
        console.log(`Raw transaction:`, rawTransaction)
        let signedTransaction = await web3.eth.accounts.signTransaction(rawTransaction, privateKey)
        console.log(`Sending fuel transaction: ${signedTransaction.transactionHash}`);
        await web3.eth.sendSignedTransaction(signedTransaction.rawTransaction)
            .on('confirmation', confirmation => {
                console.log(`confirmation:`, confirmation)
            })
        console.log(`Balance deposited`);
    }

    if (doSendFuelTx) {
        console.log(`Transferring funds inside Fuel ${fuelWalletOfficial.address.toHexString()}->${fuelWallet1.address.toHexString()}`);
        let fuelTransferTx = await fuelWalletOfficial.createTransfer(fuelWallet1.address, 10);
        fuelTransferTx.maxFee = new BN(100_000);
        console.log(`fuelTransferTx:`, fuelTransferTx);
        // slow send (signing process exposed)
        const fuelTransferFromOfficialToWallet1TxSigned = await fuelWalletOfficial.signTransaction(fuelTransferTx);
        const transactionId = fuelTransferTx.getTransactionId(fuelChainId);
        const recoveredAddress = Signer.recoverAddress(transactionId, fuelTransferFromOfficialToWallet1TxSigned);
        fuelTransferTx.updateWitnessByOwner(recoveredAddress, fuelTransferFromOfficialToWallet1TxSigned);
        let transferResult = await fuelWallet1.sendTransaction(fuelTransferTx);
        console.log(`transferResult.id:`, transferResult.id);
    }

    if (doWithdrawFromFuel) {
        let spendableCoins = await fuelWalletOfficial.getCoins();
        console.log(`spendableCoins:`, spendableCoins);
        if (spendableCoins.coins.length <= 0) {
            throw new Error("user have utxos to spend")
        }
        let utxoIds = [];
        for (let coin of spendableCoins.coins) {
            utxoIds.push(coin.id)
        }
        // TODO take ABI from generated json
        let encodedParams = web3.eth.abi.encodeParameters([
                {
                    "FvmWithdrawInput": {
                        "withdraw_amount": 'uint64',
                        "utxo_ids": 'bytes[]',
                    }
                },
            ],
            [
                {
                    "withdraw_amount": 1,
                    "utxo_ids": utxoIds,
                }
            ],
        );
        let input = Array.from(hexToBytes(encodedParams));
        console.log(`Withdrawing balance from Fuel account ${fuelWalletOfficial.address.toHexString()} to Fluent account ${account.address}`);
        let data = [];
        data = data.concat(FvmWithdrawSigBytes);
        data = data.concat(input);
        let rawTransaction = {
            from: account.address,
            gasPrice: gasPrice,
            gas: 30_000_000,
            to: FVM_PRECOMPILE_ADDRESS,
            data: Buffer.from(data),
        };
        console.log(`rawTransaction:`, rawTransaction);
        let signedTransaction = await web3.eth.accounts.signTransaction(rawTransaction, privateKey);
        console.log("sending transaction");
        await web3.eth.sendSignedTransaction(signedTransaction.rawTransaction)
            .on('confirmation', confirmation => {
                console.log(`confirmation:`, confirmation)
            })
        ;
        console.log(`Balance withdrawn`);
    }
    process.exit(0)
}

main().then(console.log).catch(console.error);

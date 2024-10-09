const {Web3, ETH_DATA_FORMAT} = require('web3');
const {Wallet, Provider, Signer} = require('fuels');
const { BN } = require('@fuel-ts/math');

const DEPLOYER_PRIVATE_KEY = 'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80';
const PRECOMPILE_FVM_ADDRESS = '0x0000000000000000000000000000000000005250';

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
    const checkFlag = (param) => {
        let indexOf = args.indexOf(param)
        if (indexOf < 0) {
            return false
        }
        args.splice(indexOf, 1)
        return true
    };
    let isLocal = checkFlag('--local')
    let isDev = checkFlag('--dev')

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

    const FvmDepositSig = 2311579102
    let FvmDepositSigBytes = dec2hexReverse(FvmDepositSig)
    const FvmWithdrawSig = 3481020119
    let FvmWithdrawSigBytes = dec2hexReverse(FvmWithdrawSig)

    let doSendEthToEthBalance = 0;
    let doDepositEthToFuel = 1;
    let doWithdrawEthFromFuel = 0;
    let doSendFuelTx = 1;

    let fvmPrecompileAddress = "0x0000000000000000000000000000000000005250";

    // let fuelTxOwnerAddress = "0x369f74918912b80c9947d6A174c0C6e2c95fAe1D";
    // let fuelTxOwnerAddressBalance = await web3.eth.getBalance(fuelTxOwnerAddress);
    // console.log(`for fuelTxOwnerAddress ${fuelTxOwnerAddress} balance ${fuelTxOwnerAddressBalance}`);
    let privateKey = process.env.DEPLOYER_PRIVATE_KEY || DEPLOYER_PRIVATE_KEY;
    let account = web3.eth.accounts.privateKeyToAccount('0x' + privateKey);
    // let accountBalance = await web3.eth.getBalance(account.address);
    // console.log(`for account ${account.address} balance ${accountBalance}`);

    if (doSendFuelTx) {
        const LOCAL_FUEL_NETWORK_PROXY = 'http://127.0.0.1:8080/v1/graphql'; // proxy
        const LOCAL_FUEL_NETWORK = 'http://127.0.0.1:4000/v1/graphql';
        // const fuelProviderOriginal = await Provider.create(LOCAL_FUEL_NETWORK);
        const fuelProviderProxy = await Provider.create(LOCAL_FUEL_NETWORK_PROXY);
        // let fuelBaseAssetId = "0xf8f8b6283d7fa5b672b530cbb84fcccb4ff8dc40f8176ef4544ddb1f1952ad07";

        // let baseAssetIdOriginal = fuelProviderOriginal.getBaseAssetId();
        // let baseAssetIdProxy = fuelProviderOriginal.getBaseAssetId();
        // let chainIdOriginal = fuelProviderOriginal.getChainId();
        // let chainIdProxy = fuelProviderOriginal.getChainId();
        // console.log(`baseAssetIdOriginal ${baseAssetIdOriginal}`)
        // console.log(`baseAssetIdProxy ${baseAssetIdProxy}`)
        // console.log(`chainIdOriginal ${chainIdOriginal}`)
        // console.log(`chainIdProxy ${chainIdProxy}`)

        // let fuelTestWallet = await generateTestWallet(fuelProvider, [
        //     [42, baseAssetId],
        // ]);
        // let fuelTestWalletCoins = await fuelProvider.getCoins(fuelTestWallet.address);
        // console.log(`fuelTestWalletCoins`, fuelTestWalletCoins);
        // let fuelTestWalletBalance = await fuelTestWallet.getBalance();
        // console.log(`fuelWalletOfficialBalance ${fuelTestWalletBalance}`);

        // let fuelSecretOfficial = "a1447cd75accc6b71a976fd3401a1f6ce318d27ba660b0315ee6ac347bf39568";
        // let fuelWalletOfficial = Wallet.fromPrivateKey(fuelSecretOfficial, fuelProvider);

        // let fuelProvider = await Provider.create(LOCAL_FUEL_NETWORK);
        let fuelProvider = fuelProviderProxy;

        let ethChainId = await web3.eth.getChainId();
        console.log(`ethChainId=${ethChainId}`);
        let fuelChainId = BigInt(fuelProvider.getChainId());
        console.log(`fuelChainId=${fuelChainId}`);
        if (ethChainId !== fuelChainId) {
            throw new Error(`ethChainId(${ethChainId}) !== fuelChainId(${fuelChainId})`)
        }

        let fuelSecretOfficial = "de97d8624a438121b86a1956544bd72ed68cd69f2c99555b08b1e8c51ffd511c";
        let fuelWalletOfficial = Wallet.fromPrivateKey(fuelSecretOfficial, fuelProvider);
        console.log(`- fuelWalletOfficial.address`, fuelWalletOfficial.address.toHexString());

        let fuelSecret1 = "0x99e87b0e9158531eeeb503ff15266e2b23c2a2507b138c9d1b1f2ab458df2d61";
        let fuelWallet1 = Wallet.fromPrivateKey(fuelSecret1, fuelProvider);
        console.log(`- fuelWallet1.address:`, fuelWallet1.signer().address.toHexString());

        if (doDepositEthToFuel) {
            console.log(`sending balance to ${account.address}->${fuelWalletOfficial.address.toHexString()}`)
            const gasPrice = await web3.eth.getGasPrice(ETH_DATA_FORMAT);
            let ethAmountToSend = web3.utils.toWei(300, "ether");
            let data = [];
            data = data.concat(FvmDepositSigBytes)
            data = data.concat(...fuelWalletOfficial.address.toBytes())

            let rawTransaction = {
                from: account.address,
                gasPrice: gasPrice,
                gas: 300_000_000,
                to: fvmPrecompileAddress,
                value: ethAmountToSend,
                data: Buffer.from(data),
            };
            console.log(`ethAmountToSend:`, ethAmountToSend)
            console.log(`rawTransaction:`, rawTransaction)
            const signedTransaction = await web3.eth.accounts.signTransaction(rawTransaction, privateKey)
            console.log("sending fuel transaction");
            await web3.eth.sendSignedTransaction(signedTransaction.rawTransaction)
                .on('confirmation', confirmation => {
                    console.log(`confirmation:`, confirmation)
                })
            ;
            console.log(`balance sent`);
            process.exit(0)
        }

        if (doWithdrawEthFromFuel) {
            console.log(`sending balance to ${fuelWalletOfficial.address.toHexString()}->${account.address}`)
            const gasPrice = await web3.eth.getGasPrice(ETH_DATA_FORMAT);
            let ethAmountToSend = web3.utils.toWei(0.01, "ether");
            let data = [];
            data = data.concat(FvmWithdrawSigBytes)
            console.log(`data: ${Buffer.from(data).toString('hex')}`)

            let rawTransaction = {
                from: account.address,
                gasPrice: gasPrice,
                gas: 300_000_000,
                to: fvmPrecompileAddress,
                value: ethAmountToSend,
                data: Buffer.from(data),
                // "chainId": 1337 // Remember to change this
            };
            console.log(`ethAmountToSend:`, ethAmountToSend)
            console.log(`rawTransaction:`, rawTransaction)
            const signedTransaction = await web3.eth.accounts.signTransaction(rawTransaction, privateKey)
            console.log("sending transaction");
            await web3.eth.sendSignedTransaction(signedTransaction.rawTransaction)
                .on('confirmation', confirmation => {
                    console.log(`confirmation:`, confirmation)
                })
            ;
            console.log(`balance sent`);
        }

        // let fuelWalletOfficialCoins = await fuelProvider.getCoins(
        //     fuelWalletOfficial.address,
        //     // "0xf8f8b6283d7fa5b672b530cbb84fcccb4ff8dc40f8176ef4544ddb1f1952ad07",
        // );
        // console.log(`- fuelWalletOfficialCoins:`, fuelWalletOfficialCoins);
        // process.exit(0)

        console.log("- fuel: creating transfer");
        let fuelTransferFromOfficialToWallet1Tx = await fuelWalletOfficial.createTransfer(fuelWallet1.address, 10);
        fuelTransferFromOfficialToWallet1Tx.maxFee = new BN(100_000);
        console.log("- fuelTransferFromOfficialToWallet1Tx:", fuelTransferFromOfficialToWallet1Tx);

        if (false) {
            // fast send (signing inside SDK?)
            let transferResult = await fuelWallet1.sendTransaction(fuelTransferFromOfficialToWallet1Tx);
            console.log(`- transferResult`, transferResult);
            console.log(`- transferResult.id`, transferResult.id);
        } else {
            // slow send (signing process exposed)
            const fuelTransferFromOfficialToWallet1TxSigned = await fuelWalletOfficial.signTransaction(fuelTransferFromOfficialToWallet1Tx);
            console.log("- fuelTransferFromOfficialToWallet1TxSigned:", fuelTransferFromOfficialToWallet1TxSigned);
            const transactionId = fuelTransferFromOfficialToWallet1Tx.getTransactionId(fuelChainId);
            console.log("- transactionId:", transactionId);
            const recoveredAddress = Signer.recoverAddress(transactionId, fuelTransferFromOfficialToWallet1TxSigned);
            console.log("- recoveredAddress:", recoveredAddress);
            fuelTransferFromOfficialToWallet1Tx.updateWitnessByOwner(recoveredAddress, fuelTransferFromOfficialToWallet1TxSigned);
            console.log("- fuelTransferFromOfficialToWallet1Tx (updated witness):", fuelTransferFromOfficialToWallet1Tx);
            let transferResult = await fuelWallet1.sendTransaction(fuelTransferFromOfficialToWallet1Tx);
            console.log(`- transferResult`, transferResult);
            console.log(`- transferResult.id`, transferResult.id);
            // let {id} = await transferResult.wait();
            // console.log(`- transfer id`, id);
        }


        // fuelWalletOfficialCoins = await fuelProvider.getCoins(fuelWalletOfficial.address);
        // console.log(`- fuelWalletOfficialCoins:`, fuelWalletOfficialCoins);
        // fuelWallet1Coins = await fuelProvider.getCoins(fuelWallet1.address);
        // console.log(`- fuelWallet1Coins:`, fuelWallet1Coins);

        process.exit(0)
    }

    process.exit(0)
}

main().then(console.log).catch(console.error);

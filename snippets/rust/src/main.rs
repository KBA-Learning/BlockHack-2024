use alloy::{
    providers::{Provider, ProviderBuilder},
    signers::local::PrivateKeySigner,
    sol,
};
use eyre::Result;

sol!(
    #[sol(rpc)]
    Storage,
    "../../common/contracts/Storage.abi"
);

#[tokio::main]
async fn main() -> Result<()> {
    let rpc_url = "http://127.0.0.1:8545".parse()?;
    let provider = ProviderBuilder::new().on_http(rpc_url);

    let signer: PrivateKeySigner =
        "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
            .parse()
            .expect("should parse private key");
    println!("Address: \x1b[32m{}\x1b[0m", signer.address());

    let balance = provider.get_balance(signer.address()).await?;
    println!("Balance: \x1b[32m{balance}\x1b[0m");

    let latest_block = provider.get_block_number().await?;
    println!("Latest Block: \x1b[32m{latest_block}\x1b[0m");

    let contract = Storage::new(
        "0x5FbDB2315678afecb367f032d93F642f64180aa3".parse()?,
        provider,
    );

    let builder = contract.store(String::from("101"), String::from("Hello, KBA!"));
    let trx = builder.send().await?;
    println!("Transaction Hash: \x1b[32m{}\x1b[0m", trx.tx_hash());

    let result = contract.collection(String::from("101")).call().await?;
    println!("Value: \x1b[32m{}\x1b[0m", result._0);

    Ok(())
}

package org.web3j;

import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.DefaultBlockParameterName;
import org.web3j.protocol.http.HttpService;
import org.web3j.tx.gas.DefaultGasProvider;

public class Web3App {
      public static final String GREEN = "\033[0;32m";
      public static final String RESET = "\033[0m";

      public static void main(String[] args) throws Exception {
            Credentials credentials = Credentials
                        .create("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80");
            System.out.println("Address: " + GREEN + credentials.getAddress() + RESET);
            Web3j web3j = Web3j.build(new HttpService("http://127.0.0.1:8545"));
            System.out.println("Balance: " + GREEN
                        + web3j.ethGetBalance(credentials.getAddress(), DefaultBlockParameterName.LATEST).send()
                                    .getBalance()
                        + RESET);
            System.out.println("Latest Block: " + GREEN + web3j.ethBlockNumber().send().getBlockNumber() + RESET);
            Storage storage = Storage.load("0x5FbDB2315678afecb367f032d93F642f64180aa3", web3j, credentials,
                        new DefaultGasProvider());
            System.out.println("Transaction Hash: " + GREEN
                        + storage.store("1995", "Hello from Java ☕!").send().getTransactionHash() + RESET);
            System.out.println("Value: " + GREEN + storage.collection("1995").send() + RESET);
      }
}

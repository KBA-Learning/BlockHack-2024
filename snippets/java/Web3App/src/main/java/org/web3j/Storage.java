package org.web3j;

import io.reactivex.Flowable;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import org.web3j.abi.EventEncoder;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Event;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.Utf8String;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.DefaultBlockParameter;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.core.RemoteFunctionCall;
import org.web3j.protocol.core.methods.request.EthFilter;
import org.web3j.protocol.core.methods.response.BaseEventResponse;
import org.web3j.protocol.core.methods.response.Log;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;
import org.web3j.tx.gas.ContractGasProvider;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/hyperledger/web3j/tree/main/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 1.6.1.
 */
@SuppressWarnings("rawtypes")
public class Storage extends Contract {
    public static final String BINARY = "608060405234801561000f575f80fd5b506107ae8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063dc642f3014610038578063f641090c14610068575b5f80fd5b610052600480360381019061004d91906102ff565b610084565b60405161005f91906103c0565b60405180910390f35b610082600480360381019061007d91906103e0565b610136565b005b5f818051602081018201805184825260208301602085012081835280955050505050505f9150905080546100b790610483565b80601f01602080910402602001604051908101604052809291908181526020018280546100e390610483565b801561012e5780601f106101055761010080835404028352916020019161012e565b820191905f5260205f20905b81548152906001019060200180831161011157829003601f168201915b505050505081565b805f8360405161014691906104ed565b9081526020016040518091039020908161016091906106a9565b508160405161016f91906104ed565b60405180910390207f05cf4538ca45b8f0999a133df2f77a0e5648ca2349b8df20169cbc6328fdae99826040516101a691906103c0565b60405180910390a25050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610211826101cb565b810181811067ffffffffffffffff821117156102305761022f6101db565b5b80604052505050565b5f6102426101b2565b905061024e8282610208565b919050565b5f67ffffffffffffffff82111561026d5761026c6101db565b5b610276826101cb565b9050602081019050919050565b828183375f83830152505050565b5f6102a361029e84610253565b610239565b9050828152602081018484840111156102bf576102be6101c7565b5b6102ca848285610283565b509392505050565b5f82601f8301126102e6576102e56101c3565b5b81356102f6848260208601610291565b91505092915050565b5f60208284031215610314576103136101bb565b5b5f82013567ffffffffffffffff811115610331576103306101bf565b5b61033d848285016102d2565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561037d578082015181840152602081019050610362565b5f8484015250505050565b5f61039282610346565b61039c8185610350565b93506103ac818560208601610360565b6103b5816101cb565b840191505092915050565b5f6020820190508181035f8301526103d88184610388565b905092915050565b5f80604083850312156103f6576103f56101bb565b5b5f83013567ffffffffffffffff811115610413576104126101bf565b5b61041f858286016102d2565b925050602083013567ffffffffffffffff8111156104405761043f6101bf565b5b61044c858286016102d2565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061049a57607f821691505b6020821081036104ad576104ac610456565b5b50919050565b5f81905092915050565b5f6104c782610346565b6104d181856104b3565b93506104e1818560208601610360565b80840191505092915050565b5f6104f882846104bd565b915081905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261055f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610524565b6105698683610524565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6105ad6105a86105a384610581565b61058a565b610581565b9050919050565b5f819050919050565b6105c683610593565b6105da6105d2826105b4565b848454610530565b825550505050565b5f90565b6105ee6105e2565b6105f98184846105bd565b505050565b5b8181101561061c576106115f826105e6565b6001810190506105ff565b5050565b601f8211156106615761063281610503565b61063b84610515565b8101602085101561064a578190505b61065e61065685610515565b8301826105fe565b50505b505050565b5f82821c905092915050565b5f6106815f1984600802610666565b1980831691505092915050565b5f6106998383610672565b9150826002028217905092915050565b6106b282610346565b67ffffffffffffffff8111156106cb576106ca6101db565b5b6106d58254610483565b6106e0828285610620565b5f60209050601f831160018114610711575f84156106ff578287015190505b610709858261068e565b865550610770565b601f19841661071f86610503565b5f5b8281101561074657848901518255600182019150602085019450602081019050610721565b86831015610763578489015161075f601f891682610672565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220db6c60d8d8359a6f32ab1ee5cac909917d78fe1828b1459d0f3aec545e4316d464736f6c63430008180033";

    private static String librariesLinkedBinary;

    public static final String FUNC_STORE = "store";

    public static final String FUNC_COLLECTION = "collection";

    public static final Event STORED_EVENT = new Event("Stored", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>(true) {}, new TypeReference<Utf8String>() {}));
    ;

    @Deprecated
    protected Storage(String contractAddress, Web3j web3j, Credentials credentials,
            BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected Storage(String contractAddress, Web3j web3j, Credentials credentials,
            ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected Storage(String contractAddress, Web3j web3j, TransactionManager transactionManager,
            BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected Storage(String contractAddress, Web3j web3j, TransactionManager transactionManager,
            ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteFunctionCall<TransactionReceipt> store(String _key, String _value) {
        final Function function = new Function(
                FUNC_STORE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Utf8String(_key), 
                new org.web3j.abi.datatypes.Utf8String(_value)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public static List<StoredEventResponse> getStoredEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = staticExtractEventParametersWithLog(STORED_EVENT, transactionReceipt);
        ArrayList<StoredEventResponse> responses = new ArrayList<StoredEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            StoredEventResponse typedResponse = new StoredEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.key = (byte[]) eventValues.getIndexedValues().get(0).getValue();
            typedResponse.value = (String) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public static StoredEventResponse getStoredEventFromLog(Log log) {
        Contract.EventValuesWithLog eventValues = staticExtractEventParametersWithLog(STORED_EVENT, log);
        StoredEventResponse typedResponse = new StoredEventResponse();
        typedResponse.log = log;
        typedResponse.key = (byte[]) eventValues.getIndexedValues().get(0).getValue();
        typedResponse.value = (String) eventValues.getNonIndexedValues().get(0).getValue();
        return typedResponse;
    }

    public Flowable<StoredEventResponse> storedEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(log -> getStoredEventFromLog(log));
    }

    public Flowable<StoredEventResponse> storedEventFlowable(DefaultBlockParameter startBlock,
            DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(STORED_EVENT));
        return storedEventFlowable(filter);
    }

    public RemoteFunctionCall<String> collection(String param0) {
        final Function function = new Function(FUNC_COLLECTION, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Utf8String(param0)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    @Deprecated
    public static Storage load(String contractAddress, Web3j web3j, Credentials credentials,
            BigInteger gasPrice, BigInteger gasLimit) {
        return new Storage(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static Storage load(String contractAddress, Web3j web3j,
            TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new Storage(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static Storage load(String contractAddress, Web3j web3j, Credentials credentials,
            ContractGasProvider contractGasProvider) {
        return new Storage(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static Storage load(String contractAddress, Web3j web3j,
            TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new Storage(contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<Storage> deploy(Web3j web3j, Credentials credentials,
            ContractGasProvider contractGasProvider) {
        return deployRemoteCall(Storage.class, web3j, credentials, contractGasProvider, getDeploymentBinary(), "");
    }

    @Deprecated
    public static RemoteCall<Storage> deploy(Web3j web3j, Credentials credentials,
            BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(Storage.class, web3j, credentials, gasPrice, gasLimit, getDeploymentBinary(), "");
    }

    public static RemoteCall<Storage> deploy(Web3j web3j, TransactionManager transactionManager,
            ContractGasProvider contractGasProvider) {
        return deployRemoteCall(Storage.class, web3j, transactionManager, contractGasProvider, getDeploymentBinary(), "");
    }

    @Deprecated
    public static RemoteCall<Storage> deploy(Web3j web3j, TransactionManager transactionManager,
            BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(Storage.class, web3j, transactionManager, gasPrice, gasLimit, getDeploymentBinary(), "");
    }

    public static void linkLibraries(List<Contract.LinkReference> references) {
        librariesLinkedBinary = linkBinaryWithReferences(BINARY, references);
    }

    private static String getDeploymentBinary() {
        if (librariesLinkedBinary != null) {
            return librariesLinkedBinary;
        } else {
            return BINARY;
        }
    }

    public static class StoredEventResponse extends BaseEventResponse {
        public byte[] key;

        public String value;
    }
}

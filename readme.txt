### Actual Lition ERC20 token contract
```
https://ropsten.etherscan.io/address/0x65fc0f7d2bb96a9be30a770fb5fcd5a7762ad659
```

### Actual Lition Registry contract
```
https://ropsten.etherscan.io/address/0xA5276B814DfBEd36Fd9bF9ecF6463BBfcbB705ef
```

### Actual Lition Validator contract
```
https://ropsten.etherscan.io/address/0xc2ff80ba9aba34b0e9b9a4aca41a5db6e4e4c8ba
```

### Contract debugging
To get require error message, use:
https://gist.github.com/msigwart/d3e374a64c8718f8ac5ec04b5093597f

# go-lition-abi
#### Prerequisites
##### solc
sudo add-apt-repository ppa:ethereum/ethereum  
sudo apt-get update  
sudo apt-get install solc  

##### abigen
go get -u github.com/ethereum/go-ethereum  
cd $GOPATH/src/github.com/ethereum/go-ethereum/  
make  
make devtools  

#### Create SC ABI
run
```
cd contracts
solc --abi --overwrite --optimize Registration.sol --output-dir client/abi
```

#### Create SC go class 
run
```
cd contracts
abigen --abi=client/abi/LitionRegistry.abi --pkg=litionScClient --out=client/litionRegistry.go
// Replace imports in client/litionRegistry.go from "github.com/ethereum/go-ethereum" to "gitlab.com/lition/lition"
```

## How to deploy contracts
- Create Ropsten account at Metamask and mint some ethers. https://metamask.io  
- Create Etherscan account and create API key. https://etherscan.io  
- Open Solidity online IDE: https://remix.ethereum.org  
- Activate "Etherscan - Contract Verification" plugin  
- Copy your Solidity code into Remix and compile it with the latest compiler  
- When deploying, use Injected Web3 as environment + select your ropsten metamask account  
- Deploy selected contract, e.g. LitionRegistry - as constructor argument use address of contract with ERC20 LIT token. Use existing contract (0x65fc0f7d2bb96a9be30a770fb5fcd5a7762ad659) or deploy new one.  
- You can interact with contract directly throught remix after clicking on LitionRegistry under Deployed contracts.   
- If you want to publish contract interface also on Etherscan, verify contract through "Contract Verification Plugin".   
  You need to provide API key from etherscan, contract name, newly deployed LitionRegistry contract address and constructor arguments in ABI encoded format. To get this format, use: https://abi.hashex.org  
- Copy there ABI - can by copied from compilation step in Remix. Put there also constructor argument and copy the encoded result into remix. You can now Verify Contract and contract interface will be available also on etherscan.


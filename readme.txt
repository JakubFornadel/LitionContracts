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
solc --abi Registration.sol --output-dir client/abi
```

#### Create SC go class 
run
```
cd contracts
abigen --abi=client/abi/LitionRegistry.abi --pkg=litionContractClient --out=client/litionRegistry.go
```


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
cd abi
solc --abi ../Registration.sol -o .
abigen --abi=LitionRegistry.abi --pkg=lition --out=LitionRegistry.go
```

#### Create SC go class 
run
```
cd go_wrapper
abigen --abi=../abi/LitionRegistry.abi --pkg=lition --out=LitionRegistry.go
```


package payment

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"codedaotoken/contract"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

const (
	addrContract = "0xdccc2dd10c7b3be9661aecf87b99f65f290c3f1f"
	addrOwner = "0x3A120D99ac41D6e0630382752F9714D8772c5432"
	// import account to unlock in geth, file keyjson in keystore
	key = `{"address":"3a120d99ac41d6e0630382752f9714d8772c5432","crypto":{"cipher":"aes-128-ctr","ciphertext":"eca3d8f8eb49b969eb17ec58b486cf4d86162cac6007f9a52aedc7c616c14830","cipherparams":{"iv":"475a6ef6eb8ca5a546b78a733c6a7bc9"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"960c0080b60884d1b9c7cbb1e655025da4fb1f9420ef945e77368ca9ff25f285"},"mac":"cb2b4a000e9ec1b3939b60cf9f92472b90bc72379830635c4a3b491234ca3916"},"id":"45185a27-5e40-4839-81f5-b052b0d487d2","version":3}`
	//pathIPC = "/home/vanthanhbk/.ethereum/rinkeby/geth.ipc"
	pathIPC = "/home/phanvanthanh_mrt/.ethereum/rinkeby/geth.ipc"
)

func GetClient () (*ethclient.Client, error) {
	return ethclient.Dial(pathIPC)
}

func GetContract () *contract.CodedaoNetwork {
	client, err := GetClient()
	token, err := contract.NewCodedaoNetwork(common.HexToAddress(addrContract), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	return token
}

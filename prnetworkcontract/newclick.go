package prnetworkcontract

import (
	"codedaotoken/payment"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"log"
	"github.com/ethereum/go-ethereum/common"
)

var (
	auth *bind.TransactOpts
	err error
	proofkey = "24e890c939b2f19273a2843f0ba1836c"
)

const (
	key = `{"address":"3a120d99ac41d6e0630382752f9714d8772c5432","crypto":{"cipher":"aes-128-ctr","ciphertext":"a844aa4d9aea200f8b914b8b55f7301a789dd306cefbe9770b114aa2949bfa9d","cipherparams":{"iv":"f488e7f5c5de84f18f9978280e428aa1"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ad69abe52dfa549b5d6c444cef3da08a8167e2f639c4e1d0d68a591039d205ed"},"mac":"690b5c57a539a5074abae6f87dd8d1094d6f76c52be971286e0698550d24d77d"},"id":"be3303af-83c3-4555-a10b-1bcf4920b659","version":3}`
)

func init(){
	auth, err = bind.NewTransactor(strings.NewReader(key), "thanhpv1234")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
}

func OnNewClick(from string, parent string, contract string) string {
	prContract := payment.GetPrContract(contract)
	tx, err := prContract.OnNewClick(auth, common.HexToAddress(from), common.HexToAddress(parent))
	if err != nil {
		log.Fatalf("Failed to connect to request OnNewClick: %v", err)
	}
	return tx.Hash().String()
}

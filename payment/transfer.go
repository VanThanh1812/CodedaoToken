package payment

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"log"
	"math/big"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"codedaotoken/models"
	"github.com/ethereum/go-ethereum/core/types"
	"errors"
	"regexp"
)

var (
	auth *bind.TransactOpts
	err error
	proofkey = "24e890c939b2f19273a2843f0ba1836c"
)

func init(){
	auth, err = bind.NewTransactor(strings.NewReader(key), "thanhpv1234")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
}

func checkAddress(addr string) bool {
	match, _ := regexp.MatchString("^0x[a-fA-F0-9]{40}$", addr)
	return match
}

func Transfer(proofkey string, rq models.TransferRequest)(*types.Transaction, error){

	if !checkProofKey(proofkey) {
		return nil, errors.New("Proof key error.")
	}

	if !checkAddress(rq.AddrTo){
		return nil, errors.New("Address To wrong.")
	}

	if rq.Amount <= 0 {
		return nil, errors.New("Amount require more than 0")
	}

	tx, err := tokenContract.Transfer(auth, common.HexToAddress(rq.AddrTo), big.NewInt(rq.Amount))
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
		return nil, err
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	return tx, nil
}

func TransferFrom(proofkey string, rq models.TransferFromRequest)(*types.Transaction, error){
	if !checkProofKey(proofkey) {
		return nil, errors.New("Proof key error.")
	}

	if !checkAddress(rq.AddrTo){
		return nil, errors.New("Address To wrong.")
	}

	if !checkAddress(rq.AddrFrom){
		return nil, errors.New("Address From wrong.")
	}

	if rq.Amount <= 0 {
		return nil, errors.New("Amount require more than 0")
	}

	tx, err := tokenContract.TransferFrom(auth, common.HexToAddress(rq.AddrFrom),common.HexToAddress(rq.AddrTo), big.NewInt(rq.Amount))
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
		return nil, err
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	return tx, nil
}

func Earn(proofkey string, eq models.EarnRequest)(*types.Transaction, error){
	if !checkProofKey(proofkey) {
		return nil, errors.New("Proof key error.")
	}

	if !checkAddress(eq.AddrTo){
		return nil, errors.New("Address To wrong.")
	}

	if eq.Amount <= 0 {
		return nil, errors.New("Amount require more than 0")
	}

	tx, err := tokenContract.Earn(auth, common.HexToAddress(eq.AddrTo), big.NewInt(eq.Amount))
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
		return nil, err
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	return tx, nil
}

func checkProofKey(key string) bool {
	if key == proofkey {
		return true
	}else{
		return false
	}
}
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
)

var (
	auth *bind.TransactOpts
)

func init(){
	auth, _ = bind.NewTransactor(strings.NewReader(key), "thanhpv1234")
}

func Transfer(rq models.TransferRequest)(*types.Transaction, error){
	tx, err := tokenContract.Transfer(auth, common.HexToAddress(rq.AddrTo), big.NewInt(rq.Amount))
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
		return nil, err
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	return tx, nil
}

func TransferFrom(rq models.TransferFromRequest)(*types.Transaction, error){
	tx, err := tokenContract.TransferFrom(auth, common.HexToAddress(rq.AddrFrom),common.HexToAddress(rq.AddrTo), big.NewInt(rq.Amount))
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
		return nil, err
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	return tx, nil
}

func Earn(eq models.EarnRequest)(*types.Transaction, error){
	tx, err := tokenContract.Earn(auth, common.HexToAddress(eq.AddrTo), big.NewInt(eq.Amount))
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
		return nil, err
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	return tx, nil
}
package payment

import (
	"codedaotoken/models"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"codedaotoken/contract"
)

var (
	tokenContract *contract.CodedaoNetwork
)

func init(){
	tokenContract = GetContract()
}

func GetOwner () (models.Address) {
	owner, _ := tokenContract.Owner(nil)
	o := models.Address{
		Address: owner.String(),
	}
	return o
}

func GetBalanceOf (addr string) (*models.BalanceInfo, error){
	balance, err := tokenContract.BalanceOf(nil, common.HexToAddress(addr))
	if err != nil{
		return nil, err
	}else{
		return &models.BalanceInfo{Addr:addr, Balance: balance}, nil
	}
}

func GetTotalSupply()(*big.Int){
	total, _ := tokenContract.TotalSupply(nil)
	return total
}
package models

import (
	"math/big"
)

type Response struct {
	Data interface{} `json:"data"`
	Error ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	Code 	int16	`json:"code"`
	Message interface{}	`json:"message"`
}

type Address struct {
	Address string 	`json:"address"`
}

type BalanceInfo struct {
	Addr string 	`json:"address"`
	Balance *big.Int 		`json:"balance"`
}

type TransferRequest struct {
	AddrTo string	`json:"addr_to"`
	Amount int64		`json:"amount"`
}

type TransferFromRequest struct {
	AddrFrom 	string 	`json:"addr_from"`
	AddrTo 		string 	`json:"addr_to"`
	Amount 		int64 	`json:"amount"`
}

type EarnRequest struct {
	AddrTo string	`json:"addr_to"`
	Amount int64		`json:"amount"`
}
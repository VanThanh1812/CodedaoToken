package main

import (
	_ "codedaotoken/routers"
	"github.com/astaxie/beego"
	"github.com/gorilla/mux"
	"net/http"
	"codedaotoken/prnetworkcontract"
)

func main() {



	/*token, err := contract.NewCodedaoNetwork(common.HexToAddress(addrContract), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	name, err := token.Name(nil)

	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Token name:", name)

	// authorized
	auth, err := bind.NewTransactor(strings.NewReader(key), "thanhpv1234")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	tx, err := token.Transfer(auth, common.HexToAddress("0x9bb555f21F17A694420bb8D84dFB11491bd92929"), big.NewInt(10))
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())*/

	/*session := &contract.CodedaoNetworkSession{
		Contract:token,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:auth.From,
			Signer:auth.Signer,
			GasLimit: 200000,
			GasPrice:big.NewInt(21),
		},
	}

	tx, err := session.Transfer(common.HexToAddress("0x9bb555f21F17A694420bb8D84dFB11491bd92929"), big.NewInt(10))
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())*/

	// api
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()

	// redirect
	r := mux.NewRouter()
	r.HandleFunc("/click", OnNewClick)
}

func OnNewClick(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	from := request.URL.Query().Get("from")
	parent := request.URL.Query().Get("parent")
	contract := request.URL.Query().Get("contract")
	link := request.URL.Query().Get("ref")

	txhash := prnetworkcontract.OnNewClick(from, parent, contract)

	http.Redirect(writer, request,link, 301)

	writer.Write([]byte(txhash))

}

package controllers

import (
	"github.com/astaxie/beego"
	"codedaotoken/payment"
	"codedaotoken/models"
	"encoding/json"
)

// PaymentController operations for Payment
type PaymentController struct {
	beego.Controller
}

// URLMapping ...
func (c *PaymentController) URLMapping() {
	c.Mapping("GetOwner", c.GetOwner)
	c.Mapping("GetBalanceOf", c.GetBalanceOf)
	c.Mapping("GetTotalSupply", c.GetTotalSupply)
	c.Mapping("Transfer", c.Transfer)
	c.Mapping("TransferFrom", c.TransferFrom)
	c.Mapping("Earn", c.Earn)
}

// GetOwner ...
// @Title GetOwner
// @Description get Owner Contract
// @Success 200 {object} models.Address
// @Failure 403  is empty
// @router /owner [get]
func (c *PaymentController) GetOwner(){
	res := models.Response{
		Data: payment.GetOwner(),
		Error: models.ErrorResponse{
			Message:"Success",
			Code:200,
		},
	}
	c.Data["json"] = res
	c.ServeJSON()
}

// GetBalanceOf by address ...
// @Title GetBalanceof
// @Description get Balance of Address
// @Param	address		query 	string	true		"the address you want to get"
// @Success 200 {object} models.Response
// @Failure 403  is error
// @router /balanceof [get]
func (c *PaymentController) GetBalanceOf(){
	addr := c.Ctx.Input.Query("address")

	if addr == "" {
		return
	}

	balance, err := payment.GetBalanceOf(addr)

	var res models.Response

	if err == nil{
		res = models.Response{
			Data: balance,
			Error: models.ErrorResponse{
				Message:"Success",
				Code:200,
			},
		}
	}else{
		res = models.Response{
			Data: 0,
			Error: models.ErrorResponse{
				Message: "Success",
				Code:403,
			},
		}
	}

	c.Data["json"] = res
	c.ServeJSON()
}

// GetTotalSupply ...
// @Title GetTotalSupply
// @Description get total Token Contract
// @Success 200 {object} big.Int
// @Failure 403  is empty
// @router /totalsupply [get]
func (c *PaymentController)GetTotalSupply()(){
	res := models.Response{
		Data: payment.GetTotalSupply(),
		Error: models.ErrorResponse{
			Message:"Success",
			Code:200,
		},
	}
	c.Data["json"] = res
	c.ServeJSON()
}

// Post ...
// @Title Send amount CNC
// @Description create request tranfer
// @Param	body		body 	models.TransferRequest	true		"body for Payment content"
// @Success 201 {object} models.Response
// @Failure 403 body is empty
// @router /transfer [post]
func (c *PaymentController) Transfer() {
	var rq models.TransferRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &rq)

	var res models.Response
	tx, err := payment.Transfer(rq)
	if err == nil{
		res = models.Response{
			Data: tx,
			Error: models.ErrorResponse{
				Message:"Success",
				Code:200,
			},
		}
	}else{
		res = models.Response{
			Data: nil,
			Error: models.ErrorResponse{
				Message: err,
				Code: 403,
			},
		}
	}

	c.Data["json"] = res
	c.ServeJSON()
}

// Post ...
// @Title Send amount CNC from to
// @Description create request tranfer
// @Param	body		body 	models.TransferFromRequest	true		"body for Payment content"
// @Success 201 {object} models.Response
// @Failure 403 body is empty
// @router /transferfrom [post]
func (c *PaymentController) TransferFrom() {
	var rq models.TransferFromRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &rq)

	var res models.Response
	tx, err := payment.TransferFrom(rq)
	if err == nil{
		res = models.Response{
			Data: tx,
			Error: models.ErrorResponse{
				Message:"Success",
				Code:200,
			},
		}
	}else{
		res = models.Response{
			Data: nil,
			Error: models.ErrorResponse{
				Message: err,
				Code: 403,
			},
		}
	}

	c.Data["json"] = res
	c.ServeJSON()
}

// Post ...
// @Title Earn amount CNC to address
// @Description create request earn CNC
// @Param	body		body 	models.EarnRequest	true		"body for Earn"
// @Success 201 {object} models.Response
// @Failure 403 body is empty
// @router /earn [post]
func (c *PaymentController) Earn(){
	var eq models.EarnRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &eq)

	var res models.Response
	tx, err := payment.Earn(eq)
	if err == nil{
		res = models.Response{
			Data: tx,
			Error: models.ErrorResponse{
				Message:"Success",
				Code:200,
			},
		}
	}else{
		res = models.Response{
			Data: nil,
			Error: models.ErrorResponse{
				Message: err,
				Code: 403,
			},
		}
	}

	c.Data["json"] = res
	c.ServeJSON()
}

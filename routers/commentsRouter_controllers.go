package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["codedaotoken/controllers:ObjectController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:ObjectController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:ObjectController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:ObjectController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:ObjectController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:PaymentController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:PaymentController"],
		beego.ControllerComments{
			Method: "GetBalanceOf",
			Router: `/balanceof`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:PaymentController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:PaymentController"],
		beego.ControllerComments{
			Method: "Earn",
			Router: `/earn`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:PaymentController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:PaymentController"],
		beego.ControllerComments{
			Method: "GetOwner",
			Router: `/owner`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:PaymentController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:PaymentController"],
		beego.ControllerComments{
			Method: "GetTotalSupply",
			Router: `/totalsupply`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:PaymentController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:PaymentController"],
		beego.ControllerComments{
			Method: "Transfer",
			Router: `/transfer`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:PaymentController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:PaymentController"],
		beego.ControllerComments{
			Method: "TransferFrom",
			Router: `/transferfrom`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:UserController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:UserController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:UserController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:UserController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:UserController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:UserController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["codedaotoken/controllers:UserController"] = append(beego.GlobalControllerRouter["codedaotoken/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}

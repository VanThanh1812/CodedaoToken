// @APIVersion 1.0.0
// @Title API hệ thống quảng bá theo mạng lưới
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact phanvanthanh.mrt@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"codedaotoken/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/payment",
			beego.NSInclude(
				&controllers.PaymentController{},
			)),
	)
	beego.AddNamespace(ns)
}

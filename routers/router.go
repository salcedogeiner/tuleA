// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"tuleapApi/controllers"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.Debug("Filters init...")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
			AllowAllOrigins: true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
			AllowCredentials: true,
	}))
	
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/historia_usuario",
			beego.NSInclude(
				&controllers.HistoriaUsuarioController{},
			),
		),

		beego.NSNamespace("/tarea",
			beego.NSInclude(
				&controllers.TareaController{},
			),
		),

		beego.NSNamespace("/comentario",
			beego.NSInclude(
				&controllers.ComentarioController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["tuleapApi/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:HistoriaUsuarioController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:HistoriaUsuarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:HistoriaUsuarioController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:HistoriaUsuarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:HistoriaUsuarioController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:HistoriaUsuarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:HistoriaUsuarioController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:HistoriaUsuarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:HistoriaUsuarioController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:HistoriaUsuarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:TareaController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:TareaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:TareaController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:TareaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:TareaController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:TareaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:TareaController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:TareaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["tuleapApi/controllers:TareaController"] = append(beego.GlobalControllerRouter["tuleapApi/controllers:TareaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}

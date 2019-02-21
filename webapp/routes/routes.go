package routes

import (
	"github.com/kataras/iris/mvc"
	"gitlab.com/z547743799/irissso/bootstrap"
	"gitlab.com/z547743799/irissso/controller"
	"gitlab.com/z547743799/irissso/service"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {

	page := mvc.New(b.Party("/"))
	page.Handle(new(controller.Controller))

	LoginService := service.NewTbLoginService()
	Login := mvc.New(b.Party("/user"))
	Login.Register(LoginService)
	Login.Handle(new(controller.LoginController))

	RegisterService := service.NewTbRegisterService()
	Register := mvc.New(b.Party("/user"))
	Register.Register(RegisterService)
	Register.Handle(new(controller.RegisterController))

	TokenService := service.NewTokenService()
	Token := mvc.New(b.Party("/user"))
	Token.Register(TokenService)
	Token.Handle(new(controller.TokenController))

	//admin := mvc.New(b.Party("/admin"))
	//admin.Router.Use(middleware.BasicAuth)
	//admin.Register(superstarService)
	//admin.Handle(new(controllers.AdminController))

	//b.Get("/follower/{id:long}", GetFollowerHandler)
	//b.Get("/following/{id:long}", GetFollowingHandler)
	//b.Get("/like/{id:long}", GetLikeHandler)
}

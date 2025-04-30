package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/ykp_pos/controller"
	"github.com/jejevj/ykp_pos/middleware"
	"github.com/jejevj/ykp_pos/service"
)

func User(route fiber.Router, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/user")

	routes.Post("", userController.Register)
	routes.Get("", userController.GetAllUser)
	routes.Post("/login", userController.Login)
	routes.Delete("", middleware.Authenticate(jwtService), userController.Delete)
	routes.Put("", middleware.Authenticate(jwtService), userController.Update)
	routes.Get("/me", middleware.Authenticate(jwtService), userController.Me)
	routes.Post("/verify_email", userController.VerifyEmail)
	routes.Post("/send_verification_email", userController.SendVerificationEmail)
}

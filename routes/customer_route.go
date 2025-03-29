package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/ykp_pos/controller"
	"github.com/jejevj/ykp_pos/middleware"
	"github.com/jejevj/ykp_pos/service"
)

func Customer(route fiber.Router, customerController controller.CustomerController, jwtService service.JWTService) {
	routes := route.Group("/customer")

	routes.Post("", customerController.AddCustomer)
	routes.Get("", customerController.GetAllCustomerWithPagination)
	routes.Delete("", middleware.Authenticate(jwtService), customerController.DeleteCustomer)
	routes.Put("", middleware.Authenticate(jwtService), customerController.UpdateCustomer)
	routes.Get("/by-id", middleware.Authenticate(jwtService), customerController.GetCustomerById)
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/ykp_pos/controller"
	"github.com/jejevj/ykp_pos/middleware"
	"github.com/jejevj/ykp_pos/service"
)

func Loading(route fiber.Router, loadingController controller.LoadingController, jwtService service.JWTService) {
	routes := route.Group("/loading")

	routes.Post("", loadingController.AddLoading)
	routes.Get("", loadingController.GetAllLoadingWithPagination)
	routes.Delete("", middleware.Authenticate(jwtService), loadingController.DeleteLoading)
	routes.Put("", middleware.Authenticate(jwtService), loadingController.UpdateLoading)
	routes.Get("/by-id", middleware.Authenticate(jwtService), loadingController.GetLoadingById)
}

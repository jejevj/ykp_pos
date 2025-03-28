package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/ykp_pos/controller"
	"github.com/jejevj/ykp_pos/middleware"
	"github.com/jejevj/ykp_pos/service"
)

func Satuan(route fiber.Router, satuanController controller.SatuanController, jwtService service.JWTService) {
	routes := route.Group("/satuan")

	routes.Post("", satuanController.AddSatuan)
	routes.Get("", satuanController.GetAllSatuanWithPagination)
	routes.Delete("", middleware.Authenticate(jwtService), satuanController.DeleteSatuan)
	routes.Put("", middleware.Authenticate(jwtService), satuanController.UpdateSatuan)
	routes.Get("/by-id", middleware.Authenticate(jwtService), satuanController.GetSatuanById)
}

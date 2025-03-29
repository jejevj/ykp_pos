package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/ykp_pos/controller"
	"github.com/jejevj/ykp_pos/middleware"
	"github.com/jejevj/ykp_pos/service"
)

func Barang(route fiber.Router, barangController controller.BarangController, jwtService service.JWTService) {
	routes := route.Group("/barang")

	routes.Post("", barangController.AddBarang)
	routes.Get("", barangController.GetAllBarangWithPagination)
	routes.Delete("", middleware.Authenticate(jwtService), barangController.DeleteBarang)
	routes.Put("", middleware.Authenticate(jwtService), barangController.UpdateBarang)
	routes.Get("/by-id", middleware.Authenticate(jwtService), barangController.GetBarangById)
}

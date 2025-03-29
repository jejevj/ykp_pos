package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/ykp_pos/controller"
	"github.com/jejevj/ykp_pos/middleware"
	"github.com/jejevj/ykp_pos/service"
)

func Transaksi(route fiber.Router, transaksiController controller.TransaksiController, jwtService service.JWTService) {
	routes := route.Group("/trx")

	routes.Post("", transaksiController.AddTransaksi)
	routes.Get("", transaksiController.GetAllTransaksiWithPagination)
	routes.Delete("", middleware.Authenticate(jwtService), transaksiController.DeleteTransaksi)
	routes.Put("", middleware.Authenticate(jwtService), transaksiController.UpdateTransaksi)
	routes.Get("/by-id", middleware.Authenticate(jwtService), transaksiController.GetTransaksiById)
}

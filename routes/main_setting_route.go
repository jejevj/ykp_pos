package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/ykp_pos/controller"
	"github.com/jejevj/ykp_pos/middleware"
	"github.com/jejevj/ykp_pos/service"
)

func MainSetting(route fiber.Router, mainSettingController controller.MainSettingController, jwtService service.JWTService) {
	routes := route.Group("/main-setting")

	routes.Post("", mainSettingController.AddMainSetting)
	routes.Get("", mainSettingController.GetAllMainSettingWithPagination)
	routes.Delete("", middleware.Authenticate(jwtService), mainSettingController.DeleteMainSetting)
	routes.Put("", middleware.Authenticate(jwtService), mainSettingController.UpdateMainSetting)
	routes.Get("/by-id", middleware.Authenticate(jwtService), mainSettingController.GetMainSettingById)
}

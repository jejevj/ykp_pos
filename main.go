package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/ykp_pos/cmd"
	"github.com/jejevj/ykp_pos/config"
	"github.com/jejevj/ykp_pos/controller"
	"github.com/jejevj/ykp_pos/middleware"
	"github.com/jejevj/ykp_pos/repository"
	"github.com/jejevj/ykp_pos/routes"
	"github.com/jejevj/ykp_pos/service"
)

func main() {
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if len(os.Args) > 1 {
		cmd.Commands(db)
		return
	}

	var (
		jwtService service.JWTService = service.NewJWTService()

		// Repository
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		// Service
		userService service.UserService = service.NewUserService(userRepository, jwtService)
		// Controller
		userController controller.UserController = controller.NewUserController(userService)

		// Satuan Service
		// Repository
		satuanRepository repository.SatuanRepository = repository.NewSatuanRepository(db)
		// Service
		satuanService service.SatuanService = service.NewSatuanService(satuanRepository, jwtService)
		// Controller
		satuanController controller.SatuanController = controller.NewSatuanController(satuanService)

		// Barang Service
		// Repository
		barangRepository repository.BarangRepository = repository.NewBarangRepository(db)
		// Service
		barangService service.BarangService = service.NewBarangService(barangRepository, jwtService)
		// Controller
		barangController controller.BarangController = controller.NewBarangController(barangService)

		// Loading Service
		// Repository
		loadingRepository repository.LoadingRepository = repository.NewLoadingRepository(db)
		// Service
		loadingService service.LoadingService = service.NewLoadingService(loadingRepository, jwtService)
		// Controller
		loadingController controller.LoadingController = controller.NewLoadingController(loadingService)

		// Transaksi Service
		// Repository
		transaksiRepository repository.TransaksiRepository = repository.NewTransaksiRepository(db)
		// Service
		transaksiService service.TransaksiService = service.NewTransaksiService(transaksiRepository, jwtService)
		// Controller
		transaksiController controller.TransaksiController = controller.NewTransaksiController(transaksiService)

		// Customer Service
		// Repository
		customerRepository repository.CustomerRepository = repository.NewCustomerRepository(db)
		// Service
		customerService service.CustomerService = service.NewCustomerService(customerRepository, jwtService)
		// Controller
		customerController controller.CustomerController = controller.NewCustomerController(customerService)
	)

	server := fiber.New()
	server.Use(middleware.CORSMiddleware())
	apiGroup := server.Group("/api")

	// routes
	routes.User(apiGroup, userController, jwtService)
	routes.Satuan(apiGroup, satuanController, jwtService)
	routes.Barang(apiGroup, barangController, jwtService)
	routes.Loading(apiGroup, loadingController, jwtService)
	routes.Transaksi(apiGroup, transaksiController, jwtService)
	routes.Customer(apiGroup, customerController, jwtService)

	server.Static("/assets", "./assets")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Listen(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

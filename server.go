package main

import (
	"golang_api_hupiutang/config"
	"golang_api_hupiutang/controllers"
	"golang_api_hupiutang/middleware"
	"golang_api_hupiutang/repository"
	"golang_api_hupiutang/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                       *gorm.DB                             = config.SetupDatabaseConnection()
	userRepository           repository.UserRepository            = repository.NewUserRepository(db)
	hutangRepository         repository.HutangRepository          = repository.NewHutangRepository(db)
	piutangRepository        repository.PiutangRepository         = repository.NewPiutangRepository(db)
	lunasHutangRepository    repository.LunasHutangRepository     = repository.NewLunasHutangRepository(db)
	lunasPiutangRepository   repository.LunasPiutangRepository    = repository.NewLunasPiutangRepository(db)
	cicilanHutangRepository  repository.CicilanHutangRepository   = repository.NewCicilanHutangRepository(db)
	cicilanPiutangRepository repository.CicilanPiutangRepository  = repository.NewCicilanPiutangRepository(db)
	jwtService               service.JWTService                   = service.NewJWTService()
	userService              service.UserService                  = service.NewUserService(userRepository)
	hutangService            service.HutangService                = service.NewHutangService(hutangRepository)
	piutangService           service.PiutangService               = service.NewPiutangService(piutangRepository)
	lunasHutangService       service.LunasHutangService           = service.NewLunasHutangService(lunasHutangRepository)
	lunasPiutangService      service.LunasPiutangService          = service.NewLunasPiutangService(lunasPiutangRepository)
	cicilanHutangService     service.CicilanHutangService         = service.NewCicilanHutangService(cicilanHutangRepository)
	cicilanPiutangService    service.CicilanPiutangService        = service.NewCicilanPiutangService(cicilanPiutangRepository)
	authService              service.AuthService                  = service.NewAuthService(userRepository)
	authController           controllers.AuthController           = controllers.NewAuthController(authService, jwtService)
	userController           controllers.UserController           = controllers.NewUserController(userService, jwtService)
	hutangController         controllers.HutangController         = controllers.NewHutangController(hutangService, jwtService)
	piutangController        controllers.PiutangController        = controllers.NewPiutangController(piutangService, jwtService)
	lunashutangController    controllers.LunasHutangController    = controllers.NewLunasHutangController(lunasHutangService, jwtService)
	lunaspiutangController   controllers.LunasPiutangController   = controllers.NewLunasPiutangController(lunasPiutangService, jwtService)
	cicilanhutangController  controllers.CicilanHutangController  = controllers.NewCicilanHutangController(cicilanHutangService, jwtService)
	cicilanpiutangController controllers.CicilanPiutangController = controllers.NewCicilanPiutangController(cicilanPiutangService, jwtService)
)

func main() {

	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/register", authController.Register)
		}

		userRoutes := v1.Group("/user", middleware.AuthorizeJWT(jwtService))
		{
			userRoutes.GET("/profile", userController.Profile)
			userRoutes.PUT("/profile", userController.Update)
		}

		hutangRoutes := v1.Group("/hutangs", middleware.AuthorizeJWT(jwtService))
		{
			hutangRoutes.GET("/", hutangController.All)
			hutangRoutes.POST("/", hutangController.Insert)
			hutangRoutes.GET("/:id", hutangController.FindByID)
			hutangRoutes.PUT("/:id", hutangController.Update)
			hutangRoutes.DELETE("/:id", hutangController.Delete)
			hutangRoutes.PUT("/lunas/:id", lunashutangController.UpdateIsLunasHutang)
			hutangRoutes.PUT("/cicilan/:id", cicilanhutangController.UpdateCicilanHutang)
		}
		piutangRoutes := v1.Group("/piutangs", middleware.AuthorizeJWT(jwtService))
		{
			piutangRoutes.GET("/", piutangController.All)
			piutangRoutes.POST("/", piutangController.Insert)
			piutangRoutes.GET("/:id", piutangController.FindByID)
			piutangRoutes.PUT("/:id", piutangController.Update)
			piutangRoutes.DELETE("/:id", piutangController.Delete)
			piutangRoutes.PUT("/lunas/:id", lunaspiutangController.UpdateIsLunasPiutang)
			piutangRoutes.PUT("/cicilan/:id", cicilanpiutangController.UpdateCicilanPiutang)
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

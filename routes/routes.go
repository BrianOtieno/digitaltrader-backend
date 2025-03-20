package routes

import (
	"digitaltrader/controllers"
	"digitaltrader/middlewares"
	"time"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Middleware(cors.Config{
		// Origins: "http://localhost",
		Origins:         "https://xxx.com",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		Credentials:     true, // Allow credentials
		MaxAge:          50 * time.Second,
		ValidateHeaders: false,
	}))

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"https://xxx.com"},
	// 	// AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	// 	ExposeHeaders:    []string{},
	// 	AllowCredentials: true,
	// 	MaxAge:           50 * time.Second,
	// }))

	// Apply the recovery middleware
	r.Use(middlewares.RecoveryMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Quantum Cartel Home")
	})

	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			//Register
			public.POST("/signup", controllers.CreateUser)
			// Sign in
			public.POST("/login", controllers.Login)
			//refresh token
			public.POST("/refresh-token", controllers.RefreshToken)
		}

		protected := api.Group("/auth").Use(middlewares.Authz())
		{
			// user
			protected.POST("/users", controllers.CreateUser)
			protected.GET("/users", controllers.GetAllUsers)
			protected.GET("/users/:id", controllers.GetUserByID)
			protected.PUT("/users/:id", controllers.UpdateUser)
			protected.DELETE("/users/:id", controllers.DeleteUser)

			// password reset
			protected.POST("/", controllers.CreatePasswordResetTokenController)
			protected.GET("/:email", controllers.GetPasswordResetTokenController)
			protected.PUT("/:email", controllers.UpdatePasswordResetTokenController)
			protected.DELETE("/:email", controllers.DeletePasswordResetTokenController)
			protected.GET("/", controllers.GetAllPasswordResetTokensController)

			// transactions
			protected.POST("/", controllers.CreateTransactionController)
			protected.GET("/:id", controllers.GetTransactionController)
			protected.PUT("/:id", controllers.UpdateTransactionController)
			protected.DELETE("/:id", controllers.DeleteTransactionController)
			protected.GET("/", controllers.GetAllTransactionsController)

			// transfers
			protected.POST("/", controllers.CreateTransferHandler)
			protected.GET("/:id", controllers.GetTransferByIDHandler)
			protected.GET("/", controllers.GetAllTransfersHandler)
			protected.PUT("/:id", controllers.UpdateTransferHandler)
			protected.DELETE("/:id", controllers.DeleteTransferHandler)

			//wallets
			protected.POST("", controllers.CreateWalletController)
			protected.GET("/:id", controllers.GetWalletController)
			protected.PUT("/:id", controllers.UpdateWalletController)
			protected.DELETE("/:id", controllers.DeleteWalletController)
			protected.GET("", controllers.GetAllWalletsController)

			// failed jobs
			protected.POST("", controllers.CreateFailedJobController)
			protected.GET("/:id", controllers.GetFailedJobController)
			protected.PUT("/:id", controllers.UpdateFailedJobController)
			protected.DELETE("/:id", controllers.DeleteFailedJobController)
			protected.GET("", controllers.GetAllFailedJobsController)

			// categories
			protected.POST("", controllers.CreateCategoryController)
			protected.GET("/:id", controllers.GetCategoryController)
			protected.PUT("/:id", controllers.UpdateCategoryController)
			protected.DELETE("/:id", controllers.DeleteCategoryController)
			protected.GET("", controllers.GetAllCategoriesController)

			// address
			protected.POST("", controllers.CreateAddressController)
			protected.GET("/:id", controllers.GetAddressController)
			protected.PUT("/:id", controllers.UpdateAddressController)
			protected.DELETE("/:id", controllers.DeleteAddressController)
			protected.GET("", controllers.GetAllAddressesController)

			// chat room
			protected.POST("", controllers.CreateChatRoomController)
			protected.GET("/:id", controllers.GetChatRoomController)
			protected.PUT("/:id", controllers.UpdateChatRoomController)
			protected.DELETE("/:id", controllers.DeleteChatRoomController)
			protected.GET("", controllers.GetAllChatRoomsController)

			// cities
			protected.POST("", controllers.CreateCityController)
			protected.GET("/:id", controllers.GetCityController)
			protected.PUT("/:id", controllers.UpdateCityController)
			protected.DELETE("/:id", controllers.DeleteCityController)
			protected.GET("", controllers.GetAllCitiesController)

			// favourites
			protected.POST("", controllers.CreateFavouriteController)
			protected.GET("/:id", controllers.GetFavouriteController)
			protected.PUT("/:id", controllers.UpdateFavouriteController)
			protected.DELETE("/:id", controllers.DeleteFavouriteController)
			protected.GET("", controllers.GetAllFavouritesController)

			// flush
			protected.POST("", controllers.CreateFlushController)
			protected.GET("/:id", controllers.GetFlushController)
			protected.PUT("/:id", controllers.UpdateFlushController)
			protected.DELETE("/:id", controllers.DeleteFlushController)
			protected.GET("", controllers.GetAllFlushesController)

			// chat message
			protected.POST("", controllers.CreateChatMessageController)
			protected.GET("/:id", controllers.GetChatMessageController)
			protected.PUT("/:id", controllers.UpdateChatMessageController)
			protected.DELETE("/:id", controllers.DeleteChatMessageController)
			protected.GET("", controllers.GetAllChatMessagesController)

			// general
			protected.POST("", controllers.CreateGeneralController)
			protected.GET("/:id", controllers.GetGeneralController)
			protected.PUT("/:id", controllers.UpdateGeneralController)
			protected.DELETE("/:id", controllers.DeleteGeneralController)
			protected.GET("", controllers.GetAllGeneralsController)

			// language
			protected.POST("", controllers.CreateLanguageController)
			protected.GET("/:id", controllers.GetLanguageController)
			protected.PUT("/:id", controllers.UpdateLanguageController)
			protected.DELETE("/:id", controllers.DeleteLanguageController)
			protected.GET("", controllers.GetAllLanguagesController)

			// otp
			protected.POST("", controllers.CreateOTPController)
			protected.GET("/:id", controllers.GetOTPController)
			protected.PUT("/:id", controllers.UpdateOTPController)
			protected.DELETE("/:id", controllers.DeleteOTPController)
			protected.GET("", controllers.GetAllOTPsController)

			// payment
			protected.POST("", controllers.CreatePaymentController)
			protected.GET("/:id", controllers.GetPaymentController)
			protected.PUT("/:id", controllers.UpdatePaymentController)
			protected.DELETE("/:id", controllers.DeletePaymentController)
			protected.GET("", controllers.GetAllPaymentsController)

		}
	}

	return r
}

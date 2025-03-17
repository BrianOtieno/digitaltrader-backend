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
		Origins:         "https://quantum-cartel.com",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		Credentials:     true, // Allow credentials
		MaxAge:          50 * time.Second,
		ValidateHeaders: false,
	}))

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"https://quantum-cartel.com"},
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

		}
	}

	return r
}

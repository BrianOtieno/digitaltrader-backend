package routes

import (
	"digitaltrader/controllers"
	"digitaltrader/middlewares"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

// SetupRouter initializes and configures the Gin router with all routes
func SetupRouter() *gin.Engine {
	// Initialize Gin with default middleware (logger and recovery)
	r := gin.Default()

	// Apply CORS middleware
	r.Use(cors.Middleware(cors.Config{
		Origins:         "https://xxx.com",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		Credentials:     true,
		MaxAge:          50 * time.Second,
		ValidateHeaders: false,
	}))

	// Apply custom recovery middleware
	r.Use(middlewares.RecoveryMiddleware())

	// Root and health check endpoints
	r.GET("/ping/", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome Home")
	})

	// API base group
	api := r.Group("/api/v1/")
	{
		// Public routes (no authentication required)
		public := api.Group("/public/")
		{
			public.POST("/signup/", controllers.CreateUser)
			public.POST("/login/", controllers.Login)
			public.POST("/refresh-token/", controllers.RefreshToken)
		}

		// Protected routes (authentication required)
		protected := api.Group("/auth/")
		protected.Use(middlewares.Authz())
		{
			// Users
			users := protected.Group("/users/")
			{
				users.POST("", controllers.CreateUser)
				users.GET("", controllers.GetAllUsers)
				users.GET(":id/", controllers.GetUserByID)
				users.PUT(":id/", controllers.UpdateUser)
				users.DELETE(":id/", controllers.DeleteUser)
			}

			// Password Reset Tokens
			passwordReset := protected.Group("/password-reset/")
			{
				passwordReset.POST("", controllers.CreatePasswordResetTokenController)
				passwordReset.GET(":email/", controllers.GetPasswordResetTokenController)
				passwordReset.PUT(":email/", controllers.UpdatePasswordResetTokenController)
				passwordReset.DELETE(":email/", controllers.DeletePasswordResetTokenController)
				passwordReset.GET("", controllers.GetAllPasswordResetTokensController)
			}

			// Transactions
			transactions := protected.Group("/transactions/")
			{
				transactions.POST("", controllers.CreateTransactionController)
				transactions.GET(":id/", controllers.GetTransactionController)
				transactions.PUT(":id/", controllers.UpdateTransactionController)
				transactions.DELETE(":id/", controllers.DeleteTransactionController)
				transactions.GET("", controllers.GetAllTransactionsController)
			}

			// Transfers
			transfers := protected.Group("/transfers/")
			{
				transfers.POST("", controllers.CreateTransferHandler)
				transfers.GET(":id/", controllers.GetTransferByIDHandler)
				transfers.GET("", controllers.GetAllTransfersHandler)
				transfers.PUT(":id/", controllers.UpdateTransferHandler)
				transfers.DELETE(":id/", controllers.DeleteTransferHandler)
			}

			// Wallets
			wallets := protected.Group("/wallets/")
			{
				wallets.POST("", controllers.CreateWalletController)
				wallets.GET(":id/", controllers.GetWalletController)
				wallets.PUT(":id/", controllers.UpdateWalletController)
				wallets.DELETE(":id/", controllers.DeleteWalletController)
				wallets.GET("", controllers.GetAllWalletsController)
			}

			// Failed Jobs
			failedJobs := protected.Group("/failed-jobs/")
			{
				failedJobs.POST("", controllers.CreateFailedJobController)
				failedJobs.GET(":id/", controllers.GetFailedJobController)
				failedJobs.PUT(":id/", controllers.UpdateFailedJobController)
				failedJobs.DELETE(":id/", controllers.DeleteFailedJobController)
				failedJobs.GET("", controllers.GetAllFailedJobsController)
			}

			// Categories
			categories := protected.Group("/categories/")
			{
				categories.POST("", controllers.CreateCategoryController)
				categories.GET(":id/", controllers.GetCategoryController)
				categories.PUT(":id/", controllers.UpdateCategoryController)
				categories.DELETE(":id/", controllers.DeleteCategoryController)
				categories.GET("", controllers.GetAllCategoriesController)
			}

			// Addresses
			addresses := protected.Group("/addresses/")
			{
				addresses.POST("", controllers.CreateAddressController)
				addresses.GET(":id/", controllers.GetAddressController)
				addresses.PUT(":id/", controllers.UpdateAddressController)
				addresses.DELETE(":id/", controllers.DeleteAddressController)
				addresses.GET("", controllers.GetAllAddressesController)
			}

			// Chat Rooms
			chatRooms := protected.Group("/chat-rooms/")
			{
				chatRooms.POST("", controllers.CreateChatRoomController)
				chatRooms.GET(":id/", controllers.GetChatRoomController)
				chatRooms.PUT(":id/", controllers.UpdateChatRoomController)
				chatRooms.DELETE(":id/", controllers.DeleteChatRoomController)
				chatRooms.GET("", controllers.GetAllChatRoomsController)
			}

			// Cities
			cities := protected.Group("/cities/")
			{
				cities.POST("", controllers.CreateCityController)
				cities.GET(":id/", controllers.GetCityController)
				cities.PUT(":id/", controllers.UpdateCityController)
				cities.DELETE(":id/", controllers.DeleteCityController)
				cities.GET("", controllers.GetAllCitiesController)
			}

			// Favourites
			favourites := protected.Group("/favourites/")
			{
				favourites.POST("", controllers.CreateFavouriteController)
				favourites.GET(":id/", controllers.GetFavouriteController)
				favourites.PUT(":id/", controllers.UpdateFavouriteController)
				favourites.DELETE(":id/", controllers.DeleteFavouriteController)
				favourites.GET("", controllers.GetAllFavouritesController)
			}

			// Flush
			flushes := protected.Group("/flushes/")
			{
				flushes.POST("", controllers.CreateFlushController)
				flushes.GET(":id/", controllers.GetFlushController)
				flushes.PUT(":id/", controllers.UpdateFlushController)
				flushes.DELETE(":id/", controllers.DeleteFlushController)
				flushes.GET("", controllers.GetAllFlushesController)
			}

			// Chat Messages
			chatMessages := protected.Group("/chat-messages/")
			{
				chatMessages.POST("", controllers.CreateChatMessageController)
				chatMessages.GET(":id/", controllers.GetChatMessageController)
				chatMessages.PUT(":id/", controllers.UpdateChatMessageController)
				chatMessages.DELETE(":id/", controllers.DeleteChatMessageController)
				chatMessages.GET("", controllers.GetAllChatMessagesController)
			}

			// General
			generals := protected.Group("/generals/")
			{
				generals.POST("", controllers.CreateGeneralController)
				generals.GET(":id/", controllers.GetGeneralController)
				generals.PUT(":id/", controllers.UpdateGeneralController)
				generals.DELETE(":id/", controllers.DeleteGeneralController)
				generals.GET("", controllers.GetAllGeneralsController)
			}

			// Languages
			languages := protected.Group("/languages/")
			{
				languages.POST("", controllers.CreateLanguageController)
				languages.GET(":id/", controllers.GetLanguageController)
				languages.PUT(":id/", controllers.UpdateLanguageController)
				languages.DELETE(":id/", controllers.DeleteLanguageController)
				languages.GET("", controllers.GetAllLanguagesController)
			}

			// OTPs
			otps := protected.Group("/otps/")
			{
				otps.POST("", controllers.CreateOTPController)
				otps.GET(":id/", controllers.GetOTPController)
				otps.PUT(":id/", controllers.UpdateOTPController)
				otps.DELETE(":id/", controllers.DeleteOTPController)
				otps.GET("", controllers.GetAllOTPsController)
			}

			// Payments
			payments := protected.Group("/payments/")
			{
				payments.POST("", controllers.CreatePaymentController)
				payments.GET(":id/", controllers.GetPaymentController)
				payments.PUT(":id/", controllers.UpdatePaymentController)
				payments.DELETE(":id/", controllers.DeletePaymentController)
				payments.GET("", controllers.GetAllPaymentsController)
			}

			// Popups
			popups := protected.Group("/popups/")
			{
				popups.POST("", controllers.CreatePopupController)
				popups.GET(":id/", controllers.GetPopupController)
				popups.PUT(":id/", controllers.UpdatePopupController)
				popups.DELETE(":id/", controllers.DeletePopupController)
				popups.GET("", controllers.GetAllPopupsController)
			}

			// Stores
			stores := protected.Group("/stores/")
			{
				stores.POST("", controllers.CreateStoreController)
				stores.GET(":id/", controllers.GetStoreController)
				stores.PUT(":id/", controllers.UpdateStoreController)
				stores.DELETE(":id/", controllers.DeleteStoreController)
				stores.GET("", controllers.GetAllStoresController)
			}

			// Sub-Categories
			subCategories := protected.Group("/sub-categories/")
			{
				subCategories.POST("", controllers.CreateSubCategoryController)
				subCategories.GET(":id/", controllers.GetSubCategoryController)
				subCategories.PUT(":id/", controllers.UpdateSubCategoryController)
				subCategories.DELETE(":id/", controllers.DeleteSubCategoryController)
				subCategories.GET("", controllers.GetAllSubCategoriesController)
			}

			// Analytics
			analytics := protected.Group("/analytics/")
			{
				analytics.POST("", controllers.CreateAnalyticsController)
				analytics.GET(":id/", controllers.GetAnalyticsController)
				analytics.PUT(":id/", controllers.UpdateAnalyticsController)
				analytics.DELETE(":id/", controllers.DeleteAnalyticsController)
				analytics.GET("", controllers.GetAllAnalyticsController)
			}

			// Contacts
			contacts := protected.Group("/contacts/")
			{
				contacts.POST("", controllers.CreateContactController)
				contacts.GET(":id/", controllers.GetContactController)
				contacts.PUT(":id/", controllers.UpdateContactController)
				contacts.DELETE(":id/", controllers.DeleteContactController)
				contacts.GET("", controllers.GetAllContactsController)
			}

			// Drivers
			drivers := protected.Group("/drivers/")
			{
				drivers.POST("", controllers.CreateDriverController)
				drivers.GET(":id/", controllers.GetDriverController)
				drivers.PUT(":id/", controllers.UpdateDriverController)
				drivers.DELETE(":id/", controllers.DeleteDriverController)
				drivers.GET("", controllers.GetAllDriversController)
			}

			// Manage
			manages := protected.Group("/manages/")
			{
				manages.POST("", controllers.CreateManageController)
				manages.GET(":id/", controllers.GetManageController)
				manages.PUT(":id/", controllers.UpdateManageController)
				manages.DELETE(":id/", controllers.DeleteManageController)
				manages.GET("", controllers.GetAllManagesController)
			}

			// Orders
			orders := protected.Group("/orders/")
			{
				orders.POST("", controllers.CreateOrderController)
				orders.GET(":id/", controllers.GetOrderController)
				orders.PUT(":id/", controllers.UpdateOrderController)
				orders.DELETE(":id/", controllers.DeleteOrderController)
				orders.GET("", controllers.GetAllOrdersController)
			}

			// Products
			products := protected.Group("/products/")
			{
				products.POST("", controllers.CreateProductController)
				products.GET(":id/", controllers.GetProductController)
				products.PUT(":id/", controllers.UpdateProductController)
				products.DELETE(":id/", controllers.DeleteProductController)
				products.GET("", controllers.GetAllProductsController)
			}

			// Redeems
			redeems := protected.Group("/redeems/")
			{
				redeems.POST("", controllers.CreateRedeemController)
				redeems.GET(":id/", controllers.GetRedeemController)
				redeems.PUT(":id/", controllers.UpdateRedeemController)
				redeems.DELETE(":id/", controllers.DeleteRedeemController)
				redeems.GET("", controllers.GetAllRedeemsController)
			}

			// Referrals
			referrals := protected.Group("/referrals/")
			{
				referrals.POST("", controllers.CreateReferralController)
				referrals.GET(":id/", controllers.GetReferralController)
				referrals.PUT(":id/", controllers.UpdateReferralController)
				referrals.DELETE(":id/", controllers.DeleteReferralController)
				referrals.GET("", controllers.GetAllReferralsController)
			}

			// ReferralCodes
			referralCodes := protected.Group("/referralcodes/")
			{
				referralCodes.POST("", controllers.CreateReferralCodeController)
				referralCodes.GET(":id/", controllers.GetReferralCodeController)
				referralCodes.PUT(":id/", controllers.UpdateReferralCodeController)
				referralCodes.DELETE(":id/", controllers.DeleteReferralCodeController)
				referralCodes.GET("", controllers.GetAllReferralCodesController)
			}

			// Store Requests
			storeRequests := protected.Group("/store-requests/")
			{
				storeRequests.POST("", controllers.CreateStoreRequestController)
				storeRequests.GET(":id/", controllers.GetStoreRequestController)
				storeRequests.PUT(":id/", controllers.UpdateStoreRequestController)
				storeRequests.DELETE(":id/", controllers.DeleteStoreRequestController)
				storeRequests.GET("", controllers.GetAllStoreRequestsController)
			}

			// Subscribers
			subscribers := protected.Group("/subscribers/")
			{
				subscribers.POST("", controllers.CreateSubscriberController)
				subscribers.GET(":id/", controllers.GetSubscriberController)
				subscribers.PUT(":id/", controllers.UpdateSubscriberController)
				subscribers.DELETE(":id/", controllers.DeleteSubscriberController)
				subscribers.GET("", controllers.GetAllSubscribersController)
			}

			// Complaints
			complaints := protected.Group("/complaints/")
			{
				complaints.POST("", controllers.CreateComplaintController)
				complaints.GET(":id/", controllers.GetComplaintController)
				complaints.PUT(":id/", controllers.UpdateComplaintController)
				complaints.DELETE(":id/", controllers.DeleteComplaintController)
				complaints.GET("", controllers.GetAllComplaintsController)
			}

			// Driver Requests
			driverRequests := protected.Group("/driver-requests/")
			{
				driverRequests.POST("", controllers.CreateDriverRequestController)
				driverRequests.GET(":id/", controllers.GetDriverRequestController)
				driverRequests.PUT(":id/", controllers.UpdateDriverRequestController)
				driverRequests.DELETE(":id/", controllers.DeleteDriverRequestController)
				driverRequests.GET("", controllers.GetAllDriverRequestsController)
			}

			// Settings
			settings := protected.Group("/settings/")
			{
				settings.POST("", controllers.CreateSettingController)
				settings.GET("", controllers.GetSettingController)
				settings.PUT("", controllers.UpdateSettingController)
				settings.DELETE("", controllers.DeleteSettingController)
				settings.GET("/all", controllers.GetAllSettingsController) // Optional, for consistency with other endpoints
			}
		}
	}

	return r
}

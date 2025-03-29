package migrations

import (
	"digitaltrader/db"
	"digitaltrader/models"
	"fmt"
)

func MigrateDB() {
	db := db.InitDB()
	fmt.Println("============")
	fmt.Printf("Firing models into %v Database", db.Migrator().CurrentDatabase())
	fmt.Println("")

	// Core models (no or minimal dependencies)
	models.MigrateUser()     // Users (needed by many others like Redeem, Complaint)
	models.MigrateCity()     // Cities (needed by StoreRequest, DriverRequest)
	models.MigrateCategory() // Categories (standalone or potentially referenced)
	models.MigrateGeneral()  // General settings (likely standalone)
	models.MigrateLanguage() // Languages (standalone)

	// Models with potential dependencies
	models.MigrateActivityLog()        // Activity logs (might reference User)
	models.MigrateAddress()            // Addresses (might reference User or City)
	models.MigrateChatRoom()           // Chat rooms (might reference User)
	models.MigrateChatMessage()        // Chat messages (depends on ChatRoom, User)
	models.MigrateFailedJob()          // Failed jobs (standalone or system-related)
	models.MigrateFavourite()          // Favourites (might reference User, Product)
	models.MigrateFlush()              // Flush (unclear purpose, assumed standalone)
	models.MigrateOTP()                // OTP (might reference User)
	models.MigratePasswordResetToken() // Password reset tokens (depends on User)
	models.MigrateWallet()             // Wallets (depends on User)

	// Payment-related models
	models.MigratePayment()     // Payments (standalone, referenced by Transaction)
	models.MigrateTransaction() // Transactions (might reference Payment, User)
	models.MigrateTransfer()    // Transfers (might reference Wallet, User)

	// Models added from recent migrations
	models.MigrateRedeem()        // Redeem (depends on User for OwnerID, RedeemerID)
	models.MigrateReferral()      // Referral (standalone)
	models.MigrateReferralCode()  // ReferralCode (depends on User for UID)
	models.MigrateStoreRequest()  // StoreRequest (depends on City for CID)
	models.MigrateSubscriber()    // Subscriber (standalone)
	models.MigrateComplaint()     // Complaint (depends on User, Order, Driver, Store, Product, Reason)
	models.MigrateDriverRequest() // DriverRequest (depends on City for CityID)
	models.MigrateSetting()       // Setting (standalone, singleton)
	models.MigratePage()          // Page (standalone)

	fmt.Println("============")
	fmt.Printf("Models Fired into %v Database Successfully!", db.Migrator().CurrentDatabase())
	fmt.Println("")
	fmt.Println("============")
}

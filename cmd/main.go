package main

import (
	"database/sql"
	"kreditplus/db"
	"kreditplus/internal/app/customers"
	"kreditplus/internal/app/transactions"
	"kreditplus/internal/app/users"
	"kreditplus/internal/repositories"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = zerolog.New(os.Stdout).With().Caller().Logger()
}

func main() {

	app := gin.Default()

	db := db.NewDbMysql()

	defer db.Close()

	setupContainer(app, db)

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to KreditPlus"})
	})
	app.Use(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "Not Found"})
	})

	log.Fatal().Msg(app.Run(":" + os.Getenv("PORT")).Error())

}

func setupContainer(app *gin.Engine, db *sql.DB) {

	// repo
	querier := repositories.New(db)

	// service
	userService := users.NewUserService(querier)
	customersService := customers.NewCustomerService(querier)
	transactionService := transactions.NewTransactionService(querier)

	// handler
	userHandler := users.NewUserHandler(userService)
	customersHandler := customers.NewCustomerHandler(customersService)
	transactionHandler := transactions.NewTransactionHandler(transactionService)

	// route
	users.RegisterUserRoute(app, userHandler)
	customers.RegisterCustomerRoute(app, customersHandler)
	transactions.RegisterTransactionRoute(app, transactionHandler)

}

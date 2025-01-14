package main

import (
	"context"
	"database/sql"
	"kreditplus/db"
	"kreditplus/internal/app/customers"
	"kreditplus/internal/app/transactions"
	"kreditplus/internal/app/users"
	"kreditplus/internal/repositories"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: app,
	}

	go func() {
		log.Info().Msgf("Starting server... on port %s", os.Getenv("PORT"))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server failed")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exiting")

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

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to KreditPlus"})
	})
	app.Use(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "Not Found"})
	})

}

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/AkifhanIlgaz/hotel-booking-app/config"
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/handlers"
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/middlewares"
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/routes"
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/services"
	"github.com/AkifhanIlgaz/hotel-booking-app/migrations"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/db"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/mail"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	mod := flag.String("mod", "dev", "Mode to run the app in: dev or prod")
	flag.Parse()

	fmt.Println("Running in", *mod, "mode")

	cfg, err := config.Load(*mod)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := db.Connect(cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer db.Close()

	err = migrations.Init(db)
	if err != nil {
		log.Fatalf("failed to seed databases: %v", err)
	}

	tokenManager, err := token.NewTokenManager(db, &cfg.Token)
	if err != nil {
		log.Fatalf("failed to create token manager: %v", err)
	}

	mailManager := mail.NewManager(cfg.SMTP)

	server := gin.Default()

	router := server.Group("/api")

	userService := services.NewUserService(db)
	otpService := services.NewOTPService(db)
	hotelService := services.NewHotelService(db)

	authHandler := handlers.NewAuthHandler(userService, otpService, tokenManager, mailManager)
	hotelHandler := handlers.NewHotelHandler(hotelService)
	authMiddleware := middlewares.NewAuthMiddleware(tokenManager)

	routeManager := routes.NewManager(router, authHandler, hotelHandler, authMiddleware)

	routeManager.SetupRoutes()

	port := ":8080"
	err = server.Run(port)
	if err != nil {
		panic("Gin sunucusu başlatılamadı: " + err.Error())
	}

}

// ! For development
func setCors(server *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))
}

package routes

import (
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/handlers"
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type Manager struct {
	r              *gin.RouterGroup
	authMiddleware *middlewares.AuthMiddleware
	authHandler    *handlers.AuthHandler
	hotelHandler   *handlers.HotelHandler
}

func NewManager(r *gin.RouterGroup, authHandler *handlers.AuthHandler, hotelHandler *handlers.HotelHandler, authMiddleware *middlewares.AuthMiddleware) *Manager {
	return &Manager{
		r:              r,
		authMiddleware: authMiddleware,
		authHandler:    authHandler,
		hotelHandler:   hotelHandler,
	}
}

func (m *Manager) SetupRoutes() {
	m.authRoutes()
	m.hotelRoutes()
}

func (m Manager) authRoutes() {
	auth := m.r.Group("/auth")
	{
		auth.POST("/login", m.authHandler.Login)
		auth.POST("/register", m.authHandler.Register)
		auth.POST("/logout", m.authHandler.Logout)
		auth.POST("/refresh", m.authHandler.Refresh)

		auth.POST("/change-password", m.authHandler.ChangePassword)
		auth.POST("/forgot-password", m.authHandler.ForgotPassword)
		auth.POST("/verify-otp", m.authHandler.VerifyOTP)

		auth.GET("/test", m.authMiddleware.AccessToken())
	}
}

func (m Manager) hotelRoutes() {
	hotel := m.r.Group("/hotels")

	{
		hotel.GET("/", m.hotelHandler.Hotels)
		hotel.GET("/:id", m.hotelHandler.Hotel)
	}
}

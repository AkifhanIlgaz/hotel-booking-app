package routes

import (
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/handlers"
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type Manager struct {
<<<<<<< HEAD
	r              *gin.RouterGroup
	authMiddleware *middlewares.AuthMiddleware
	authHandler    *handlers.AuthHandler
}

func NewManager(r *gin.RouterGroup, authHandler *handlers.AuthHandler, authMiddleware *middlewares.AuthMiddleware) *Manager {
	return &Manager{
		r:              r,
		authMiddleware: authMiddleware,
		authHandler:    authHandler,
=======
	r           *gin.RouterGroup
	authHandler *handlers.AuthHandler
}

func NewManager(r *gin.RouterGroup, authHandler *handlers.AuthHandler) *Manager {
	return &Manager{
		r:           r,
		authHandler: authHandler,
>>>>>>> cc2aac0c3ed8781304723a6f811ddfba44838dad
	}
}

func (m *Manager) SetupRoutes() {
	m.userRoutes()
}

func (m Manager) userRoutes() {
	auth := m.r.Group("/auth")
	{
		auth.POST("/login", m.authHandler.Login)
		auth.POST("/register", m.authHandler.Register)
		auth.POST("/logout", m.authHandler.Logout)
		auth.POST("/refresh", m.authHandler.Refresh)
<<<<<<< HEAD

		auth.POST("/change-password", m.authHandler.ChangePassword)
		auth.POST("/forgot-password", m.authHandler.ForgotPassword)
		auth.POST("/verify-otp", m.authHandler.VerifyOTP)

		auth.GET("/test", m.authMiddleware.AccessToken())
=======

		auth.POST("/change-password", m.authHandler.ChangePassword)
		auth.POST("/forgot-password", m.authHandler.ForgotPassword)
		auth.POST("/verify-otp", m.authHandler.VerifyOTP)
>>>>>>> cc2aac0c3ed8781304723a6f811ddfba44838dad
	}
}

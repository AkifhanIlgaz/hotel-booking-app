package routes

import (
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Manager struct {
	r           *gin.RouterGroup
	authHandler *handlers.AuthHandler
}

func NewManager(r *gin.RouterGroup, authHandler *handlers.AuthHandler) *Manager {
	return &Manager{
		r:           r,
		authHandler: authHandler,
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

		auth.POST("/change-password", m.authHandler.ChangePassword)
		auth.POST("/forgot-password", m.authHandler.ForgotPassword)
		auth.POST("/verify-otp", m.authHandler.VerifyOTP)
	}
}

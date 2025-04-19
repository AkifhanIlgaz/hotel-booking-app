package routes

import (
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Manager struct {
	r           *gin.RouterGroup
	userHandler *handlers.UserHandler
}

func NewManager(r *gin.RouterGroup, userHandler *handlers.UserHandler) *Manager {
	return &Manager{
		r:           r,
		userHandler: userHandler,
	}
}

func (m *Manager) SetupRoutes() {
	m.userRoutes()
}

func (m Manager) userRoutes() {
	auth := m.r.Group("/auth")
	{
		auth.POST("/login", m.userHandler.Login)
		auth.POST("/register", m.userHandler.Register)

		auth.POST("/refresh", m.userHandler.Refresh)
		// TODO: Implement

		auth.POST("/forgot-password")
		auth.POST("/logout")

	}
}

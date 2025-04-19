package handlers

import (
	"net/http"

	"github.com/AkifhanIlgaz/hotel-booking-app/internal/models"
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/services"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/errors"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/messages"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/response"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	userService  *services.UserService
	tokenManager *token.Manager
}

func NewAuthHandler(userService *services.UserService, tokenManager *token.Manager) *AuthHandler {
	return &AuthHandler{
		userService:  userService,
		tokenManager: tokenManager,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req models.RegistrationRequest
	if err := ctx.BindJSON(&req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			if len(validationErrors) > 0 {
				fe := validationErrors[0]
				var params []string
				switch fe.Tag() {
				case "min", "max":
					params = []string{fe.Param()}
				}
				msg := messages.ErrorMessage{
					Message: messages.MessageForTag(fe.Tag(), params...),
				}
				response.WithError(ctx, http.StatusBadRequest, msg.Message, fe)
				return
			}
		}
		response.WithError(ctx, http.StatusBadRequest, messages.InvalidJSONOrMissingFields, err)
		return
	}

	id, err := h.userService.RegisterUser(req)
	if err != nil {
		if errors.Is(err, errors.ErrEmailTaken) {
			response.WithError(ctx, http.StatusConflict, messages.EmailAlreadyRegistered, err)
			return
		}

		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	// Todo: role enum
	accessToken, err := h.tokenManager.GenerateAccessToken(id.String(), "user")
	if err != nil {
		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	refreshToken, err := h.tokenManager.GenerateRefreshToken(id)
	if err != nil {
		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	response.WithSuccess(ctx, http.StatusCreated, messages.SuccessfullyRegistered, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.BindJSON(&req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			if len(validationErrors) > 0 {
				fe := validationErrors[0]
				var params []string
				switch fe.Tag() {
				case "min", "max":
					params = []string{fe.Param()}
				}
				msg := messages.ErrorMessage{
					Message: messages.MessageForTag(fe.Tag(), params...),
				}
				response.WithError(ctx, http.StatusBadRequest, msg.Message, fe)
				return
			}
		}
		response.WithError(ctx, http.StatusBadRequest, messages.InvalidJSONOrMissingFields, err)
		return
	}

	user, err := h.userService.AuthenticateUser(req)
	if err != nil {
		if errors.Is(err, errors.ErrUserNotFound) {
			response.WithError(ctx, http.StatusNotFound, messages.UserNotFound, err)
			return
		}
		if errors.Is(err, errors.ErrWrongPassword) {
			response.WithError(ctx, http.StatusUnauthorized, messages.WrongPassword, err)
			return
		}
	}

	accessToken, err := h.tokenManager.GenerateAccessToken(user.Id.String(), "user")
	if err != nil {
		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	refreshToken, err := h.tokenManager.GenerateRefreshToken(user.Id)
	if err != nil {
		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	response.WithSuccess(ctx, http.StatusOK, messages.SuccessfullyLoggedIn, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}

func (h *AuthHandler) Refresh(ctx *gin.Context) {
	var req models.RefreshRequest
	if err := ctx.BindJSON(&req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			if len(validationErrors) > 0 {
				fe := validationErrors[0]
				var params []string
				switch fe.Tag() {
				case "min", "max":
					params = []string{fe.Param()}
				}
				msg := messages.ErrorMessage{
					Message: messages.MessageForTag(fe.Tag(), params...),
				}
				response.WithError(ctx, http.StatusBadRequest, msg.Message, fe)
				return
			}
		}
		response.WithError(ctx, http.StatusBadRequest, messages.InvalidJSONOrMissingFields, err)
		return
	}

	uid, err := h.tokenManager.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		if errors.Is(err, errors.ErrTokenExpired) {
			response.WithError(ctx, http.StatusUnauthorized, messages.TokenExpired, err)
			return
		}
	}
	accessToken, err := h.tokenManager.GenerateAccessToken(uid.String(), "user")
	if err != nil {
		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	refreshToken, err := h.tokenManager.GenerateRefreshToken(uid)
	if err != nil {
		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	response.WithSuccess(ctx, http.StatusOK, messages.SuccessfullyLoggedIn, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	var req models.RefreshRequest
	if err := ctx.BindJSON(&req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			if len(validationErrors) > 0 {
				fe := validationErrors[0]
				var params []string
				switch fe.Tag() {
				case "min", "max":
					params = []string{fe.Param()}
				}
				msg := messages.ErrorMessage{
					Message: messages.MessageForTag(fe.Tag(), params...),
				}
				response.WithError(ctx, http.StatusBadRequest, msg.Message, fe)
				return
			}
		}
		response.WithError(ctx, http.StatusBadRequest, messages.InvalidJSONOrMissingFields, err)
		return
	}

	uid, err := h.tokenManager.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		if errors.Is(err, errors.ErrTokenExpired) {
			response.WithError(ctx, http.StatusUnauthorized, messages.TokenExpired, err)
			return
		}
	}

	err = h.tokenManager.DeleteRefreshToken(uid)
	if err != nil {
		if errors.Is(err, errors.ErrNotFoundRefreshToken) {
			response.WithError(ctx, http.StatusNotFound, messages.TokenNotFound, err)
			return
		}
	}

	response.WithSuccess(ctx, http.StatusOK, messages.SuccessfullyLoggedOut, nil)
}

func (h *AuthHandler) ForgotPassword(ctx *gin.Context) {}

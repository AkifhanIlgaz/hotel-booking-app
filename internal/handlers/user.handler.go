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

type UserHandler struct {
	userService  *services.UserService
	tokenManager *token.Manager
}

func NewUserHandler(userService *services.UserService, tokenManager *token.Manager) *UserHandler {
	return &UserHandler{
		userService:  userService,
		tokenManager: tokenManager,
	}
}

func (uh *UserHandler) Register(ctx *gin.Context) {
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

	id, err := uh.userService.RegisterUser(req)
	if err != nil {
		if errors.Is(err, errors.ErrEmailTaken) {
			response.WithError(ctx, http.StatusConflict, messages.EmailAlreadyRegistered, err)
			return
		}

		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	// Todo: role enum
	accessToken, err := uh.tokenManager.GenerateAccessToken(id.String(), "user")
	if err != nil {
		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	// TODO: Better error handling
	refreshToken, err := uh.tokenManager.GenerateRefreshToken(id)
	if err != nil {
		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	response.WithSuccess(ctx, http.StatusCreated, messages.SuccessfullyRegistered, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}

func (uh *UserHandler) Login(ctx *gin.Context) {
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

	user, err := uh.userService.AuthenticateUser(req)
	if err != nil {
		// Todo: Error constants
		if errors.Is(err, errors.ErrUserNotFound) {
			response.WithError(ctx, http.StatusNotFound, messages.UserNotFound, err)
			return
		}
		if errors.Is(err, errors.ErrWrongPassword) {
			response.WithError(ctx, http.StatusUnauthorized, messages.WrongPassword, err)
			return
		}
	}

	accessToken, err := uh.tokenManager.GenerateAccessToken(user.Id.String(), "user")
	if err != nil {

		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	// TODO: Better error handling
	refreshToken, err := uh.tokenManager.GenerateRefreshToken(user.Id)
	if err != nil {
		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	response.WithSuccess(ctx, http.StatusOK, messages.SuccessfullyLoggedIn, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}

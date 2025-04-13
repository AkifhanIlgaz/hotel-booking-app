package handlers

import (
	"net/http"

	"github.com/AkifhanIlgaz/hotel-booking-app/internal/models"
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/services"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/errors"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/messages"
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
					Field:   fe.Field(),
					Message: messages.MessageForTag(fe.Tag(), params...),
				}
				ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": messages.InvalidJSONOrMissingFields})
		return
	}

	id, err := uh.userService.RegisterUser(req)
	if err != nil {
		if errors.Is(err, errors.ErrEmailTaken) {
			ctx.JSON(http.StatusConflict, gin.H{"error": messages.EmailAlreadyRegistered})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": messages.EmailAlreadyRegistered,
		})
		return
	}

	_ = id

	ctx.JSON(http.StatusOK, gin.H{"message": "Kayıt başarılı!"})

	// TODO: Generate jwt tokens and return to user
}

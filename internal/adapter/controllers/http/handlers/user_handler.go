package handlers

import (
	"github.com/NorthDice/ReflectDiary/internal/usecase/user"
	apperr "github.com/NorthDice/ReflectDiary/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	registerUC *user.RegisterUseCase
	loginUC    *user.LoginUseCase
}

func NewUserHandler(useCase *user.RegisterUseCase, loginUC *user.LoginUseCase) *UserHandler {
	return &UserHandler{
		registerUC: useCase,
		loginUC:    loginUC,
	}
}

func (uh *UserHandler) Register(c *gin.Context) {
	var req user.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		appErr := apperr.NewValidationError("Invalid request format", err.Error())
		uh.handleError(c, appErr)
		return
	}

	//resp, err := uh.registerUC.Register(req)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, err)
	//	return
	//}

	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}

func (uh *UserHandler) handleError(c *gin.Context, err error) {
	if appErr, ok := apperr.IsAppError(err); ok {
		c.JSON(appErr.GetHTTPStatus(), gin.H{
			"success": false,
			"error":   appErr,
		})
		return
	}

	internalErr := apperr.NewInternalError(err.Error())
	c.JSON(internalErr.GetHTTPStatus(), gin.H{
		"success": false,
		"error":   internalErr,
	})
}

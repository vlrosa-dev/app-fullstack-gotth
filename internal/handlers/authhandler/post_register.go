package authhandler

import (
	"app-fullstack-gotth/internal/services/authservices"
	"app-fullstack-gotth/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostRegisterHandler struct {
	authServices *authservices.RegisterServices
}

func NewPostRegisterHandler(authServices *authservices.RegisterServices) *PostRegisterHandler {
	return &PostRegisterHandler{authServices: authServices}
}

func (auth *PostRegisterHandler) Serve(e echo.Context) error {
	var register structs.AuthRegisterFormRequest

	if err := e.Bind(&register); err != nil {
		return e.String(http.StatusBadRequest, "bad request")
	}

	if err := auth.authServices.Register(register); err != nil {
		return err
	}

	return nil
}

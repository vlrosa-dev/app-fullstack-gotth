package authhandler

import (
	"app-fullstack-gotth/share"
	"app-fullstack-gotth/views/authviews"

	"github.com/labstack/echo/v4"
)

type GetRegisterHandler struct {
}

func NewGetRegisterHandler() *GetRegisterHandler {
	return &GetRegisterHandler{}
}

func (auth *GetRegisterHandler) Serve(e echo.Context) error {
	registerView := authviews.Register()
	return share.RenderView(e, authviews.RegisterIndex(registerView))
}

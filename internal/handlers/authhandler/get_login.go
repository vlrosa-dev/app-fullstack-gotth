package authhandler

import (
	"app-fullstack-gotth/internal/views/authviews"
	"app-fullstack-gotth/share"

	"github.com/labstack/echo/v4"
)

type GetLoginHandler struct {
}

func NewGetLoginHandler() *GetLoginHandler {
	return &GetLoginHandler{}
}

func (auth *GetLoginHandler) Serve(e echo.Context) error {
	homeView := authviews.Login()
	return share.RenderView(e, authviews.LoginIndex(homeView))
}

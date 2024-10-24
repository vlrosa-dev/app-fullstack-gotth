package share

import (
	"app-fullstack-gotth/internal/views/errorsviews"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, e echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	var errorPage templ.Component

	switch code {
	case 401:
		errorPage = errorsviews.Error401()
	case 404:
		errorPage = errorsviews.Error404()
	case 500:
		errorPage = errorsviews.Error500()
	}

	RenderView(e, errorPage)
}

package httpHandler

import (
	"github.com/labstack/echo/v4"
)

type mw struct {
}

type IMiddleware interface {
	permission(next echo.HandlerFunc) echo.HandlerFunc
}

func (m mw) permission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
		// clientIP := utils.GetClientIP(c.Request())
		// companyToken := c.Request().Header.Get("company_token")
		// err := m.uc.VerifyPermission(c.Request().Context(), c.Param("token"), companyToken, clientIP)
		// if err != nil {
		// 	if err.Error() == "access forbidden: IP address not whitelisted" && utils.IsLocalhostIP(clientIP) {
		// 		return next(c)
		// 	}
		// 	return ResponseError(http.StatusBadRequest, err.Error(), c)
		// }
		// return next(c)
	}
}

func NewMiddlewareHandler() IMiddleware {
	return mw{}
}

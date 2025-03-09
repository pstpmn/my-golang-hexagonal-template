package httpHandler

import (
	"fmt"
	"net/http"

	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pstpmn/my-golang-hexagonal-template/conf"
	_ "github.com/pstpmn/my-golang-hexagonal-template/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type (
	router struct {
		echo       *echo.Echo
		cfg        conf.App
		handler    IHttpHandler
		middleware IMiddleware
	}
)

func errorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			if he, ok := err.(*echo.HTTPError); ok {
				code := he.Code
				var message string
				switch code {
				case http.StatusNotFound:
					message = "404 Not Found"
				case http.StatusInternalServerError:
					message = "Internal Server Error Message"
				default:
					return ResponseError(code, he.Message.(string), c)
				}
				return ResponseError(code, message, c)
			}
		}
		return nil
	}
}

func NewRouter(cfg conf.App, h IHttpHandler, mw IMiddleware) *router {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Decompress())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))
	e.Use(echoPrometheus.MetricsMiddleware())
	// e.Static("/assets/blog", "public/assets/blog")
	e.Use(errorHandler)

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/document/*", echoSwagger.WrapHandler)
	r := e.Group("v1")
	{
		user := r.Group("/user")
		{
			user.GET("", h.GetAll)
			user.GET("/:userId", h.GetUserById)
			r.GET("/users", h.GetAll)
		}

		// admin := r.Group("/admin")
		// {
		// }
	}

	return &router{
		echo:       e,
		cfg:        cfg,
		handler:    h,
		middleware: mw,
	}
}

func (r *router) Serve() {
	if err := r.echo.Start(r.cfg.Host); err != http.ErrServerClosed {
		panic(fmt.Errorf("error http server %s", err))
	}
}

func CustomHTTPErrorHandler(err error, ctx echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		code := he.Code
		var message string
		switch code {
		case http.StatusNotFound:
			message = "Custom 404 Not Found Message"
		case http.StatusInternalServerError:
			message = "Custom 500 Internal Server Error Message"
		default:
			// For other HTTP errors, return the default error response
			ResponseError(code, he.Message.(string), ctx)
			return
		}
		ResponseError(code, message, ctx)
		return
	}
	ResponseError(http.StatusInternalServerError, "Internal Server Error", ctx)
	ctx.Error(err)
}

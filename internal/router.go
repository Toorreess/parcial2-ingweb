package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, ctr IController) *echo.Echo {
	api := e.Group("/api")
	v1 := api.Group("/v1")

	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"*"},
	}))

	v1.GET("/e/:id", func(c echo.Context) error {
		return ctr.Get(c)
	})

	v1.POST("/e", func(c echo.Context) error {
		return ctr.Create(c)
	})

	v1.PUT("/e/:id", func(c echo.Context) error {
		obj := make(map[string]interface{})
		c.Bind(&obj)
		return ctr.Update(c, obj)
	})

	v1.DELETE("/e/:id", func(c echo.Context) error {
		return ctr.Delete(c)
	})

	v1.GET("/e", func(c echo.Context) error {
		return ctr.List(c)
	})

	return e
}

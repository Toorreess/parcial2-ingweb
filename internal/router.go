package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, ctr AppController) *echo.Echo {
	api := e.Group("/api")
	v1 := api.Group("/v1")

	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"*"},
	}))

	/* Entity1 endpoints */
	v1.GET("/e1/:id", func(c echo.Context) error {
		return ctr.Entity1.Get(c)
	})

	v1.POST("/e1", func(c echo.Context) error {
		return ctr.Entity1.Create(c)
	})

	v1.PUT("/e1/:id", func(c echo.Context) error {
		obj := make(map[string]interface{})
		c.Bind(&obj)
		return ctr.Entity1.Update(c, obj)
	})

	v1.DELETE("/e1/:id", func(c echo.Context) error {
		return ctr.Entity1.Delete(c)
	})

	v1.GET("/e1", func(c echo.Context) error {
		return ctr.Entity1.List(c)
	})

	/* Entity2 endpoints */
	v1.GET("/e2/:id", func(c echo.Context) error {
		return ctr.Entity2.Get(c)
	})

	v1.POST("/e2", func(c echo.Context) error {
		return ctr.Entity2.Create(c)
	})

	v1.PUT("/e2/:id", func(c echo.Context) error {
		obj := make(map[string]interface{})
		c.Bind(&obj)
		return ctr.Entity2.Update(c, obj)
	})

	v1.DELETE("/e2/:id", func(c echo.Context) error {
		return ctr.Entity2.Delete(c)
	})

	v1.GET("/e2", func(c echo.Context) error {
		return ctr.Entity2.List(c)
	})

	return e
}

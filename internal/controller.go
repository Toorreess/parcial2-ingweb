package internal

import (
	"net/http"
	"parcial2-ingweb/internal/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Context echo.Context

type IController interface {
	Create(c Context) error
	Get(c Context) error
	Update(c Context, body map[string]interface{}) error
	Delete(c Context) error

	List(c Context) error
}

type controller struct {
	Interactor IInteractor
}

func NewController(i IInteractor) IController {
	return &controller{i}
}

func (ctr *controller) Create(c Context) error {
	var entity *model.Entity

	if err := c.Bind(&entity); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	entity, err := ctr.Interactor.Create(entity)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	return c.JSON(http.StatusCreated, entity)
}

func (ctr *controller) Get(c Context) error {
	var entity *model.Entity

	entity, err := ctr.Interactor.Get(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, entity)
}

func (ctr *controller) Update(c Context, body map[string]interface{}) error {
	var entity *model.Entity

	entity, err := ctr.Interactor.Update(c.Param("id"), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, entity)
}

func (ctr *controller) Delete(c Context) error {
	err := ctr.Interactor.Delete(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.NoContent(http.StatusNoContent)
}

func (ctr *controller) List(c Context) error {
	query := c.QueryParams()

	q := make(map[string]string)
	for k, v := range query {
		if k != "limit" && k != "offset" && k != "orderBy" && k != "order" {
			q[k] = v[0]
		}
	}

	limitStr := query.Get("limit")
	if limitStr == "" {
		limitStr = "20"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"status": "Limit must be a number"})
	}

	offsetStr := query.Get("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"status": "Offset must be a number"})
	}

	list, err := ctr.Interactor.List(q, limit, offset, query.Get("orderBy"), query.Get("order"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, list)
}

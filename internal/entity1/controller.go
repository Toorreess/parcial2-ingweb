package entity1

import (
	"net/http"
	"parcial2-ingweb/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Context echo.Context

type IE1Controller interface {
	Create(c Context) error
	Get(c Context) error
	Update(c Context, body map[string]interface{}) error
	Delete(c Context) error

	List(c Context) error
}

type e1Controller struct {
	e1Interactor IE1Interactor
}

func NewE1Controller(i IE1Interactor) IE1Controller {
	return &e1Controller{i}
}

func (ec *e1Controller) Create(c Context) error {
	var entity *models.Entity1

	if err := c.Bind(&entity); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	entity, err := ec.e1Interactor.Create(entity)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	return c.JSON(http.StatusCreated, entity)
}

func (ec *e1Controller) Get(c Context) error {
	var entity *models.Entity1

	entity, err := ec.e1Interactor.Get(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, entity)
}

func (ec *e1Controller) Update(c Context, body map[string]interface{}) error {
	var entity *models.Entity1

	entity, err := ec.e1Interactor.Update(c.Param("id"), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, entity)
}

func (ec *e1Controller) Delete(c Context) error {
	err := ec.e1Interactor.Delete(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.NoContent(http.StatusNoContent)
}

func (ec *e1Controller) List(c Context) error {
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

	list, err := ec.e1Interactor.List(q, limit, offset, query.Get("orderBy"), query.Get("order"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, list)
}

package handler

import (
	"net/http"
	"strings"

	"hhcshare/model"

	"github.com/labstack/echo/v4"
)

// ListOffices
func (s HandlerStore) ListOffices(context echo.Context) error {
	a := context.QueryParam("a")

	var oo any
	var err error

	if strings.ToLower(a) == "full" {
		oo, err = model.ListOfficesFull(s.Connection)
	} else {
		oo, err = model.ListOffices(s.Connection)
	}

	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, oo)
}

// OfficeGOfficeGeometryeoJSON
func (s HandlerStore) OfficeGeometry(context echo.Context) error {
	office := context.Param("office")
	officeLower := strings.ToUpper(office)
	oo, err := model.OfficeGeometry(s.Connection, officeLower)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, oo)
}

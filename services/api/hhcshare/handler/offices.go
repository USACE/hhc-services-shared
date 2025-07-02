package handler

import (
	"fmt"
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
	options := []string{"cw", "fuds", "mil", "reg"}
	office := strings.ToUpper(context.Param("office"))
	aor := strings.ToLower(context.QueryParam("aor"))
	var aorTable string

	// check the aor
	if aor == "" {
		aor = "cw"
	}
	isValid := false
	for _, option := range options {
		if aor == option {
			isValid = true
			aorTable = "office_aor_" + option
			break
		}
	}
	if !isValid {
		msg := fmt.Sprintf("Mission option '%s' not available", aor)
		return context.JSON(http.StatusInternalServerError, map[string]string{"message": msg})
	}

	oo, err := model.OfficeGeometry(s.Connection, office, aorTable)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, oo)
}

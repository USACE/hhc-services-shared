package handler

import (
	"net/http"
	"strings"

	"github.com/hhc-services-shared/services/api/model"

	"github.com/labstack/echo/v4"
)

// ListOffices
func (s HandlerStore) ListOffices(context echo.Context) error {
	allowed := []string{"office_aor", "parent_office"}
	filtered := make([]string, 0)

	// get the column names, split on comma, trim spaces, and add to filtered if allowed
	n := context.QueryParam("property")
	for n := range strings.SplitSeq(n, ",") {
		nt := strings.TrimSpace(n)
		for _, a := range allowed {
			if a == nt {
				filtered = append(filtered, nt)
			}
		}
	}

	oo, err := model.ListOffices(s.Connection, filtered)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, oo)
}

// OfficeGeometry
func (s HandlerStore) OfficeGeometry(context echo.Context) error {
	allowed := []string{"military", "civil_works", "fuds", "regulatory"}
	office := strings.ToUpper(context.Param("office"))

	filtered := make([]string, 0)
	aors := context.QueryParam("aor")
	for aor := range strings.SplitSeq(aors, ",") {
		aort := strings.TrimSpace(aor)
		for _, a := range allowed {
			if a == aort {
				filtered = append(filtered, aort)
			}
		}
	}

	if len(filtered) == 0 {
		filtered = append(filtered, "civil_works")
	}
	oo, err := model.OfficeGeometry(s.Connection, office, filtered)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, oo)
}

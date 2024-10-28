package rest

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:generate mockery --name BmiService
type BmiService interface {
	Calculate(ctx context.Context, weight float64, height float64) (float64, error)
}

type BmiHandler struct {
	Service BmiService
}

func NewBmiHandler(e *echo.Echo, svc BmiService) {
	handler := &BmiHandler{
		Service: svc,
	}
	e.POST("/calculate-bmi", handler.Calculate)
}

func (h *BmiHandler) Calculate(c echo.Context) error {
	ctx := c.Request().Context()

	type Request struct {
		Weight float64 `json:"weight"`
		Height float64 `json:"height"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bmi, err := h.Service.Calculate(ctx, req.Weight, req.Height)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	resp := map[string]interface{}{
		"bmi": bmi,
	}
	return c.JSON(http.StatusOK, resp)
}

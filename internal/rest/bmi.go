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

type CalculateRequest struct {
	Weight float64 `json:"weight"` // น้ำหนัก (kg)
	Height float64 `json:"height"` // ส่วนสูง (meters)
}

type CalculateResponse struct {
	BMI float64 `json:"bmi"`
}

// Calculate godoc
// @Summary คำนวณค่า BMI
// @Description คำนวณค่า BMI จากน้ำหนักและส่วนสูง
// @Tags BMI
// @Accept json
// @Produce json
// @Param Request body CalculateRequest true "ข้อมูลการคำนวณ BMI"
// @Success 200 {object} CalculateResponse
// @Failure 500 {object} ResponseError
// @Router /calculate-bmi [post]
func (h *BmiHandler) Calculate(c echo.Context) error {
	ctx := c.Request().Context()

	req := CalculateRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	bmi, err := h.Service.Calculate(ctx, req.Weight, req.Height)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	resp := CalculateResponse{bmi}
	return c.JSON(http.StatusOK, resp)
}

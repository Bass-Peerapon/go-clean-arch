package rest_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/go-clean-arch/internal/rest"
	"github.com/bxcodec/go-clean-arch/internal/rest/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCalculateBmi(t *testing.T) {
	type Request struct {
		Weight float64 `json:"weight"`
		Height float64 `json:"height"`
	}

	mockRequest := Request{
		Weight: 70,
		Height: 1.75,
	}
	mockUCase := new(mocks.BmiService)

	j, err := json.Marshal(mockRequest)
	assert.NoError(t, err)

	mockUCase.On("Calculate", mock.Anything, mock.AnythingOfType("float64"), mock.AnythingOfType("float64")).Return(22.857142857142858, nil)

	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(), echo.POST, "/calculate-bmi", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/calculate-bmi")

	handler := rest.BmiHandler{
		Service: mockUCase,
	}
	err = handler.Calculate(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

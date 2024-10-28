package bmi

import (
	"context"
	"errors"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Calculate(ctx context.Context, weight float64, height float64) (float64, error) {
	if height <= 0 {
		return 0, errors.New("height must be greater than zero")
	}
	bmi := weight / (height * height)
	return bmi, nil
}

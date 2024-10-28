package bmi_test

import (
	"context"
	"testing"

	"github.com/bxcodec/go-clean-arch/bmi"
)

func TestService_Calculate(t *testing.T) {
	service := bmi.NewService()
	ctx := context.Background()

	tests := []struct {
		name    string
		weight  float64
		height  float64
		wantBMI float64
		wantErr bool
	}{
		{
			name:    "Normal BMI",
			weight:  70,
			height:  1.75,
			wantBMI: 22.857142857142858,
			wantErr: false,
		},
		{
			name:    "Zero Height",
			weight:  70,
			height:  0,
			wantBMI: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBMI, err := service.Calculate(ctx, tt.weight, tt.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && gotBMI != tt.wantBMI {
				t.Errorf("Calculate() = %v, want %v", gotBMI, tt.wantBMI)
			}
		})
	}
}

package main

import "testing"

func TestHitungHargaTotal(t *testing.T) {
	tests := []struct {
		name    string
		product Product
		want    float64
		wantErr bool
	}{
		{
			name: "failed in item price == 0",
			product: Product{
				Name:         "Test Product",
				Price:        0,
				Quantity:     1,
				ShippingCost: 1000,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failed in qty item == 0",
			product: Product{
				Name:         "Test Product",
				Price:        100,
				Quantity:     0,
				ShippingCost: 1000,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.product.CalculateTotalPrice()
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateTotalPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateTotalPrice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatePayment(t *testing.T) {
	tests := []struct {
		name    string
		payment Payment
		wantErr bool
	}{
		{
			name: "failed in total price == 0",
			payment: Payment{
				TotalPrice:    0,
				PaymentMethod: "cod",
				Installment:   true,
			},
			wantErr: true,
		},
		{
			name: "failed in the payment method does not exist",
			payment: Payment{
				TotalPrice:    4000,
				PaymentMethod: "ngutang",
				Installment:   true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.payment.ValidatePayment()
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

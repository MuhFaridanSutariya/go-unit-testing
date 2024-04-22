package main

import (
	"errors"
	"fmt"
)

type Product struct {
	Name         string
	Price        float64
	Quantity     int
	ShippingCost float64
}

type Payment struct {
	TotalPrice    float64
	PaymentMethod string
	Installment   bool
}

type Order struct {
	Product
	Payment
}

const (
	TaxRate          = 10
	AdditionalCharge = 2000
)

func (p *Product) CalculateTotalPrice() (float64, error) {
	if p.Price <= 0 {
		return 0, errors.New("harga barang tidak boleh nol")
	}

	if p.Quantity <= 0 {
		return 0, errors.New("jumlah barang tidak boleh nol")
	}

	totalItemPrice := p.Price * float64(p.Quantity)
	totalAfterShipping := totalItemPrice + p.ShippingCost
	tax := totalItemPrice * TaxRate / 100
	total := totalAfterShipping + tax + AdditionalCharge

	return total, nil
}

func (p *Payment) ValidatePayment() error {
	if p.TotalPrice <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	validMethods := []string{"cod", "transfer", "debit", "credit", "gerai"}
	valid := false
	for _, method := range validMethods {
		if p.PaymentMethod == method {
			valid = true
			break
		}
	}
	if !valid {
		return errors.New("metode tidak dikenali")
	}

	if p.Installment {
		if p.PaymentMethod != "credit" || p.TotalPrice < 500000 {
			return errors.New("cicilan tidak memenuhi syarat")
		}
	} else {
		if p.PaymentMethod == "credit" {
			return errors.New("credit harus dicicil")
		}
	}

	return nil
}

func main() {
	product := Product{Name: "Laptop", Price: 10000000, Quantity: 1, ShippingCost: 100000}
	totalPrice, err := product.CalculateTotalPrice()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	payment := Payment{TotalPrice: totalPrice, PaymentMethod: "credit", Installment: true}
	err = payment.ValidatePayment()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	order := Order{Product: product, Payment: payment}
	fmt.Println("Order berhasil:", order)
}

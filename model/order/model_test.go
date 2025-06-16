package model

import (
	"testing"
)

// func setup() {
// 	LoadData("../../data.json")
// }

func TestValidateCoupon(t *testing.T) {
	// setup()
	o := Order{CouponCode: "SUPER100"}

	//validate Coupon
	couponValid := o.ValidateCoupon("SUPER100")
	if couponValid {
		t.Errorf("Expected invalid coupon reponse. Got %v", couponValid)
	}
}

func TestSave(t *testing.T) {
	// setup()
	o := Order{CouponCode: "FIFTYOFF"}
	items := []Item{{ProductId: "1", Quantity: 1}}
	req := OrderReq{CouponCode: "FIFTYOFF", Items: items}

	//Save Order
	err := o.Save(req)
	if err != nil {
		t.Errorf("error while saving order %v ", err)
	}

	if o.Id == "" {
		t.Errorf("Expected id. Got %v", o.Id)
	}

	if len(*o.Products) < 1 {
		t.Errorf("Expected products in order. Got %v", *o.Products)
	}
}

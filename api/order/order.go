package api

import (
	"encoding/json"
	"net/http"
	"oolio/app"
	model "oolio/model/order"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := r.Header.Get("x-api-key")

	if key != app.Config.ApiKey {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req model.OrderReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	o := &model.Order{CouponCode: req.CouponCode, Items: req.Items}

	//validate Coupon
	couponValid := o.ValidateCoupon(req.CouponCode)
	if !couponValid {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//Save in DB
	err := o.Save(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error in saving order " + err.Error()))
		return
	}

	//Encode response in http Write
	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	o := &model.Order{}
	res := o.Get()
	//Encode response in http Write
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

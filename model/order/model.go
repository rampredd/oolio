package model

import (
	"bufio"
	"log"
	model "oolio/model/product"
	"os"

	"github.com/google/uuid"
)

type Item struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type OrderReq struct {
	CouponCode string `json:"couponCode"`
	Items      []Item `json:"items"`
}

type Order struct {
	Id         string           `json:"id"`
	CouponCode string           `json:"couponCode"`
	Items      []Item           `json:"items"`
	Products   *[]model.Product `json:"products"`
}

var (
	orders    []Order
	couponMap map[string]map[string]bool
)

func LoadCoupons(location string) error {
	log.Printf("Loading coupon files from %s into DB ", location)

	fileNames := []string{"couponbase1", "couponbase2", "couponbase3"}
	keyNames := []string{"c1", "c2", "c3"}
	fileName := ""
	couponMap = make(map[string]map[string]bool)

	for ind, v := range fileNames {
		fileName = location + v

		log.Printf("Loading file %s", fileName)

		f, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		couponMap[keyNames[ind]] = make(map[string]bool)

		//append all keys
		for scanner.Scan() {
			couponMap[keyNames[ind]][scanner.Text()] = false
		}
	}

	log.Printf("Coupon codes loading completed")
	return nil
}

func (o *Order) ValidateCoupon(coupon string) bool {
	keyNames := []string{"c1", "c2", "c3"}
	findCouponCount := 0
	for _, k := range keyNames {
		if _, ok := couponMap[k][coupon]; ok {
			findCouponCount += 1
		}
	}

	//coupon is not valid if not present in atleast in 2 files
	if findCouponCount < 2 {
		return false
	}
	return true
}

/*
Get List of products
Returns list of products
*/
func (o *Order) Save(r OrderReq) error {
	//generate uuid
	o.Id = uuid.New().String()

	var err error
	p := &model.Product{}
	ids := []string{}
	for _, i := range r.Items {
		ids = append(ids, i.ProductId)
	}

	//get list of products
	o.Products, err = p.GetProductList(ids)
	if err != nil {
		return err
	}
	orders = append(orders, *o)
	return nil
}

func (o *Order) Get() []Order {
	return orders
}

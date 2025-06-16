package model

import (
	"testing"
)

func setup() {
	LoadData("../../data.json")
}

func TestGetProductList(t *testing.T) {
	setup()
	p := Product{}

	//Get list of products from Data set
	ids := []string{}
	response, err := p.GetProductList(ids)
	if err != nil {
		t.Errorf("error while fetching list of products %v ", err)
	}

	if len(*response) < 1 {
		t.Errorf("Expected array with data. Got %v", *response)
	}
}

func TestGetProductInfo(t *testing.T) {
	setup()
	p := Product{Id: "1"}

	//Get list of products from Data set
	err := p.GetProductInfo()
	if err != nil {
		t.Errorf("error while fetching details of one product %v ", err)
	}

	if len(p.Id) < 1 {
		t.Errorf("Expected array with data. Got %v", p.Id)
	}
}

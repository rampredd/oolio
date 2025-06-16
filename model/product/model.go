package model

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"sort"
)

type Image struct {
	Thumbnail string `json:"thumbnail"`
	Mobile    string `json:"mobile"`
	Tablet    string `json:"tablet"`
	Desktop   string `json:"desktop"`
}

type Product struct {
	Id       string `json:"id"`
	Image    `json:"image"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float32 `json:"price"`
}

var (
	MapOfProducts map[string]Product
	pIds          []string
)

func LoadData(location string) error {
	// location := "./data.json"
	log.Printf("Loading config file %s", location)
	var products []Product
	f, err := os.Open(location)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = json.NewDecoder(f).Decode(&products); err != nil {
		return err
	}
	MapOfProducts = make(map[string]Product)
	for _, p := range products {
		MapOfProducts[p.Id] = p
	}

	pIds = make([]string, 0)
	for _, val := range MapOfProducts {
		pIds = append(pIds, val.Id)
	}

	//sort ids in ascending order
	sort.Strings(pIds)

	return nil
}

/*
Get List of products
Returns list of products
*/
func (P *Product) GetProductList(ids []string) (*[]Product, error) {
	products := make([]Product, 0)
	if len(ids) > 0 {
		for _, id := range ids {
			products = append(products, MapOfProducts[id])
		}
	} else {
		for _, id := range pIds {
			products = append(products, MapOfProducts[id])
		}
	}
	return &products, nil
}

/*
Get Details of one product
Returns product info
*/
func (P *Product) GetProductInfo() error {
	val, ok := MapOfProducts[P.Id]
	if !ok {
		return errors.New(ProductNotFound)
	}
	P.Category = val.Category
	P.Name = val.Name
	P.Price = val.Price
	P.Image = val.Image
	return nil
}

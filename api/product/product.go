package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	model "oolio/model/product"
)

/*
API route: api/product
Response: 200
*/
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := &model.Product{}

	//Get list of products from Data set
	ids := []string{}
	response, err := p.GetProductList(ids)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//Encode response in http Write
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

/*
API route: api/product/{productId}
Response: 200 if success

	400 if productId not found in request
	404 if product not found in Data
*/
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, ok := params["productId"]

	if !ok {
		err := json.NewEncoder(w).Encode(map[string]string{"description": model.InvalidProductId})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	p := &model.Product{Id: id}

	//Get details of one product from Data set
	err := p.GetProductInfo()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	//Encode response in http Write
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

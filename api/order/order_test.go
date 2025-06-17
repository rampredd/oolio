package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"oolio/app"
	model "oolio/model/order"
	"testing"

	"github.com/gorilla/mux"
)

var (
	Router *mux.Router
)

func setUp() {
	Router = mux.NewRouter()
	if err := app.LoadConfig("../../config.json"); err != nil {
		log.Fatal(err)
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestCreateOrder(t *testing.T) {
	setUp()
	items := []model.Item{{ProductId: "1", Quantity: 1}}
	body := &model.OrderReq{CouponCode: "FIFTYOFF", Items: items}
	order, err := json.Marshal(body)
	if err != nil {
		t.Errorf("impossible to marshall order: %s", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8020/api/order", bytes.NewReader(order))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", "secretkey")
	if err != nil {
		t.Errorf("error in preparing POST /api/order api %v", err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("error in reponse %v", err.Error())
	}

	checkResponseCode(t, http.StatusOK, res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error in reading body %v", err)
	}

	if len(resBody) < 1 {
		t.Errorf("body expected with some data. Got empty array")
	}
}

func TestCreateOrderBadRequest(t *testing.T) {
	req, err := http.NewRequest("POST", "http://localhost:8020/api/order", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", "secretkey")
	if err != nil {
		t.Errorf("error in preparing POST /api/order api %v", err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("error in reponse %v", err.Error())
	}

	checkResponseCode(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateOrderUnauthorized(t *testing.T) {
	req, err := http.NewRequest("POST", "http://localhost:8020/api/order", nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("error in preparing POST /api/order api %v", err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("error in reponse %v", err.Error())
	}

	checkResponseCode(t, http.StatusUnauthorized, res.StatusCode)
}

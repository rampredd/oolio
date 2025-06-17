package api

import (
	"io"
	"log"
	"net/http"
	"oolio/app"
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

func TestGetProducts(t *testing.T) {
	setUp()

	req, err := http.NewRequest("GET", "http://localhost:8020/api/product", nil)
	if err != nil {
		t.Errorf("error in preparing GET /api/product api %v", err.Error())
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

func TestGetProduct(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8020/api/product/1", nil)
	if err != nil {
		t.Errorf("error in preparing GET /api/product/1 api %v", err.Error())
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
		t.Errorf("body expected with some data. Got empty byte array")
	}
}

func TestGetProductNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8020/api/product/12", nil)
	if err != nil {
		t.Errorf("error in preparing GET /api/product/12 api %v", err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("error in reponse %v", err.Error())
	}

	checkResponseCode(t, http.StatusNotFound, res.StatusCode)
}

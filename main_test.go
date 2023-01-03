package main

import (
	"bytes"
	"go-learn/controller/cake"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/cakes", cake.Create).Methods("POST")
	router.HandleFunc("/cakes", cake.GetAll).Methods("GET")
	router.HandleFunc("/cakes/{id}", cake.Update).Methods("PUT")
	router.HandleFunc("/cakes/{id}", cake.GetOne).Methods("GET")
	router.HandleFunc("/cakes/{id}", cake.Delete).Methods("DELETE")
	return router
}
func TestDisplayCakes(t *testing.T) {
	request, _ := http.NewRequest("GET", "/cakes", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestDisplayCake(t *testing.T) {
	request, _ := http.NewRequest("GET", "/cakes/6", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestInsertCake(t *testing.T) {
	testBody := `{
		"title": "Lemon Testing",
		"description": "A cheesecake made of lemon",
		"rating": 7,
		"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
	  }
	  `
	request, _ := http.NewRequest("POST", "/cakes", bytes.NewBufferString(testBody))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	checkResponseCode(t, http.StatusCreated, response.Code)
}

func TestUpdateCake(t *testing.T) {
	testBody := `{
		"title": "test",
		"description": "update dulu la",
		"rating": 7,
		"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
	  }
	  `
	request, _ := http.NewRequest("PUT", "/cakes/5", bytes.NewBufferString(testBody))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestDeleteCake(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/cakes/5", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

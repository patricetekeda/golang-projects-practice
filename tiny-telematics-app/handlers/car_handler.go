package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/patriceteda/tiny-telematics-app/models"
	"github.com/patriceteda/tiny-telematics-app/services"
)

var cars []models.Car

// RegisterCar handles POST /register
func RegisterCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); 
	err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	cars = append(cars, car)

	// Notify the analytics service
	go services.NotifyAnalytics(car.VIN)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Car registered successfully: %+v\n", car)
}

// ListCars handles GET /cars
func ListCars(w http.ResponseWriter, r *http.Request) {
	if len(cars) == 0 {
		http.Error(w, "No cars registered", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

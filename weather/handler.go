package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}

	getWeather(w, r)
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var coords Coordinates

	if err := json.NewDecoder(r.Body).Decode(&coords); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Latitude: %f, Longitude: %f\n", coords.Latitude, coords.Longitude)

	points, err := FetchPoints(r.Context(), coords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the weather data as JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(points)
}

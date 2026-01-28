package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		getWeather(w, r)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var coords Coordinates

	if err := json.NewDecoder(r.Body).Decode(&coords); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Latitude: %f, Longitude: %f\n", coords.Latitude, coords.Longitude)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "success"})
}

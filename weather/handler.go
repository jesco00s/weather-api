package weather

import (
	"encoding/json"
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

	points, err := FetchPoints(r.Context(), coords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	forecast, errForecast := FetchForecast(r.Context(), points)
	if errForecast != nil {
		http.Error(w, errForecast.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(forecast)
}

package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const nationalWeatherPointsUrl = "https://api.weather.gov/points/%f,%f"

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

	points, err := fetchPoints(coords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the weather data as JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(points)
}

func fetchPoints(coords Coordinates) (PointResponse, error) {
	var pointResponse PointResponse
	var pointsUrl = fmt.Sprintf(nationalWeatherPointsUrl, coords.Latitude, coords.Longitude)

	req, err := http.NewRequest(http.MethodGet, pointsUrl, nil)
	if err != nil {
		return pointResponse, err
	}

	req.Header.Set("Accept", "application/geo+json")
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return pointResponse, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return pointResponse, fmt.Errorf("Error retrieving points from NWS status=%s body=%s", resp.Status, string(b))
	}

	err = json.NewDecoder(resp.Body).Decode(&pointResponse)
	if err != nil {
		return pointResponse, fmt.Errorf("Error unmarshalling response body. Error: %w ", err)
	}

	return pointResponse, nil
}

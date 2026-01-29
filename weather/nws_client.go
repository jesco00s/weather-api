package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func FetchPoints(ctx context.Context, coords Coordinates) (Grid, error) {
	const pointsURL = "https://api.weather.gov/points/%f,%f"
	var url = fmt.Sprintf(pointsURL, coords.Latitude, coords.Longitude)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Grid{}, err
	}

	req.Header.Set("Accept", "application/geo+json")
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return Grid{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return Grid{}, fmt.Errorf("nws points request failed: status=%s body=%s", resp.Status, string(b))
	}

	var pointResponse PointResponse
	err = json.NewDecoder(resp.Body).Decode(&pointResponse)
	if err != nil {
		return Grid{}, fmt.Errorf("failed to decode nws points response: %w ", err)
	}

	return pointResponse.Properties, nil
}

func FetchForecast(ctx context.Context, grid Grid) (WeatherResponse, error) {
	const forecastURL = "https://api.weather.gov/gridpoints/TOP/%d,%d/forecast"
	var url = fmt.Sprintf(forecastURL, grid.GridX, grid.GridY)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return WeatherResponse{}, err
	}

	req.Header.Set("Accept", "application/geo+json")
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return WeatherResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return WeatherResponse{}, fmt.Errorf("nws forecast request failed: status=%s body=%s", resp.Status, string(b))
	}

	var forecastResponse ForecastResponse
	err = json.NewDecoder(resp.Body).Decode(&forecastResponse)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("failed to decode nws forecast response: %w ", err)
	}

	return getForecast(forecastResponse.Properties.Periods)
}

func getForecast(periods []Period) (WeatherResponse, error) {
	var weatherResponse WeatherResponse

	if periods[0].Temperature >= 80 {
		weatherResponse.TemperatureDescription = "hot"
	} else if periods[0].Temperature >= 60 {
		weatherResponse.TemperatureDescription = "moderate"
	} else {
		weatherResponse.TemperatureDescription = "cold"
	}
	weatherResponse.ShortForecast = periods[0].ShortForecast

	return weatherResponse, nil
}

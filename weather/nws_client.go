package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const nationalWeatherPointsURL = "https://api.weather.gov/points/%f,%f"

func FetchPoints(ctx context.Context, coords Coordinates) (Grid, error) {
	var pointResponse PointResponse
	var url = fmt.Sprintf(nationalWeatherPointsURL, coords.Latitude, coords.Longitude)

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

	err = json.NewDecoder(resp.Body).Decode(&pointResponse)
	if err != nil {
		return Grid{}, fmt.Errorf("failed to decode nws points response: %w ", err)
	}

	return pointResponse.Properties, nil
}

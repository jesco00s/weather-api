package weather

import "net/http"

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		getWeather(w)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func getWeather(w http.ResponseWriter) {
	w.Write([]byte("success"))
}

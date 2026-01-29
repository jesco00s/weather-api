# weather-api

A simple API that returns weather data for a given set of geographic coordinates.
Data is sourced from the U.S. National Weather Service.

## Instructions for running
- run `go run cmd/main.go`
- make a POST request to url `http://localhost:8080/weather`
- in the body put 
`{
  "latitude": 39.7456,
  "longitude": -97.0892
}`
- can also see a sample request in `bruno/weather-api/Weather Get.bru`
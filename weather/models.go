package weather

type Coordinates struct {
	Latitude  float32 `json:"latitude`
	Longitude float32 `json:"longitude`
}

type PointResponse struct {
	Properties PropertiesResponse `json:"properties"`
}

type PropertiesResponse struct {
	GridX int `json:"gridX"`
	GridY int `json:"gridY"`
}

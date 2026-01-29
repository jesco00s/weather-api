package weather

type Coordinates struct {
	Latitude  float32 `json:"latitude`
	Longitude float32 `json:"longitude`
}

type PointResponse struct {
	Properties Grid `json:"properties"`
}

type Grid struct {
	GridX int `json:"gridX"`
	GridY int `json:"gridY"`
}

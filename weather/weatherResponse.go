package weather

type PointResponse struct {
	Properties PropertiesResponse `json:"properties"`
}

type PropertiesResponse struct {
	GridX int `json:"gridX"`
	GridY int `json:"gridY"`
}

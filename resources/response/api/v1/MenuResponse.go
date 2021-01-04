package v1Response

type MenuResponse struct {
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Image    string  `json:"image,omitempty"`
	Status   string  `json:"status"`
	Category string  `json:"category"`
}

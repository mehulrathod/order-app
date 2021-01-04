package v1Response

type OrderResponse struct {
	Id       uint64              `json:"id,"`
	UserId   uint64              `json:"user_id,omitempty"`
	MenuData string              `json:"MenuData"`
	Status   string              `json:"status"`
	Total    string              `json:"total"`
	//Items    []OrderMenuResponse `json:"items"`
}

package v1Response

//import "time"

type UserResponse struct {
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Image  string `json:"image,omitempty"`
	Mobile int64  `json:"mobile,omitempty"`
}

type LoginResponse struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Image  string `json:"image,omitempty"`
	Mobile int64  `json:"mobile,omitempty"`
	Token  string `json:"token"`
}

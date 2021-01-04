package apiHelpers

import (
	"encoding/json"
	"net/http"
)

const (
	ResponseSuccess = 1
	ResponseError = 0
)

type ResponseData struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{})  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetError(field, tag string) string {
	if tag == "required" {
		return field + " is " + tag
	} else if tag == "email" {
		return field + "not valid"
	}
	return ""
}
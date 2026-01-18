package models

type APIResponse struct {
	Response string `json:"Response"`
	Error    string `json:"Error,omitempty"`
}

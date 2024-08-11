package models

type BasicResp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type BasicRespMesg struct {
	Message string `json:"message"`
}

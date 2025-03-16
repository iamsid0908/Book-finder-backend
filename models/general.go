package models

type BasicResp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type BasicRespWithMeta struct {
	Message string         `json:"message"`
	Data    interface{}    `json:"data,omitempty"`
	Meta    MetaPagination `json:"meta,omitempty"`
}
type MetaPagination struct {
	PageNumber   int64 `json:"page_number"`
	PageSize     int64 `json:"page_size"`
	TotalPages   int64 `json:"total_pages"`
	TotalRecords int64 `json:"total_records"`
}

type BasicRespMesg struct {
	Message string `json:"message"`
}

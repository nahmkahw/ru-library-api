package entity

type Response struct {
	Code         string      `json:"code"`
	Data         interface{} `json:"data,omitempty"`
	HttpCode     int         `json:"http_code"`
	ErrorMessage string      `json:"error_message"`
}

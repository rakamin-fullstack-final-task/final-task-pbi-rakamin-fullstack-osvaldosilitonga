package dto

type WebResponse struct {
	Code     int    `json:"code"`
	Data     any    `json:"data"`
	Messages string `json:"messages"`
	Errors   string `json:"errors"`
}

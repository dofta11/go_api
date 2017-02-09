package error

type HttpError struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

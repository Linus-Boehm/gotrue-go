package schema

type APIError struct {
	StatusCode int    `json:"code"`
	Message    string `json:"msg"`
}

func (e *APIError) Error() string {
	return e.Message
}

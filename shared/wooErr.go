package shared

import (
	"encoding/json"
	"fmt"
)

type WooError struct {
	HttpStatusCode int
	Code           int
	Message        string
}

func NewWooError(statusCode int, raw []byte) *WooError {
	data := map[string]any{}

	err := json.Unmarshal(raw, &data)
	if err != nil {
		return &WooError{
			HttpStatusCode: statusCode,
			Code:           -1,
			Message:        fmt.Sprintf("%s", raw),
		}
	}

	code := -1
	message := "A unknown error occured"

	if parsedCode, ok := data["code"].(int); ok {
		code = parsedCode
	}

	if parsedMessage, ok := data["message"].(string); ok {
		message = parsedMessage
	}

	return &WooError{
		HttpStatusCode: statusCode,
		Code:           code,
		Message:        message,
	}
}

func (e *WooError) Error() string {
	return e.Message
}

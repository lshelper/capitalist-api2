package capitalist

import "fmt"

type APIError struct {
	Message    string
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	if e.StatusCode == 0 {
		return e.Message
	}
	return fmt.Sprintf("capitalist api request failed: status %d: %s", e.StatusCode, e.Message)
}

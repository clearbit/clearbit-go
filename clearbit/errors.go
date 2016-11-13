package clearbit

import (
	"fmt"
)

// APIError represents a Clearbit API Error response
// https://clearbit.com/docs#errors
type APIError struct {
	Errors []ErrorDetail `json:"error"`
}

// ErrorDetail represents an individual item in an APIError.
type ErrorDetail struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e APIError) Error() string {
	if len(e.Errors) > 0 {
		err := e.Errors[0]
		return fmt.Sprintf("clearbit: %d %v", err.Type, err.Message)
	}
	return ""
}

// Empty returns true if empty. Otherwise, at least 1 error message/code is
// present and false is returned.
func (e APIError) Empty() bool {
	if len(e.Errors) == 0 {
		return true
	}
	return false
}

// relevantError returns any non-nil http-related error (creating the request,
// getting the response, decoding) if any. If the decoded apiError is non-zero
// the apiError is returned. Otherwise, no errors occurred, returns nil.
func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}
	if apiError.Empty() {
		return nil
	}
	return apiError
}

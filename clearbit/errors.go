package clearbit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// apiError represents a Clearbit API Error response
// https://clearbit.com/docs#errors
type apiError struct {
	HTTPStatus string        `json:"httpStatus"`
	Errors     []ErrorDetail `json:"error"`
}

type ErrorDetailWrapper struct {
	Err ErrorDetail `json:"error"`
}

// ErrorDetail represents an individual item in an apiError.
type ErrorDetail struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// ErrorDetail represents an individual item in an apiError.
func (e apiError) Error() string {
	msg := e.HTTPStatus
	if len(e.Errors) > 0 {
		err := e.Errors[0]
		if err.Type != "" || err.Message != "" {
			msg += fmt.Sprintf(" type=%s message=%v", err.Type, err.Message)
		}
	}
	if msg != "" {
		msg = "clearbit: " + msg
	}
	return msg
}

// UnmarshalJSON is used to be able to read dynamic json
//
// This is because sometimes our errors are not arrays of ErrorDetail but a
// single ErrorDetail
func (e *apiError) UnmarshalJSON(b []byte) error {
	errorWrapper := ErrorDetailWrapper{}
	errors := []ErrorDetail{}

	if tmpErr := json.Unmarshal(b, &errorWrapper); tmpErr == nil { // && errorDetail.Type != ""
		errors = append(errors, errorWrapper.Err)
		e.Errors = errors
		return nil
	}

	if tmpErr := json.Unmarshal(b, &errors); tmpErr == nil {
		e.Errors = errors
		return nil
	}

	return nil
}

// Empty returns true if empty. Otherwise, at least 1 error message/code is
// present and false is returned.
func (e *apiError) Empty() bool {
	if len(e.Errors) == 0 {
		return true
	}
	return false
}

// relevantError returns any non-nil http-related error (creating the request,
// getting the response, decoding) if any. If the decoded apiError is non-zero
// the apiError is returned. Otherwise, no errors occurred, returns nil.
func relevantError(resp *http.Response, httpError error, ae apiError) error {
	if httpError != nil {
		return httpError
	}
	if ae.Empty() {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		ae.HTTPStatus = resp.Status
	}
	return ae
}

package types

// Error represents an API error with optional code and message.
type Error struct {
	Code    int    `json:"code,omitempty"`    // Error code
	Message string `json:"message,omitempty"` // Description of the error
}

// NewError creates a new Error with the given code and message.
func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// Response standardizes API response structures.
type Response struct {
	Success bool        `json:"success"`          // Success status of the operation
	Error   *Error      `json:"error,omitempty"`  // Details of any error that occurred
	Result  interface{} `json:"result,omitempty"` // Result data of the operation
}

// NewResponseError returns a Response indicating a failure with the specified error details.
func NewResponseError(code int, message string) *Response {
	return &Response{
		Success: false,
		Error:   NewError(code, message),
	}
}

// NewResponseResult returns a Response indicating a success with the provided result data.
func NewResponseResult(v interface{}) *Response {
	return &Response{
		Success: true,
		Result:  v,
	}
}

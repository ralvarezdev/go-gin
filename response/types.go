package response

type (
	// Response struct
	Response struct {
		data interface{}
		code *int
	}

	// JSONErrorResponse struct
	JSONErrorResponse struct {
		Error string `json:"error"`
	}
)

// NewResponseWithCode creates a new response with a code
//
// Parameters:
//
//   - data: the data to be sent in the response
//   - code: the HTTP status code to be sent in the response
//
// Returns:
//
//   - *Response: a pointer to the Response struct
func NewResponseWithCode(data interface{}, code int) *Response {
	return &Response{data: data, code: &code}
}

// NewErrorResponseWithCode creates a new error response with a code
//
// Parameters:
//
//   - err: the error to be sent in the response
//   - code: the HTTP status code to be sent in the response
//
// Returns:
//
//   - *Response: a pointer to the Response struct
func NewErrorResponseWithCode(err error, code int) *Response {
	return &Response{data: NewJSONErrorResponse(err), code: &code}
}

// NewResponse creates a new response
//
// Parameters:
//
//   - data: the data to be sent in the response
//
// Returns:
//
//   - *Response: a pointer to the Response struct
func NewResponse(data interface{}) *Response {
	return &Response{data: data}
}

// Code returns the HTTP status code
//
// Returns:
//
//   - *int: a pointer to the HTTP status code (nil if not set)
func (r Response) Code() *int {
	return r.code
}

// Data returns the response data
//
// Returns:
//
//   - interface{}: the response data
func (r Response) Data() interface{} {
	return r.data
}

// NewErrorResponse creates a new error response
//
// Parameters:
//
//   - err: the error to be sent in the response
//
// Returns:
//
//   - *Response: a pointer to the Response struct
func NewErrorResponse(err error) *Response {
	return &Response{data: NewJSONErrorResponse(err)}
}

// NewJSONErrorResponse creates a new error response
//
// Parameters:
//
//   - err: the error to be sent in the response
//
// Returns:
//
//   - JSONErrorResponse: the JSONErrorResponse struct
func NewJSONErrorResponse(err error) JSONErrorResponse {
	return JSONErrorResponse{Error: err.Error()}
}

// NewJSONErrorResponseFromString creates a new error response
//
// Parameters:
//
//   - err: the error message to be sent in the response
//
// Returns:
//
//   - JSONErrorResponse: the JSONErrorResponse struct
func NewJSONErrorResponseFromString(err string) JSONErrorResponse {
	return JSONErrorResponse{Error: err}
}

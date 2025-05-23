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
func NewResponseWithCode(data interface{}, code int) *Response {
	return &Response{data: data, code: &code}
}

// NewErrorResponseWithCode creates a new error response with a code
func NewErrorResponseWithCode(err error, code int) *Response {
	return &Response{data: NewJSONErrorResponse(err), code: &code}
}

// NewResponse creates a new response
func NewResponse(data interface{}) *Response {
	return &Response{data: data}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(err error) *Response {
	return &Response{data: NewJSONErrorResponse(err)}
}

// NewJSONErrorResponse creates a new error response
func NewJSONErrorResponse(err error) JSONErrorResponse {
	return JSONErrorResponse{Error: err.Error()}
}

// NewJSONErrorResponseFromString creates a new error response
func NewJSONErrorResponseFromString(err string) JSONErrorResponse {
	return JSONErrorResponse{Error: err}
}

package response

// ResponseRequest wrap data to store data or error
type ResponseRequest struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
}

// ResponseRequest extract error to struct
func (r ResponseRequest) ParseError() ErrorRequestResponse {
	err := r.Error.(ErrorRequestResponse)

	// set message error
	err.error()
	return err
}

// DataRequestResponse wrap response data
type DataRequestResponse struct {
	Code   int         `json:"code,omitempty"`
	Status int         `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Paging interface{} `json:"paging,omitempty"`
}

// ErrorRequestResponse wrap error response
type ErrorRequestResponse struct {
	Status   int    `json:"status"`
	LogLevel int    `json:"-"`
	Code     int    `json:"-"`
	Param    string `json:"-"`
	Message  string `json:"message"`
	Error    error  `json:"-"`
	Event    string `json:"-"`
}

// error define below:
// invalidDatabse - error database:
// invalidData - error:
func (e *ErrorRequestResponse) error() {}

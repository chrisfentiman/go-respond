package respond

import "net/http"

// SuccessResponse is a standard HTTP JSON response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

// Ok returns a 200 Ok response
func (resp *HTTPResponse) Ok(v interface{}) {
	resp.WriteSuccessResponse(http.StatusOK, v)
}

// Created returns a 201 Created response
func (resp *HTTPResponse) Created(v interface{}) {
	resp.WriteSuccessResponse(http.StatusCreated, v)
}

// WriteSuccessResponse is the error response writer
func (resp *HTTPResponse) WriteSuccessResponse(code int, v interface{}) {
	resp.SetStatusCode(code)
	body := &SuccessResponse{true, v}
	resp.SetBody(resp.MarshallJSON(body)).
		WriteResponse()
}
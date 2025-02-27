// Copyright 2018 Adam S Levy. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package jsonrpc2

import "fmt"

// Error represents a JSON-RPC 2.0 Error object, which is used in the Response
// object.
type Error struct {
	Code    ErrorCode   `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewError returns an Error with the given code, message, and data.
func NewError(code ErrorCode, message string, data interface{}) *Error {
	return &Error{Code: code, Message: message, Data: data}
}

// Error implements the error interface.
func (e Error) Error() string {
	s := fmt.Sprintf("jsonrpc2.Error{Code:%v, Message:%#v", e.Code, e.Message)
	if e.Data != nil {
		s += fmt.Sprintf(", Data:%#v", e.Data)
	}
	return s + "}"
}

// Official JSON-RPC 2.0 Errors

// InvalidParams returns a pointer to a new Error initialized with the
// InvalidParamsCode and InvalidParamsMessage and the user provided data.
// MethodFuncs are responsible for detecting and returning this error.
func InvalidParams(data interface{}) *Error {
	return NewError(InvalidParamsCode, InvalidParamsMessage, data)
}

// internalError returns a pointer to a new Error initialized with the
// InternalErrorCode and InternalErrorMessage and the user provided data.
func internalError(data interface{}) *Error {
	return NewError(InternalErrorCode, InternalErrorMessage, data)
}

// parseError returns a pointer to a new Error initialized with the
// ParseErrorCode and ParseErrorMessage and the user provided data.
func parseError(data interface{}) *Error {
	return NewError(ParseErrorCode, ParseErrorMessage, data)
}

// invalidRequest returns a pointer to a new Error initialized with the
// InvalidRequestCode and InvalidRequestMessage and the user provided data.
func invalidRequest(data interface{}) *Error {
	return NewError(InvalidRequestCode, InvalidRequestMessage, data)
}

// methodNotFound returns a pointer to a new Error initialized with the
// MethodNotFoundCode and MethodNotFoundMessage and the user provided data.
func methodNotFound(data interface{}) *Error {
	return NewError(MethodNotFoundCode, MethodNotFoundMessage, data)
}

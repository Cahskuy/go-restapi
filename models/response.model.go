package models

import "encoding/json"

type ErrorInputResponse struct {
	FieldName string `json:"fieldName"`
	Message   string `json:"message"`
}

type WebResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type WebErrorInputResponse struct {
	Code       int             `json:"code"`
	Status     string          `json:"status"`
	ErrorField json.RawMessage `json:"errorField"`
}

type ServiceResponse struct {
	Message string `json:"message"`
}

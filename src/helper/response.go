package helpers

import (
	"fmt"
	"reflect"
)

// APIResponse ...
type APIResponse struct {
	// Message string 			`default:"nil" json:"message" xml:"message"`
	Meta interface{} `default:"nil" json:"meta" xml:"meta"`
	Data interface{} `default:"nil" json:"data" xml:"data"`
}

// ErrorResponseWithPayload ...
type ErrorResponseWithPayload struct {
	Payload       interface{} `json:"payload" xml:"payload"`
	StatusCode    int         `json:"status_code" xml:"status_code"`
	StatusMessage string      `json:"status_message" xml:"status_message"`
}

// Run ...
func Run() {
	fmt.Println("helpers")
}

// OutputErrorResponseWithPayload ...
func OutputErrorResponseWithPayload(params map[string]interface{}) interface{} {
	var payload interface{}
	statusMessage := "error validation"
	code := 422
	for key, val := range params {
		if key == "payload" {
			payload = val
		}
		if key == "code" {
			code = val.(int)
		}
		if key == "statusMessage" {
			statusMessage = val.(string)
		}
	}

	return &ErrorResponseWithPayload{
		Payload:       payload,
		StatusCode:    code,
		StatusMessage: statusMessage,
	}
}

// OutputAPIResponseWithPayload ...
func OutputAPIResponseWithPayload(params map[string]interface{}) interface{} {
	var payload interface{}
	var meta interface{}
	for key, val := range params {
		if key == "payload" {
			payload = val
		}
		if key == "meta" {
			if reflect.TypeOf(val).Kind() == reflect.String {
				meta = map[string]interface{}{"message": val}
			} else {
				meta = val
			}
		}
	}

	return &APIResponse{Data: payload, Meta: meta}
}

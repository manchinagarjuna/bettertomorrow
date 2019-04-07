package util

import "encoding/json"

type RequestMessage struct {
	Operation string          `json:operation`
	Data      json.RawMessage `json:data`
}

type ResponseMessage struct {
	Data  interface{} `json:data`
	Error string      `json:error`
}

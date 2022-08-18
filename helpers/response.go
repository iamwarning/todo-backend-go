package helpers

import (
	"encoding/json"
	"net/http"
	"reflect"
)

func Response(w http.ResponseWriter, statusCode int, payload interface{}) {
	content := contentType(payload)
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", content)
	w.WriteHeader(statusCode)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}

func contentType(payload interface{}) string {
	content := reflect.ValueOf(payload).Kind().String()
	if content == "string" {
		return "text/plain"
	}
	return "application/json"
}

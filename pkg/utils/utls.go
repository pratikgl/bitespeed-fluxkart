package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody is a helper function which is intended for
// parsing the JSON request body of an HTTP request and
// unmarshaling it into a Go struct or interface.
func ParseRequestBody(r *http.Request, x interface{}) {
	// Read the body of the request.
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

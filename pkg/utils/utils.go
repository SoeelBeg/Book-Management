package utils

import (
	"encoding/json"
	"net/http"
)

// ParseBody parses the body of the request into the provided struct
func ParseBody(r *http.Request, x interface{}) error {
	return json.NewDecoder(r.Body).Decode(x)
}

// package utils

// import (
// 	"encoding/json"
// 	"io"
// 	"net/http"
// 	"log"
// )

// // ParseBody reads the body from the HTTP request and unmarshals it into the provided interface.
// func ParseBody(r *http.Request, x interface{}) {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		log.Printf("Error reading body: %v", err)
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := json.Unmarshal(body, x); err != nil {
// 		log.Printf("Error unmarshalling JSON: %v", err)
// 	}
// }

// listing create lawyers profiles and sharing
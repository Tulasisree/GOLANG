//thsi file handles all rough routing
package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	//in go /users , /users/ is different
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJson(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

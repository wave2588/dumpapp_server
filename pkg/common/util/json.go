package util

import (
	"encoding/json"
	"net/http"
)

// ArgsValidator help you to validate the incoming data.
type ArgsValidator interface {
	Validate() error
}

// JSONArgs fill the out using request Body explained in json format.
func JSONArgs(r *http.Request, out ArgsValidator) error {
	if err := json.NewDecoder(r.Body).Decode(out); err != nil {
		return err
	}

	if err := out.Validate(); err != nil {
		return err
	}

	return nil
}

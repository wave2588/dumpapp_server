package formatter

import (
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/errors"
)

func RenderError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	v, ok := err.(*errors.APIError)
	if ok {
		w.WriteHeader(v.HttpStatus())
	}

	util.RenderJSON(w, err)
}

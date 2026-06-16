package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

)


func WriteJSONResponse(w http.ResponseWriter, status int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	e := json.NewEncoder(w)
	e.SetIndent("", "    ")
	e.Encode(response)

}

func ReadBody(r *http.Request) []byte {

	body := r.Body
	defer body.Close()

	var b bytes.Buffer
	if _, err := io.Copy(&b, body); err != nil {
		// handle err
	}
	data := b.Bytes()
	body.Close()

	return data
}

func WriteISError(w http.ResponseWriter, estr string) {
		WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"error": estr,
		})

}

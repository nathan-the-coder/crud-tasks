package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/devjefster/GoShortUniqueID/idgen"
)

func IDGen() string {
	idGen := idgen.New(6, "", "")
	id := ""

	for range 5 {
		id = idGen.Generate()
	}
	return id
}

func WriteJSONResponse(w http.ResponseWriter, status int, response any) {
	w.Header().Set("Content-Type", "application/json")
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

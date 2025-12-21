package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/rwwiv/rwwiv.com/transform"
	"github.com/rwwiv/rwwiv.com/util"
)

type TransformSpec struct {
	Type   string         `json:"type"`
	Params map[string]any `json:"-"`
}

func (t *TransformSpec) UnmarshalJSON(data []byte) error {
	// First unmarshal to get the type
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Extract type
	if typeVal, ok := raw["type"].(string); ok {
		t.Type = typeVal
	}

	// Copy all other fields to Params
	t.Params = make(map[string]any)
	for k, v := range raw {
		if k != "type" {
			t.Params[k] = v
		}
	}

	return nil
}

func ChainHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20) // max 32MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	contentType := r.PostFormValue("content_type")

	img, err := util.Decode(file, contentType)
	if err != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// Parse transforms JSON array
	transformsJSON := r.PostFormValue("transforms")
	if transformsJSON == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var specs []TransformSpec
	if err := json.Unmarshal([]byte(transformsJSON), &specs); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Apply each transform in order
	currentImg := img
	for _, spec := range specs {
		constructor, ok := transform.Registry[spec.Type]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t, err := constructor(spec.Params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		currentImg, err = t.Apply(currentImg)
		if err != nil {
			panic(err)
		}
	}

	outputBuf := new(bytes.Buffer)
	err = util.Encode(outputBuf, currentImg, contentType)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(outputBuf.Bytes())
}

package api

import (
	"bytes"
	"net/http"

	"github.com/rwwiv/rwwiv.com/transform"
	"github.com/rwwiv/rwwiv.com/util"
)

func InvertHandler(w http.ResponseWriter, r *http.Request) {
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

	invert, err := transform.NewInvert(nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	processedImg, err := invert.Apply(img)
	if err != nil {
		panic(err)
	}

	outputBuf := new(bytes.Buffer)
	err = util.Encode(outputBuf, processedImg, contentType)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(outputBuf.Bytes())
}

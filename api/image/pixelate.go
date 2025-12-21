package api

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/rwwiv/rwwiv.com/transform"
	"github.com/rwwiv/rwwiv.com/util"
)

func PixelateHandler(w http.ResponseWriter, r *http.Request) {
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

	pixelSize, err := strconv.Atoi(r.PostFormValue("pixel_size"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pixelate, err := transform.NewPixelate(map[string]any{"pixel_size": pixelSize})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	processedImg, err := pixelate.Apply(img)
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

package api

import (
	"encoding/json"
	"net/http"
)

type Foo struct {
	Bar string `json:"bar"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	foo := Foo{Bar: "Baz"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(foo)
	w.Write(json)
}

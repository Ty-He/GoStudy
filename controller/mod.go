package controller 

import (
    "net/http"
)

func registerModHandle() {
    http.HandleFunc("/mod", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("handle mod"))
    })
}

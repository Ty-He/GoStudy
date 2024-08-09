package controller 

import (
    "net/http"
)

func registerDelHandle() {
    http.HandleFunc("/del", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("handle del"))
    })
}

package controller 

import (
    "net/http"
)

func registerAddHandle() {
    http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("handle add"))
    })
}

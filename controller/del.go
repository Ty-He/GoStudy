package controller 

import (
    "log"
    "net/http"

    "github.com/ty/crud_web/model"
)

func registerDelHandle() {
    http.HandleFunc("/del", func(w http.ResponseWriter, r *http.Request) {
        // if r.Method != http.MethodDelete {
        //     http.NotFound(w, r)
        //     return
        // }
        log.Println("handle delete method")
        // CROS
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        // handle preflight request
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        id := r.URL.Query().Get("id")
        // delete record by id 
        err := model.DeleteUser(id)
        if err != nil {
            log.Println("DeleteUser: ", err)
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        w.WriteHeader(http.StatusOK)
    })
}

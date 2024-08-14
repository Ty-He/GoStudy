package controller 

import (
    "log"
    "net/http"
    "strconv"

    "github.com/ty/crud_web/model"
)

func registerDelHandle() {
    http.HandleFunc("/del", func(w http.ResponseWriter, r *http.Request) {
        // if r.Method != http.MethodDelete {
        //     http.NotFound(w, r)
        //     return
        // }

        id_str := r.URL.Query().Get("id")
        id, err := strconv.Atoi(id_str)
        if err != nil {
            log.Println("Atoi err:", err)
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        // delete record by id 
        err = model.DeleteUser(uint32(id))
        if err != nil {
            log.Println("DeleteUser: ", err)
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        w.WriteHeader(http.StatusOK)
    })
}

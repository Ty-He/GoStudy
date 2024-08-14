package controller 

import (
    "log"
    "net/http"
    "encoding/json"

    "github.com/ty/crud_web/model"
)

func registerModHandle() {
    http.HandleFunc("/mod", func(w http.ResponseWriter, r *http.Request) {
        if (r.Method != http.MethodPatch) {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        // parse json data
        decoder := json.NewDecoder(r.Body)
        u := &model.UserInformation{}

        // Decode
        if err:= decoder.Decode(u); err != nil {
            log.Panicln("decoder.Decode() err:", err)
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        // updata db
        if err := u.Update(); err != nil {
            log.Println("Updata error:", err)
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        w.WriteHeader(http.StatusOK)
    })
}

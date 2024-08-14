package controller 

import (
    "log"
    "net/http"
    
    "github.com/ty/crud_web/model"
    "github.com/ty/crud_web/view"
)

func registerHomeHandle() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // first, get home template
        // then, find all user data
        users, err := model.GetTotalUsers()
        if err != nil {
            log.Println("GetTotalUsers: ", err)
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        // load html
        tmplData := &view.TmplData{
            Type: "home",
            Value: users,
        }
        if err := view.ExecuteTemplate(w, tmplData); err != nil {
            log.Println("ExecuteTemplate: ", err)
            w.WriteHeader(http.StatusBadRequest)
        }
    })
}

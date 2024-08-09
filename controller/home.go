package controller 

import (
    "net/http"
    "log"
    
    "github.com/ty/crud_web/model"
    "github.com/ty/crud_web/view"
)

func registerHomeHandle() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // first, get home template
        // then, find all user data
        users := model.GetTotalUsers()
        // load html
        err := view.HomeTemplate.Execute(w, users)
        if err != nil {
            log.Println(err)
        }
    })
}

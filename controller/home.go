package controller 

import (
    "log"
    "net/http"
    
    "github.com/ty/crud_web/model"
    "github.com/ty/crud_web/view"
)

func registerHomeHandle() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("handle home")
        // first, get home template
        // then, find all user data
        users := model.GetTotalUsers()
        // load html
        tmplData := &view.TmplData{
            Type: "home",
            Value: users,
        }
        if err := view.ExecuteTemplate(w, tmplData); err != nil {
            log.Println(err)
        }
    })
}

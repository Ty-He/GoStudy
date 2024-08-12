package controller 

import (
    "log"
    "net/http"

    "github.com/ty/crud_web/model"
    "github.com/ty/crud_web/view"
)


// get operator web page
func registerAddPageHandle() {
    http.HandleFunc("/add_operator_page", func(w http.ResponseWriter, r *http.Request) {
        if err := view.ExecuteTemplate(w, &view.TmplData{
            Type: "add_user",
        }); err != nil {
            log.Println(err)
        }
    })
}

// handle add Request
func registerAddHandle() {
    http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
        err := r.ParseForm()
        if err != nil {
            log.Println("r.ParseForm() err", err)
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        u := model.NewUser(
            r.PostFormValue("name"),
            r.PostFormValue("gender"),
            r.PostFormValue("introduction"),
            r.PostFormValue("password"),
        )
        err = u.Add()
        var tmplStatusData *view.TmplData
        if err != nil {
            log.Println("user add db err: ", err)
            tmplStatusData = view.NewStatusTmplData("occur error!", err.Error())
        } else {
            tmplStatusData = view.NewStatusTmplData(
                "operate ok!",
                u.String(),
            )
        }
        err = view.ExecuteTemplate(w, tmplStatusData)
        if err != nil {
            log.Println(err)
        }
    })
}

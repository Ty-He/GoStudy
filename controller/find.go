package controller 

import (
    "log"
    "net/http"
    "strconv"

    "github.com/ty/crud_web/model"
    "github.com/ty/crud_web/view"
)

// get find page
func registerFindPageHandle() {
    http.HandleFunc("/find_page", func(w http.ResponseWriter, r *http.Request) {
        if err := view.ExecuteTemplate(w, &view.TmplData{
            Type: "find_user",
        }); err != nil {
            log.Println(err)
        }
    })
}

func registerFindHandle() {
    http.HandleFunc("/find", findHandle)
}

func findHandle(w http.ResponseWriter, r * http.Request) {
    if (r.Method != http.MethodGet) {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    err := r.ParseForm()
    if err != nil {
        log.Println("r.ParseForm() err", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    selectedOption := r.FormValue("option")
    log.Println("selectedOption =>", selectedOption)

    switch selectedOption {
    case "id":
        findByIdHandle(w, r)
    case "name":
        findByNameHandle(w, r)
    default:
        w.WriteHeader(http.StatusBadRequest)
    }
}

func findByIdHandle(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.FormValue("id"))
    if err != nil {
        log.Println("Atoi err:", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    u, err := model.GetUserById(uint32(id))
    if err != nil {
        log.Println("GetUserById err:", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    tmplData := &view.TmplData{
        Type: "home",
        Value: u,
    }
    if err := view.ExecuteTemplate(w, tmplData); err != nil {
        log.Println("ExecuteTemplate: ", err)
        w.WriteHeader(http.StatusBadRequest)
    }

}

func findByNameHandle(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    users, err := model.GetUsersByName(name)
    if err != nil {
        log.Println("GetUsersByName: ", err)
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
}

package controller 

import "net/http"

func RegisterHandler() {
    registerHomeHandle()
    registerAddPageHandle()
    registerAddHandle()
    registerDelHandle()
    registerModHandle()
    http.Handle("/css/", http.FileServer(http.Dir("templates")))
    http.Handle("/js/", http.FileServer(http.Dir("templates")))
}

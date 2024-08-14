package controller 

import "net/http"

func RegisterHandler() {
    // home page
    registerHomeHandle()

    // add user
    registerAddPageHandle()
    registerAddHandle()

    // updata or delete
    registerDelHandle()
    registerModHandle()
    
    // find
    registerFindPageHandle()
    registerFindHandle()
    http.Handle("/css/", http.FileServer(http.Dir("templates")))
    http.Handle("/js/", http.FileServer(http.Dir("templates")))
}

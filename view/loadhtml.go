package view 

import (
    "log"
    "text/template"
)

var HomeTemplate *template.Template

func init() {
    HomeTemplate = getHomeTemplate()
    log.Println("init template finnish")
}

func getHomeTemplate() *template.Template {
    t := template.New("home.html")
    t = template.Must(t.ParseFiles("./templates/home.html"))
    return t
} 

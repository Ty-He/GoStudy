package view 

import (
    "log"
    "io"
    "text/template"
)

var layoutTemplateObj *template.Template

func init() {
    layoutTemplateObj = getFileTemplate()
    log.Println("init template finnish")
}

type TmplData struct {
    // the name of template
    Type string 

    // params of tempalte
    Value any
} 

// create a status result
func NewStatusTmplData(status, content string) *TmplData {
    return &TmplData{
        Type: "response_status",
        Value: map[string]any {
            "Status": status,
            "Content": content,
        },
    }
}

func ExecuteTemplate(w io.Writer, tmplData *TmplData) error {
    return layoutTemplateObj.ExecuteTemplate(w, "layout", tmplData)
}

func getFileTemplate() *template.Template {
    t := template.New("layout")
    // t = template.Must(t.ParseFiles(
    //     "./templates/layout.html",
    //     "./templates/home.html",
    //     "./templates/add_user.html"))
    t, err := t.ParseGlob("./templates/*.html")
    if err != nil {
        panic(err)
    }
    return t
} 



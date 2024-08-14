package main 

import (
    "net/http"
    "fmt"

    "github.com/ty/crud_web/controller"
    "github.com/ty/crud_web/middleware"
)

func main() {
    controller.RegisterHandler()
    s := http.Server{
        Addr: "192.168.18.128:8888",
        Handler: &middleware.CrosMiddleawre{},
    }

    fmt.Println("http server start...")
    s.ListenAndServe()
}

package main 

import (
    "net/http"
    "fmt"

    "github.com/ty/crud_web/controller"
)

func main() {
    controller.RegisterHandler()

    fmt.Println("http server start...")
    http.ListenAndServe("192.168.18.128:8888", nil)
}

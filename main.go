package main

import "golang-im-system/server"

func main() {
    s := server.NewServer("192.168.18.128", 9999)
    s.Run()
}

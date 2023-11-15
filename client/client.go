package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type client struct {
    dest_ip string
    dest_port int
    name string
    conn net.Conn
    mode int;
}

func NewClient(ip string, port int) *client{
    // get client obj
    clt := &client{
        dest_ip: ip,
        dest_port: port,
        mode: -1,
    }
    // connect server
    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
    if err != nil {
        fmt.Println("Dial Error: ", err)
        return nil
    }
    clt.conn = conn
    return clt
}

func (this *client) menu() bool {
    fmt.Println("1. Public Mode")
    fmt.Println("2. Private Mode")
    fmt.Println("3. Rename")
    fmt.Println("0. Exit")
    var m int
    fmt.Scanln(&m)
    if m < 0 || m > 3 {
        fmt.Println("Input Error.")
        return false
    }
    this.mode = m
    return true
}

func (this *client) undateName() bool {
    fmt.Println("Please input your new name: ")
    fmt.Scanln(&this.name)
    sendMsg := "rename|" + this.name + "\n";
    _ ,err := this.conn.Write([]byte(sendMsg))
    
    if err != nil {
        fmt.Println("Conn Write Error: ", err)
        return false
    }
    return true
}

// recv server's message
func (this *client) recvMsg() {
    // if the stdout have any buffer, this fn will print
    // At the same time, this fu is clog
    io.Copy(os.Stdout, this.conn)

    // the line is same with:
    /*
    for {
        buf := make([]byte, 256)
        _, err := this.conn.Read(buf)
        if err != nil {
            fmt.Println("Read Error, ", err)
            continue
        }
        fmt.Println(buf)
    }
    */
}

// send a message to server
func (this *client) sendMsg(msg string) bool {
    // this '\n' is for agreement
    msg += "\n"
    _, err := this.conn.Write([]byte(msg))
    if err != nil {
        fmt.Println("Write Error,", err)
        return false
    }
    return true
}

func (this *client) publicChat() {
    for {
        fmt.Println("Send Msg(public), input 'q' for exit")
        var chat string
        fmt.Scanln(&chat)
        if len(chat) != 0 {
            if chat == "q" {
                break
            }
            flag := this.sendMsg(chat)
            if !flag {
                break
            }
        }
    }
}

func (this *client) privateChat() {
    for {
        var remoteName string
        // show all user which online
        this.sendMsg("who")
        fmt.Println("Please select a user name ('q' for quit)")
        fmt.Scanln(&remoteName)
        if remoteName == "q" {
            break
        }
        var content string
        for content != "q"{
            fmt.Println("Please input you content ('q' for select)")
            fmt.Scanln(&content)
            // send
            flag := this.sendMsg("to|" + remoteName + "|" + content + "\n")
            if !flag {
                break
            }
        }
    }
}

func (this *client) run() {
    for this.mode != 0 {
        for !this.menu() {}
        switch this.mode {
        case 1:
            // fmt.Println("Public Mode")
            this.publicChat()
            break
        case 2:
            // fmt.Println("Private Mode")
            this.privateChat()
            break
        case 3:
            // fmt.Println("Rename")
            this.undateName()
            break
        case 0:
            break
        }
    }
}

// 
var serverIp string
var serverPort int

func init() {
    flag.StringVar(&serverIp, "ip", "192.168.18.128", "Set server's ip.")
    flag.IntVar(&serverPort, "port", 9999, "Set server's port.")
}

func main() {
    // Command line parsing
    flag.Parse()
    clt := NewClient("192.168.18.128", 9999)
    if clt == nil {
        fmt.Println("Error")
    }
    fmt.Println("Conn successfully!")

    // goroutine, for reacing server Message
    go clt.recvMsg()

    clt.run()
}

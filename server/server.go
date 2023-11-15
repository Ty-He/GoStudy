package server

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
    Ip string
    Port int
    OnlineMap map[string]*User
    MapLock sync.RWMutex
    Message chan string
}

func NewServer(ip string, port int) *Server {
    server := &Server{
        Ip: ip,             
        Port: port,
        OnlineMap: make(map[string]*User),
        Message: make(chan string),
        }
    
    return server
}

// Listen Broadcast, if it is not empty, send all client
func (this *Server) ListenMsg() {
    // this goroutine should be working forever bug1
    for {
        msg := <-this.Message
        this.MapLock.Lock()
        for _,clt := range this.OnlineMap {
            clt.C <- msg
        }
        this.MapLock.Unlock()
    }
}

// send broadcast
func (this *Server) SendMessage(user *User, msg string) {
    sendMsg := user.Addr + " [" + user.Name + "] " + msg + " ";
    this.Message <- sendMsg
}

// after Accept a conn, this function is for communicate
func (this *Server) Handler(conn net.Conn) {
    fmt.Println("Accept a conn.")

    // save User information
    user := NewUser(conn, this)
    /*
    this.MapLock.Lock()
    this.OnlineMap[user.Name] = user
    this.MapLock.Unlock()

    // send broadast
    this.SendMessage(user, "is online.")
    */
    user.Online()

    active := make(chan bool)

    // read cfd
    go func() {
        buf := make([]byte, 1024)
        for {
            n, err := conn.Read(buf)
            if n == 0 {
                // this.SendMessage(user, "is offline.")
                user.Offline()
                return
            } 
            if err != nil && err != io.EOF {
                fmt.Println("Read error : ", err)
                return
            }

            // the n - 1 is for removing \n
            msg := string(buf[:n-1])
            // this.SendMessage(user, msg)
            user.DoMessage(msg)

            // if DoMessage, this user is active
            active <- true
        }
    } ()
    // 
    for {
        select {
        case <- active:
            // do nothing
        case <- time.After(300 * time.Second):
            user.SendMessage("You are not active!\n")
            // destroy memory
            close(user.C)
            close(active)
            conn.Close()
            // exit current goroutine
            return
        }
    }
}

func (this *Server) Run() {
    // sock-fd init set listen
    listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
    if err != nil {
        fmt.Println("listen", err)
        return
    }
    defer listener.Close()
    
    // run ListenMsg
    go this.ListenMsg()
    
    // accept conn
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Accept", err)
            continue
        }

        // new a goroutine to communicate
        go this.Handler(conn)

    }
}

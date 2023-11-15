package server

import (
	"net"
	"strings"
)

type User struct {
    Name string
    Addr string
    C chan string
    conn net.Conn
    server *Server
}

// get a user objection
func NewUser(c net.Conn, s *Server) *User {
    userAddr := c.RemoteAddr().String()
    user := &User{
        Name: userAddr,
        Addr: userAddr,
        C: make(chan string),
        conn: c,
        server: s,
    }

    // run ListenMsg goroutine
    go user.ListenMsg()

    return user
}

func (this *User) Online() {
    this.server.MapLock.Lock()
    this.server.OnlineMap[this.Name] = this
    this.server.MapLock.Unlock()

    this.server.SendMessage(this, "is online")
}

func (this *User) Offline() {
    this.server.SendMessage(this, "is offline")
    
    // remove this from hashtabl
    this.server.MapLock.Lock()
    delete(this.server.OnlineMap, this.Name)
    this.server.MapLock.Unlock()
}

// this func is send to only one, differ with broadcast
func (this *User) SendMessage(msg string) {
    this.conn.Write([]byte(msg))
}

func (this *User) DoMessage(msg string) {
    if msg == "who" {
        this.server.MapLock.Lock()
        // get all user information
        for _, user := range this.server.OnlineMap {
            info := user.Addr + " [" + user.Name + "] is online\n" 
            this.SendMessage(info)
        }
        this.server.MapLock.Unlock()
    } else if len(msg) > 7 && msg[:7] == "rename|" {
        newName := msg[7:]
        // strings.Split(msg, "|")[1]
        this.server.MapLock.Lock()
        // only care this key is exist or not
        _, exist := this.server.OnlineMap[newName]
        if exist {
            this.SendMessage("This name is already exist!\n")
        } else {
            delete(this.server.OnlineMap, this.Name)
            this.Name = newName
            this.server.OnlineMap[this.Name] = this
            this.SendMessage("Rename successfully!\n")
        }
        this.server.MapLock.Unlock()
    } else if len(msg) > 4 && msg[:3] == "to|"{
        // get the person which recv message 's name
        substr := strings.Split(msg, "|")
        toName := substr[1]
        if toName == "" {
            this.SendMessage("Format error, for example, to|ty|hello.\n")
            return
        }
        // get User
        destUser, exist := this.server.OnlineMap[toName]
        if !exist {
            this.SendMessage("This person is not exist!\n")
            return
        }
        // get content
        content := substr[2]
        if content == "" {
            this.SendMessage("Error:Send empty content.\n")
            return
        }
        destUser.SendMessage(this.Name + ":(private)" + content)
    } else {
        this.server.SendMessage(this, msg)
    }

}

// listen user's channel
func (this *User) ListenMsg() {
    for {
        msg := <-this.C
        
        // Write to client
        this.conn.Write([]byte(msg + "\n"))
    }
}

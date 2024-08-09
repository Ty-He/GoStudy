package model 

import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
const dsn string = "root:245869@tcp(192.168.10.106:3306)/ty_data"

func init() {
    var err error
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalln("sql.Open() error", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatalln("db.Ping() error", err)
    }
    
    log.Println("init db finish")
}

type UserInformation struct {
    Uid uint32 `json:"uid"`
    Name string `json:"name"`
    Gender string `json:"gender"`
    Introduction string `json:"introduction"`
    Password string `json:"password"`
}

func GetTotalUsers() (users []*UserInformation) {
    query_sql := "select * from user_informations;"
    rows, err := db.Query(query_sql)
    if err != nil {
        log.Println("db.Query() error", err)
        return nil
    }
    defer rows.Close()

    for rows.Next() {
        u := &UserInformation{}
        err := rows.Scan(&u.Uid, &u.Name, &u.Gender, &u.Introduction, &u.Password)
        if err != nil {
            log.Println("rows.Scan() error", err)
            return nil
        }
        users = append(users, u)
    }

    return
}

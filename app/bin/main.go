package main

import (

  "encoding/json"

  "net/http"

  "fmt"

  "github.com/gorilla/mux"

  "github.com/jinzhu/gorm"

  "github.com/rs/cors"

  _ "github.com/jinzhu/gorm/dialects/mssql"
)

type Pop struct {

  //Id     uint        `gorm:"primary_key"`
  gorm.Model

  Name   string

  Color  string
}

func main() {
  router := mux.NewRouter()

  router.HandleFunc("/", GetPops).Methods("GET")

  handler := cors.Default().Handler(router)

  http.ListenAndServe(":8080", handler)
}

func GetPops(w http.ResponseWriter, r *http.Request) {

  var db *gorm.DB

  const (
    server     = "db"
    port     = 1433
    user     = "sa"
    password = "z!oBx1ab"
    database   = "beverage"
  )

  sqlserverInfo := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)

  db, err := gorm.Open("mssql", sqlserverInfo)

  if err != nil {
    panic(err)
  }

  if db != nil {
gorm.DefaultCallback.Create().Remove("mssql:set_identity_insert")
    db.Debug().DropTableIfExists(&Pop{})
    //Drops table if already exists
    var items = []Pop{

            {Name: "Cola", Color: "carmel"},

            {Name: "Gingerale",  Color: "brown"},

            {Name: "Lime",  Color: "green"},

            {Name: "Cherry",  Color: "red"},

            {Name: "Grape",  Color: "purple"},
        }
    db.Debug().AutoMigrate(&Pop{})
    //Auto create table based on Model
    for index := range items {

          db.Create(&items[index])

      }

    var pops []Pop

    db.Find(&pops)

    defer db.Close()

    json.NewEncoder(w).Encode(&pops)
  } else {
    fmt.Fprintf(w, "Hello astaxie!")
  }
}

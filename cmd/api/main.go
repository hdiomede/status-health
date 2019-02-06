package main

import (
    "os"
    "fmt"
    "github.com/hdiomede/status-health/server"
    "github.com/hdiomede/status-health/mongo"
    "github.com/hdiomede/status-health/model"
	"gopkg.in/mgo.v2"
)

var s *server.Server
var conn *mgo.Session

func main() {
    fmt.Println("hello world")
    
	dialInfo, _ := mgo.ParseURL(os.Getenv("MONGO_URI"))
	conn, err := mgo.DialWithInfo(dialInfo)

	conn.SetMode(mgo.Monotonic, true)

	if err != nil {
        fmt.Errorf("mongo: could not dial: %v", err)
        return
    }
    
    requestRepo, err := mongo.NewRequestRepository("status_health", conn)

    recurso := &model.Request{}

    recurso.Name = "Goiabainha"

    requestRepo.Store(recurso)

    s = server.New()
    
    addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	s.Echo.Logger.Fatal(s.Echo.Start(addr))
}
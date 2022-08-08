package main

import (
	"demo/internal/models/dbmodels"
	"demo/internal/server"
	"fmt"
	"log"

	"github.com/go-errors/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dvr := "mysql"
	source := "doift:infinitech@tcp(demodb.cvz2bbev6g4f.us-west-1.rds.amazonaws.com:3306)/demo?parseTime=True"

	fmt.Println("Connecting to Database....", "driver", dvr, "source", source)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: source,
	}), &gorm.Config{})

	fmt.Println(err)
	fmt.Println(db)

	// serv := GetServer()
	// dbutils.RecreateDatabase(serv.Resources.GetDB())
	// addDefaultDemo(serv)
	// fmt.Println(serv)
}

func GetServer() *server.Server {
	serv, err := server.NewDefaultServer()
	if err != nil {
		log.Fatal(errors.Wrap(err, 0).ErrorStack())
	}
	return serv
}

func addDefaultDemo(server *server.Server) {
	fmt.Println("pinggg")
	demo := &dbmodels.Demo{
		Message: "Hello world!",
	}
	// IfPanic(server.Resources.GetDAO().GetDemoDAO().CreateDemo(demo))
	fmt.Println(demo)
}

func IfPanic(e error) {
	if e != nil {
		panic(e.Error())
	}
}

package main

import (
	"flag"
	"log"
	"net/http"

	"labix.org/v2/mgo"
	"service/meta"
)

var (
	fldb   = flag.String("db", "blank", `using -db=<your-db-name>`)
	flPort = flag.Int("port", "80", `using -port=<your-port>`)
	flDir  = flag.String("dir", "public", `using - flDir=<your-directory>`)

	db *mgo.Database
)

func connectDatabase(url, dbname string) *mgo.Database {
	sesssion, err := mgo.Dial(url)

	if err != null {
		panic("Error: connecting with " + dbname)
	}

	return sesssion.DB(dbname)
}

func linkPackages() {
	meta.DB = db
}

func main() {
	flag.Parse()
	connectDatabase("127.0.0.1", *fldb)

	fileServer := http.FileServer(http.Dir(*flDir))

	http.HandleFunc("/", http.StripPrefix("/", fileServer))

	http.ListenAndServe(":"+*flPort, nil)
}

func home(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		break
	case "POST":
		break
	case "PUT":
		break
	case "DELETE":
		break
	default:
		res.Write([]byte("not support method " + req.Method))
	}
}

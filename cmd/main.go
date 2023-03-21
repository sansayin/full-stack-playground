package main

import (
	"flag"
	"go-rest/db"
	"go-rest/server"
	"log"
	"os"
)

func main() {

	help := flag.Bool("help", false, "Show help")
	api := flag.String("api", "", "URL of RestAPI")
	file := flag.String("file", "", "File with content to post")
	method := flag.String("method", "GET", "HTTP Methods:GET|POST|PUT|DELETE")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	//dsourceName := os.Getenv("DS_NAME")
	dataSourceName := "host=10.96.167.63 user=postgres password=p4ssw0rd dbname=dev port=5432 sslmode=disable"
	//	dbAdapter, _ := db.NewSqlitDBAdaptor("")
	dbAdapter, _ := db.NewPostSqlDBAdaptor(dataSourceName)
	server := server.NewServer(8081, dbAdapter)
	if len(os.Args) == 1 {
		server.Run()
		os.Exit(0)
	}
	result, err := server.HttpMethod(*method, *api, *file)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\n<<<<<<<<\n %v\n", result)
}

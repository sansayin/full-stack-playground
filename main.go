package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	API "go-rest/api"
 Graph 	"go-rest/graph"
 "go-rest/db"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type httpServer struct {
	e    *echo.Echo
  da   *db.Adapter
	port int
}

type Options struct {
	METHOD   string
	ENDPOINT string
	CONTENT  string
}

func NewServer(port int, da *db.Adapter) *httpServer {
  return &httpServer{port: port,da:da}
}

func (server httpServer) getIPs(ch chan bool) {
	cmd := "ip -f inet a|grep -oP '(?<=inet ).+(?=\\/)'"
	out, _ := exec.Command("bash", "-c", cmd).Output()
	ips := strings.Split(string(out), "\n")
	for i, ip := range ips {
		if len(strings.TrimSpace(ip)) > 0 {
			fmt.Printf("\t%v. http://%v:%v\n", i+1, ip, server.port)
		}
	}
	ch <- true
}
func (server httpServer) startCWDHttp() {

	var wg sync.WaitGroup
	wg.Add(1)
	done := make(chan bool)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Can not get current directory")
	}
	server.e = echo.New()
	server.e.Use(middleware.Recover())
	server.e.Use(middleware.CORS())
	server.e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   ".",
		Browse: true,
	})) // e.POST("/users", saveUser)
	server.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}, ${uri}, ${status}\n",
	}))

	API.NewRest(server.e, server.da)

  Graph.NewGraph(server.e, server.da)

	go server.e.Start(":" + strconv.Itoa(server.port))
	log.Printf("[Serving in directory] %v\n", dir)
	go server.getIPs(done)
	<-done
	wg.Wait()
}


func (server httpServer) doHttpMethod(method string, endpoint string, content string) (string, error) {

	body, err := os.ReadFile(content)

	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewReader(body))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")
	//	req.ContentLength = contentLength
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	ret, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, ret, "", " ")
	if error != nil {
		log.Fatal("JSON parse error: ", error)
	}
	return prettyJSON.String(), nil
}

func main() {
	help := flag.Bool("help", false, "Show help")
	endpot := flag.String("api", "", "URL of RestAPI")
	conent := flag.String("content", "", "File with content to post")
	method := flag.String("method", "GET", "HTTP Methods:GET|POST|PUT|DELETE")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	//dsourceName := os.Getenv("DS_NAME")
  dataSourceName:="host=postgres user=postgres password=p4ssw0rd dbname=dev port=54320 sslmode=disable"
  dbAdapter, err := db.NewPostSqlDBAdaptor(dataSourceName)
	
  if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	server := NewServer(8081,dbAdapter)
	if len(os.Args) == 1 {
		server.startCWDHttp()
		os.Exit(0)
	}
	result, err := server.doHttpMethod(*method, *endpot, *conent)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\n<<<<<<<<\n %v\n", result)
}

package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	API "go-rest/api"
	"go-rest/db"
	Graph "go-rest/graph"
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

func NewServer(port int, da *db.Adapter) *httpServer {
	return &httpServer{port: port, da: da}
}

func (server httpServer) getIPs() {
	cmd := "ip -f inet a|grep -oP '(?<=inet ).+(?=\\/)'"
	out, _ := exec.Command("bash", "-c", cmd).Output()
	ips := strings.Split(string(out), "\n")
	for i, ip := range ips {
		if len(strings.TrimSpace(ip)) > 0 {
			fmt.Printf("\t%v. http://%v:%v\n", i+1, ip, server.port)
		}
	}
}

func (server httpServer) HttpMethod(method string, url string, file string) (string, error) {
	body, err := os.ReadFile(file)

	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	ret, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, ret, "", " "); err != nil {
		log.Fatal("JSON parse error: ", err)
	}
	return prettyJSON.String(), nil
}

func (server httpServer) Run() {

	var wg sync.WaitGroup
	dir, _ := os.Getwd()
	server.e = echo.New()
	server.e.Use(middleware.Recover())
	server.e.Use(middleware.CORS())
	server.e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "./html",
		Browse: true,
	})) 
	server.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}, ${uri}, ${status}\n",
	}))

	API.NewRest(server.e, server.da)
	Graph.NewGraph(server.e, server.da)

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("[Serving in directory] %v/html\n", dir)
		server.e.Start(":" + strconv.Itoa(server.port))
	}()

  wg.Add(1)
	go func() {
		defer wg.Done()
		server.getIPs()
	}()
	wg.Wait()
}

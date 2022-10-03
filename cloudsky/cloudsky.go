package cloudsky

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)



func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05",
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

// HandlerFunc   defines the request handler used by cloudsky
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor of cloudsky.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) PUT(pattern string, handler HandlerFunc) {
	engine.addRoute("PUT", pattern, handler)
}

func (engine *Engine) DELETE(pattern string, handler HandlerFunc) {
	engine.addRoute("DELETE", pattern, handler)
}

func (engine *Engine) JSON(resp http.ResponseWriter, req *http.Request, data map[interface{}]interface{})   {
	req.Header.Set("content-type", "application/json")
	JsonStr,err := json.Marshal(data)
	if err!=nil{
		log.Error("JSON json.Marshal err! ", err)
	}
	
	_, err = resp.Write(JsonStr)
	if err != nil {
		log.Error(err)
	}

}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	log.Infof("cloudsky engine serve at %s", addr)
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		dataLength, err := fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
		if err != nil {
			fmt.Printf("response data length: %d\n",dataLength)
			return
		}
	}
}
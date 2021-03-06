/* Obsoleted by gen_router.go :( */
package main

//import "fmt"
import "strings"
import "sync"
import "net/http"

// TO-DO: Support the new handler signatures created by our efforts to move the PreRoute middleware into the generated router
type Router struct {
	sync.RWMutex
	routes map[string]func(http.ResponseWriter, *http.Request)
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

func (router *Router) Handle(pattern string, handle http.Handler) {
	router.Lock()
	router.routes[pattern] = handle.ServeHTTP
	router.Unlock()
}

func (router *Router) HandleFunc(pattern string, handle func(http.ResponseWriter, *http.Request)) {
	router.Lock()
	router.routes[pattern] = handle
	router.Unlock()
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if len(req.URL.Path) == 0 || req.URL.Path[0] != '/' {
		w.WriteHeader(405)
		w.Write([]byte(""))
		return
	}

	var /*extra_data, */prefix string
	if req.URL.Path[len(req.URL.Path) - 1] != '/' {
		//extra_data = req.URL.Path[strings.LastIndexByte(req.URL.Path,'/') + 1:]
		prefix = req.URL.Path[:strings.LastIndexByte(req.URL.Path,'/') + 1]
	} else {
		prefix = req.URL.Path
	}

	router.RLock()
	handle, ok := router.routes[prefix]
	router.RUnlock()

	if ok {
		handle(w,req)
		return
	}
	//log.Print("req.URL.Path[:strings.LastIndexByte(req.URL.Path,'/')]",req.URL.Path[:strings.LastIndexByte(req.URL.Path,'/')])
	NotFound(w,req)
}

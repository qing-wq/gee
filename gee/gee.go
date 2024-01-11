package gee

import "net/http"

type H map[string]interface{}

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.router.addRouter("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.router.addRouter("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// ServeHTTP Take over all HTTP requests
func (e *Engine) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c := newContext(rw, req)
	e.router.handle(c)
}

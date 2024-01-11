package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern

	root, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	root.insert(pattern, parsePattern(pattern), 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (string, map[string]string) {
	root, ok := r.roots[method]
	if !ok {
		return "", nil
	}

	searchParts := parsePattern(path)
	pattern := root.search(searchParts, 0)
	if pattern == "" {
		return "", nil
	}

	params := make(map[string]string)
	for i, part := range parsePattern(pattern) {
		if part[0] == ':' {
			params[part[1:]] = searchParts[i]
		}

		if part[0] == '*' {
			params[part[1:]] = strings.Join(searchParts[i:], "/")
			break
		}
	}
	return pattern, nil
}

func (r *router) handle(c *Context) {
	pattern, params := r.getRoute(c.Method, c.Path)
	if pattern != "" {
		c.Params = params
		reqKey := c.Method + "-" + pattern
		r.handlers[reqKey](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

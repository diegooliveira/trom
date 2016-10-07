package router

import (
	"sort"
	"strings"
)

// Router takes care of link incoming paths to backend nodes
type Router struct {
	routes []*Route
	count  int
}

// CreateNew gives you a prepared Router instance
func CreateNew() *Router {
	g := new(Router)
	g.routes = make([]*Route, 0)

	return g
}

func (g *Router) GetRouteFor(path string) *Route {

	for _, r := range g.routes {

		if strings.HasPrefix(path, r.path) {

			return r
		}
	}

	panic("")
}

func (g *Router) AddRoute(path string) *Route {

	route := new(Route)
	route.path = path
	route.nodes = make([]string, 0)

	g.routes = append(g.routes, route)
	g.count++

	sort.Sort(g)

	return route
}

func (g Router) Len() int {
	return g.count
}

func (g Router) Less(i, j int) bool {
	return len(g.routes[j].path) < len(g.routes[i].path)
}

func (g Router) Swap(i, j int) {
	g.routes[i], g.routes[j] = g.routes[j], g.routes[i]
}

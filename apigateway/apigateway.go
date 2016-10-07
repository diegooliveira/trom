package apigateway

import (
	"log"
	"net/http"
	"net/http/httputil"
	"trom/router"
)

type PreHandler interface {
	PreHandle(rw http.ResponseWriter, req *http.Request) bool
}

// ApiGateway that start listening connections
type ApiGateway struct {
	g           *router.Router
	preHandlers []PreHandler
	proxy       http.Handler
}

func New(g *router.Router) *ApiGateway {
	apiGateway := new(ApiGateway)
	apiGateway.g = g
	apiGateway.preHandlers = make([]PreHandler, 0)
	apiGateway.proxy = &httputil.ReverseProxy{
		Director: apiGateway.hostResolver,
	}
	return apiGateway
}

func (ag *ApiGateway) AddPreHandler(pre PreHandler) {
	ag.preHandlers = append(ag.preHandlers, pre)
}

func (ag *ApiGateway) hostResolver(req *http.Request) {

	route := ag.g.GetRouteFor(req.URL.Path)

	req.URL.Scheme = "http"
	req.URL.Host = route.Host()
}

func (ag *ApiGateway) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	for _, preHandler := range ag.preHandlers {
		if !preHandler.PreHandle(rw, req) {
			return
		}
	}

	ag.proxy.ServeHTTP(rw, req)

}

func (ag *ApiGateway) Start() {
	log.Fatal(http.ListenAndServe(":9090", ag))
}

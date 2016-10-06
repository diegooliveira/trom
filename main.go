package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"trom/auth"
	"trom/gateway"
	"trom/server"
)

type Registry map[string][]string

func extractNameVersion(target *url.URL) (name, version string, err error) {
	path := target.Path
	// Trim the leading `/`
	if len(path) > 1 && path[0] == '/' {
		path = path[1:]
	}
	// Explode on `/` and make sure we have at least
	// 2 elements (service name and version)
	tmp := strings.Split(path, "/")
	if len(tmp) < 2 {
		return "", "", fmt.Errorf("Invalid path")
	}
	name, version = tmp[0], tmp[1]
	// Rewrite the request's path without the prefix.
	target.Path = "/" + strings.Join(tmp[2:], "/")
	return name, version, nil
}

// NewMultipleHostReverseProxy creates a reverse proxy that will randomly
// select a host from the passed `targets`
func NewMultipleHostReverseProxy(reg Registry) *httputil.ReverseProxy {
	director := func(req *http.Request) {

		for key, val := range reg {
			if strings.HasPrefix(req.URL.Path, key) {
				log.Printf(val[0])

				req.URL.Scheme = "http"
				req.URL.Host = val[0]
				return
			}
		}
	}
	return &httputil.ReverseProxy{
		Director: director,
	}
}

func main() {

	gateway := gateway.New()

	route := gateway.Route("/economia")
	route.AddNode("localhost:9001")
	route.AddNode("localhost:9002")

	route = gateway.Route("/samba")
	route.AddNode("localhost:8001")
	route.AddNode("localhost:8002")

	auth := auth.New()

	server := server.New(gateway, auth)
	server.Start()

	proxy := NewMultipleHostReverseProxy(Registry{
		"/economia": {
			"",
		},
		"/educacao": {
			"localhost:8888",
		},
		"/": {
			"localhost:7777",
		},
	})
	log.Fatal(http.ListenAndServe(":9090", proxy))
}

package gateway

type Gateway struct {
}

func New() *Gateway {
	g := new(Gateway)

	return g
}

type Route struct {
	contextPath string
	target      []string
}

func (p *Route) AddNode(host string) {

}

func (g *Gateway) Route(contextPath string) *Route {

	target := new(Route)
	target.contextPath = contextPath

	return target
}

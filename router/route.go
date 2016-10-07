package router

type Route struct {
	path  string
	nodes []string
}

func (p *Route) AddNode(node string) {

	p.nodes = append(p.nodes, node)

}

func (r *Route) Host() string {
	return r.nodes[0]
}

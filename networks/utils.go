package networks

import "github.com/gin-gonic/gin"

type Method int8

const (
	GET Method = iota
	POST
)

func (n *Network) Router (method Method, path string, handler gin.HandlerFunc) {
	switch method  {
		case GET:
			n.engine.GET(path, handler)
		case POST:
			n.engine.POST(path, handler)
	}
}

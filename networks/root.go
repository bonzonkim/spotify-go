package networks

import "github.com/gin-gonic/gin"

type Network struct {
	engine *gin.Engine
}

func NewNetwork() *Network {
	n := &Network{
		engine: gin.New(),
	}
	return n
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}

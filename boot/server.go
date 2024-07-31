package boot

import (
	"fmt"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/types"
)

// struct server in project
//
// host:port - 127.0.0.1:8000
type _Server struct {
	address string //
}

// NewServer creates a new server instance
// Use for fx Framework and more
func NewServer() IServer {
	return &_Server{
		address: fmt.Sprintf("%s:%s", config.Host, config.Port),
	}
}

// Connect to module via adapter
//
//	func main() {
//		var (
// 			baseService    = base.New()
// 			baseController = controller.New(baseService)
// 			baseRouter     = router.New(baseController)
//			httpServer = webapi.NewServer(baseRouter)
//			adapt      = webapi.NewBaseAdapter(httpServer)
//			server     = boot.NewServer()
//
//		)
//
//		log.Fatal(server.Connect(adapt))
//	}
func (server *_Server) Connect(adapter types.IAdapter) error {
	return adapter.Run(server.address)
}

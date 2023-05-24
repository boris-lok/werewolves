package lib

import (
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Run(config Configuration, listener net.Listener) {
  engine := gin.Default()

  mode := strings.ToLower(config.Application.Mode)
  switch mode {
    case "test":
      gin.SetMode(gin.TestMode)
    case "release":
      gin.SetMode(gin.ReleaseMode)
    default:
      gin.SetMode(gin.DebugMode)
  }

  srv := &http.Server{
    Handler: engine.Handler(),
  }

  go func() {
    if err := srv.Serve(listener); err != nil && err != http.ErrServerClosed {
      log.Fatalf("Can't create a server, get an {%v}", err)
    }
  }()

  return srv, nil
}

package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GenericAPIServer struct {
	middlewares         []string
	SecureServingInfo   *SecureServingInfo
	InsecureServingInfo *InsercureServingInfo

	ShutdownTimeout time.Duration
	*gin.Engine
	healthz         bool
	EnableMetrics   bool
	EnableProfiling bool

	insecureServer, secureServer http.Server
}

func intGenericAPIServers(s *GenericAPIServer) {

}

func (g *GenericAPIServer) InstallAPIs() {
	// install healthz handler
	if g.healthz {
		g.GET("/healthz", func(ctx *gin.Context) {
			// todo: core...
		})
	}

	if g.EnableMetrics {
		// todo
	}

	if g.EnableProfiling {
		// todo  pprof handler
	}
	g.GET("/version", func(ctx *gin.Context) {
		// todo   core

	})
}

func (g *GenericAPIServer) SetUp() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		// todo
	}
}

func (g *GenericAPIServer) InstallMiddlerwares() {
	// g.Use()
	// g.Use()

	for _, m := range g.middlewares {
		fmt.Println(m)
		// ...
		// g.Use(m)
	}
}

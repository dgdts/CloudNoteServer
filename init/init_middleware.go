package init

import (
	"context"
	"sort"

	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgdts/UniversalServer/internal/middleware"
	"github.com/dgdts/UniversalServer/pkg/config"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/pprof"
)

func initMiddleware(s *server.Hertz, config *config.GlobalConfig) {
	// pprof, should not be used in production environment
	if config.Hertz.EnablePprof {
		pprof.Register(s)
	}

	// gzip
	if config.Hertz.EnableGzip {
		s.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if config.Hertz.EnableAccessLog {
		s.Use(accesslog.New())
	}

	// recovery
	s.Use(recovery.Recovery())

	// cors
	defaultCorsConfig := cors.DefaultConfig()
	defaultCorsConfig.ExposeHeaders = append(defaultCorsConfig.ExposeHeaders, "Content-Disposition")
	defaultCorsConfig.AllowHeaders = append(defaultCorsConfig.AllowHeaders,
		"Authorization",
		"Content-Type",
		"Origin",
		"Accept",
		"X-Requested-With",
	)
	defaultCorsConfig.AllowAllOrigins = true
	defaultCorsConfig.AllowCredentials = true
	s.Use(cors.New(defaultCorsConfig))

	// custom middleware
	initCustomMiddleware(s)
}

// initCustomMiddleware used to check the http header for business logic
func initCustomMiddleware(s *server.Hertz) {
	filters := middleware.GetAllFilters()

	if len(filters) == 0 {
		return
	}

	sort.Slice(filters, func(i, j int) bool {
		return filters[i].GetOrder() < filters[j].GetOrder()
	})

	for _, filter := range filters {
		s.Use(filter.DoFilter)
		hlog.CtxInfof(context.Background(), "init filter: %v", filter.Name())
	}
}

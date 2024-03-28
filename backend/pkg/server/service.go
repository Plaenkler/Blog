package server

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/rs/zerolog/log"

	"github.com/plaenkler/blog/pkg/config"
	"github.com/plaenkler/blog/pkg/server/routes/web"
)

var (
	//go:embed routes/web/static
	static embed.FS
	server = &http.Server{
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
)

func init() {
	r := NewRouter()
	registerRoutes(r)
	registerStaticFiles(r)
	server.Handler = r.ServeMux
}

func registerRoutes(r *Router) {
	for _, path := range web.GetRoutes() {
		r.Handle(path, gziphandler.GzipHandler(http.HandlerFunc(web.ProvideRoute)))
	}
}

func registerStaticFiles(r *Router) {
	staticHandler := createStaticHandler()
	r.Handle("/img/", staticHandler)
	r.Handle("/js/", staticHandler)
	r.Handle("/dist/", staticHandler)
	r.Handle("/assets/", staticHandler)
}

func createStaticHandler() http.Handler {
	fs, err := fs.Sub(static, "routes/web/static")
	if err != nil {
		log.Fatal().Err(err).Msg("[server-createStaticHandler-1] could not create static handler")
	}
	return gziphandler.GzipHandler(controlCache(http.FileServer(http.FS(fs))))
}

func Start() {
	server.Addr = fmt.Sprintf(":%v", config.Get().Port)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("[server-Start-1] could not initialize server")
	}
}

func Stop() {
	err := server.Shutdown(context.Background())
	if err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("[server-Stop-1] could not shutdown server")
	}
}

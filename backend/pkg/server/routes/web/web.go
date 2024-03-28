package web

import (
	"bytes"
	"embed"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"github.com/plaenkler/blog/pkg/config"
	"github.com/rs/zerolog/log"
)

var (
	//go:embed static
	static embed.FS
	mutex  sync.RWMutex
	routes = make(map[string]route)
)

type route struct {
	page    bool
	content []byte
}

type urlset struct {
	XMLName string `xml:"urlset"`
	XMLNS   string `xml:"xmlns,attr"`
	URLs    []url  `xml:"url"`
}

type url struct {
	Loc string `xml:"loc"`
}

// Register index as route
func init() {
	content := parseTemplate(
		"index",
		"static/html/pages/index.html",
	)
	mutex.Lock()
	defer mutex.Unlock()
	routes["/"] = route{page: true, content: content}
}

// Register imprint as route
func init() {
	content := parseTemplate(
		"imprint",
		"static/html/pages/imprint.html",
	)
	mutex.Lock()
	defer mutex.Unlock()
	routes["/imprint"] = route{page: true, content: content}
}

// Register sitemap.xml as route
func init() {
	sitemap := urlset{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  make([]url, 0),
	}
	mutex.RLock()
	for path := range routes {
		if !routes[path].page {
			continue
		}
		sitemap.URLs = append(sitemap.URLs, url{
			Loc: fmt.Sprintf("%s%s", config.Get().URL, path),
		})
	}
	mutex.RUnlock()
	output, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		log.Error().Err(err).Msg("[web-init-7] could not marshal sitemap")
	}
	mutex.Lock()
	defer mutex.Unlock()
	routes["/sitemap.xml"] = route{content: append([]byte(xml.Header), output...)}
}

// Register robots.txt as route
func init() {
	mutex.Lock()
	defer mutex.Unlock()
	routes["/robots.txt"] = route{content: []byte(fmt.Sprintf("User-agent: *\nSitemap: %s/sitemap.xml", config.Get().URL))}
}

// Provide URL specific content
func ProvideRoute(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	defer mutex.RUnlock()
	route, ok := routes[r.URL.Path]
	if !ok {
		route, ok = routes["/"]
		if !ok {
			return
		}
	}
	if route.page {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
	}
	_, err := w.Write(route.content)
	if err != nil {
		log.Error().Err(err).Msgf("[web-ProvidePage-1] could not write page %s", r.URL.Path)
	}
}

// Return all registered routes
func GetRoutes() []string {
	mutex.RLock()
	defer mutex.RUnlock()
	routePaths := make([]string, 0)
	for path := range routes {
		routePaths = append(routePaths, path)
	}
	return routePaths
}

// Parse html template files
func parseTemplate(name string, files ...string) []byte {
	tpl, err := template.New(name).ParseFS(static,
		files...,
	)
	if err != nil {
		log.Error().Err(err).Msgf("[web-parseTemplate-1] could not parse template %s", name)
		return []byte{}
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, nil)
	if err != nil {
		log.Error().Err(err).Msgf("[web-parseTemplate-2] could not render template %s", name)
		return []byte{}
	}
	return buf.Bytes()
}

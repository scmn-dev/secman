package api

import (
	"os"
	"path"
	"strings"
	"net/http"
	"path/filepath"
)

type StaticServer struct {
	Root            http.Dir
	FallbackHandler http.Handler
}

func NewSFS(root http.Dir, fallback http.HandlerFunc) *StaticServer {
	return &StaticServer{
		Root:            root,
		FallbackHandler: fallback,
	}
}

func (fs StaticServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path

	if !strings.HasPrefix(p, "/") {
		p = "/" + p
		r.URL.Path = p
	}

	p = path.Clean(p)

	name := path.Join(string(fs.Root), filepath.FromSlash(p))

	f, err := os.Open(name)

	if err != nil {
		if os.IsNotExist(err) {
			fs.FallbackHandler.ServeHTTP(w, r)
			return
		}
	}

	defer f.Close()

	http.ServeFile(w, r, name)
}

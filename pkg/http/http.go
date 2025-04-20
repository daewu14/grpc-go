package httppkg

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// NewServer creates and initializes a new grpcPkg instance with the given context and list of gRPC services.
func NewServer(ctx context.Context, addr string) *httpPkg {
	return &httpPkg{
		handlers: make([]map[string]Http, 0),
		ctx:      ctx,
		address:  addr,
	}
}

type httpPkg struct {
	address  string
	ctx      context.Context
	handlers []map[string]Http
}

func (h *httpPkg) Handle(path string, server Http) {
	path = "/" + path
	h.handlers = append(h.handlers, map[string]Http{path: server})
}

func (h *httpPkg) Run() {
	router := http.NewServeMux()

	for _, s := range h.handlers {
		for path, http2 := range s {
			router.HandleFunc(fmt.Sprintf("%s %s", http2.Method(), path), http2.Handle)
		}
	}

	log.Print("http server start run with total handlers: ", len(h.handlers), " on port: ", h.address)
	err := http.ListenAndServe(fmt.Sprintf(":%s", h.address), router)
	if err != nil {
		log.Fatalf("http server listen err: %v", err)
	}
}

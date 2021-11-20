package http

import (
	"api/internal/store"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ServerOption func(srv *Server)

func WithAddress(address string) ServerOption {
	return func(srv *Server) {
		srv.Address = address
	}
}

func WithStore(store store.Store) ServerOption {
	return func(srv *Server) {
		srv.store = store
	}
}

func WithElastic(es esapi.Info) ServerOption {
	return func(srv *Server) {
		srv.es = es
	}
}
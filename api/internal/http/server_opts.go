package http

import (
	"api/internal/store"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/cache/v8"
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

func WithRedisCache(mycache *cache.Cache) ServerOption {
	return func(srv *Server) {
		srv.mycache = mycache
	}
}

func WithElastic(es *elasticsearch.Client) ServerOption {
	return func(srv *Server) {
		srv.es = es
	}
}
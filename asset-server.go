package sk

import (
	"os"
	"sync"
)

type AssetServer struct {
	cache map[string]*Handler
	mutex sync.RWMutex
}

func NewAssetServer() *AssetServer {
	return &AssetServer{
		cache: make(map[string]*Handler),
		mutex: sync.RWMutex{},
	}
}

func (as *AssetServer) fetch(path string, handler *Handler) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		handler.SetError(err)
	} else {
		handler.SetData(bytes)
	}
}

func (as *AssetServer) Load(path string) *Handler {
	as.mutex.Lock()
	handler, ok := as.cache[path]

	if ok && (handler.ready || handler.loading) {
		as.mutex.Unlock()
		return handler
	}

	if !ok {
		handler = NewHandler(path)
		as.cache[path] = handler
	}

	as.fetch(path, handler)

	as.mutex.Unlock()
	return handler
}

func (as *AssetServer) LoadAsync(path string) *Handler {
	as.mutex.Lock()
	handler, ok := as.cache[path]

	if !ok {
		handler = NewHandler(path)
		as.cache[path] = handler
	}
	as.mutex.Unlock()

	go as.Load(path)
	return handler
}

func (as *AssetServer) Unload(path string) {
	as.mutex.Lock()
	delete(as.cache, path)
	as.mutex.Unlock()
}

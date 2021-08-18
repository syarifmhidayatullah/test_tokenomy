package repository

import (
	entity "github.com/syarifmhidayatullah/test_tokenomy/internal/model"
)

type DataRepo struct {
	Cache MemoryCache
}

func NewDataRepo(cache MemoryCache) DataRepo {
	data := DataRepo{
		Cache: cache,
	}
	return data
}

type MemoryCache struct {
	DataMap map[int]entity.Data
}

func NewMemoryCache() MemoryCache {
	return MemoryCache{}
}

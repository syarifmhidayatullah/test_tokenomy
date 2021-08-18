package repository

import (
	entity "github.com/syarifmhidayatullah/test_tokenomy/internal/model"
)

func (m *MemoryCache) ConstructData() {
	dataMap := map[int]entity.Data{}
	for _, v := range entity.ListData {
		dataMap[v.Id] = v
	}
	m.SetDataCache(dataMap)
}

func (m *MemoryCache) SetDataCache(dataMap map[int]entity.Data) {
	m.DataMap = dataMap
}

func (m *MemoryCache) GetDataCache() (dataMap map[int]entity.Data) {
	return m.DataMap
}

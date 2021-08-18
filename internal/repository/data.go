package repository

import (
	entity "github.com/syarifmhidayatullah/test_tokenomy/internal/model"
)

func (d *DataRepo) GetDataList() (dataList []entity.Data) {
	return entity.ListData
}

func (d *DataRepo) GetDataListByParam(param []int) (dataList []entity.Data) {
	dataCache := d.Cache.GetDataCache()
	for _, v := range param {
		if _, ok := dataCache[v]; !ok {
			continue
		}
		dataList = append(dataList, dataCache[v])
	}
	return dataList
}

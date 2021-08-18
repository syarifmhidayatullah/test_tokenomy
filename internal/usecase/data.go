package usecase

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	entity "github.com/syarifmhidayatullah/test_tokenomy/internal/model"
)

func (d *DataCase) GetDataList() (dataList []entity.Data) {
	dataList = d.DataRepo.GetDataList()
	return
}

func (d *DataCase) GetDataListByParam(param string) (dataList []entity.Data, errType int, err error) {
	dataList = []entity.Data{}
	paramList := strings.Split(param, ",")
	switch len(paramList) {
	case 0:
		return dataList, http.StatusBadRequest, fmt.Errorf("invalid or empty ID: \"%s\"", param)
	case 1:
		return d.processSingleParam(paramList[0])
	default:
		return d.processMultiParam(paramList)
	}
}

func (d *DataCase) processMultiParam(params []string) (dataList []entity.Data, errType int, err error) {
	dataList = []entity.Data{}
	paramInts := []int{}
	for _, v := range params {
		paramInt, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		paramInts = append(paramInts, paramInt)
	}
	if len(paramInts) <= 0 {
		return dataList, http.StatusBadRequest, fmt.Errorf("invalid or empty ID: \"%s\"", params)
	}

	dataList = d.DataRepo.GetDataListByParam(paramInts)
	if dataList == nil {
		return dataList, http.StatusNotFound, fmt.Errorf("resource with ID %s not exist", params)
	}
	return dataList, http.StatusOK, nil
}

func (d *DataCase) processSingleParam(param string) (dataList []entity.Data, errType int, err error) {
	dataList = []entity.Data{}
	errType = http.StatusOK
	paramInt, err := strconv.Atoi(param)
	if err != nil {
		return dataList, http.StatusBadRequest, fmt.Errorf("invalid or empty ID: \"%s\"", param)
	}

	dataList = d.DataRepo.GetDataListByParam([]int{paramInt})
	if len(dataList) <= 0 {
		return dataList, http.StatusNotFound, fmt.Errorf("resource with ID %d not exist", paramInt)
	}

	return
}

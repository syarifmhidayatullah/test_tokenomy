package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	entity "github.com/syarifmhidayatullah/test_tokenomy/internal/model"
)

func (a *Api) GetDataList(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("id")
	if value == "" {
		log.Println("value empty return all list")
		dataList := a.DataCase.GetDataList()
		resp := entity.JsonResponse{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    dataList,
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	dataList, errType, err := a.DataCase.GetDataListByParam(value)
	resp := entity.JsonResponse{
		Code:    errType,
		Message: http.StatusText(http.StatusOK),
		Data:    dataList,
	}

	if errType != http.StatusOK {
		resp.Message = err.Error()
	}

	json.NewEncoder(w).Encode(resp)
}

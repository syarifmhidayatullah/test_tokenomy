package delivery

import (
	"net/http"

	usecase "github.com/syarifmhidayatullah/test_tokenomy/internal/usecase"
)

type Api struct {
	DataCase usecase.DataCase
}

func NewApi(dataCase usecase.DataCase) Api {
	api := Api{
		DataCase: dataCase,
	}
	api.RegisterRoutes()
	return api
}

func (a *Api) RegisterRoutes() {
	http.HandleFunc("/", a.GetDataList)
}

package usecase

import (
	repo "github.com/syarifmhidayatullah/test_tokenomy/internal/repository"
)

type DataCase struct {
	DataRepo repo.DataRepo
}

func NewDataCase(dataRepo repo.DataRepo) DataCase {
	dataCase := DataCase{
		DataRepo: dataRepo,
	}
	return dataCase
}

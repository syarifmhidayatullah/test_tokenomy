package main

import (
	"log"
	"net/http"

	del "github.com/syarifmhidayatullah/test_tokenomy/internal/delivery"
	repo "github.com/syarifmhidayatullah/test_tokenomy/internal/repository"
	ucase "github.com/syarifmhidayatullah/test_tokenomy/internal/usecase"
)

func main() {
	cache := repo.NewMemoryCache()
	cache.ConstructData()
	repository := repo.NewDataRepo(cache)
	usecase := ucase.NewDataCase(repository)
	del.NewApi(usecase)
	log.Println("listen on port 9999")
	http.ListenAndServe(":9999", nil)
}

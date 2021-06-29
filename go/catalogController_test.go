package main

import (
	"encoding/json"
	"github.com/sogilis/poc-wiismile/model"
	"github.com/sogilis/poc-wiismile/repository"
	"github.com/sogilis/poc-wiismile/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func resetData() {
	repository.Catalogs = []model.Catalog{
		{
			Id:           utils.CreateUuid(),
			Enabled:      true,
			SupplierName: "Leclerc",
			CreatedAt:    "admin",
			CreatedBy:    "06/03/2021",
		},
		{
			Id:           utils.CreateUuid(),
			Enabled:      false,
			SupplierName: "Fnac",
			CreatedAt:    "user",
			CreatedBy:    "08/03/2021",
		},
	}
}

func structToString(a interface{}) string {
	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func TestCatalogController(t *testing.T) {
	t.Run("Get catalog list", GetCatalogList)
	t.Run("Get catalog by ID", GetCatalogById)
	t.Run("Add catalog", AddCatalog)
	t.Run("Update catalog", UpdateCatalog)
	t.Run("Delete catalog", DeleteCatalog)
}

func GetCatalogList(t *testing.T) {

	// given
	resetData()
	router := initRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/API/catalogs", nil)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, structToString(repository.Catalogs), w.Body.String())
}

func GetCatalogById(t *testing.T) {

	// given
	resetData()
	router := initRouter()
	w := httptest.NewRecorder()

	catalogToAssert := repository.Catalogs[0]

	req, _ := http.NewRequest("GET", "/API/catalog/"+catalogToAssert.Id.String(), nil)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, structToString(catalogToAssert), w.Body.String())
}

func AddCatalog(t *testing.T) {

	// given
	resetData()
	router := initRouter()
	w := httptest.NewRecorder()

	catalogToAdd := model.Catalog{
		Id:           utils.CreateUuid(),
		SupplierName: "test",
		Enabled:      true,
		CreatedBy:    "admin",
		CreatedAt:    "07/08/2012",
	}

	c := strings.NewReader(structToString(catalogToAdd))
	req, _ := http.NewRequest("POST", "/API/catalog", c)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, 3, len(repository.Catalogs))
	assert.Equal(t, repository.Catalogs[2], catalogToAdd)
}

func UpdateCatalog(t *testing.T) {

	// given
	resetData()
	router := initRouter()
	w := httptest.NewRecorder()

	catalogToUpdate := repository.Catalogs[1]

	catalogToUpdate.SupplierName = "TestUpdate"
	catalogToUpdate.Enabled = true
	catalogToUpdate.CreatedBy = "TestAdmin"

	c := strings.NewReader(structToString(catalogToUpdate))
	req, _ := http.NewRequest("POST", "/API/catalog", c)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 2, len(repository.Catalogs))
	assert.Equal(t, repository.Catalogs[1], catalogToUpdate)
}

func DeleteCatalog(t *testing.T) {

	// given
	resetData()
	router := initRouter()
	w := httptest.NewRecorder()

	id := repository.Catalogs[1].Id.String()

	req, _ := http.NewRequest("DELETE", "/API/catalog/"+id, nil)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, len(repository.Catalogs))
	assert.NotEqual(t, repository.Catalogs[0].Id.String(), id)
}

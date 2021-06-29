package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sogilis/poc-wiismile/model"
	"github.com/sogilis/poc-wiismile/utils"
)

var Catalogs = []model.Catalog{
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

func GetCatalogList() ([]model.Catalog, error) {
	return Catalogs, nil
}

func GetCatalogById(id uuid.UUID) (model.Catalog, error) {
	for _, catalog := range Catalogs {
		if catalog.Id == id {
			return catalog, nil
		}
	}

	return model.Catalog{}, nil
}

func GetCatalogsCount() (int, error) {
	return len(Catalogs), nil
}

func SaveCatalog(catalog model.Catalog) (bool, error) {
	fmt.Println("test")
	isCatalogExist, err := GetCatalogById(catalog.Id)
	fmt.Println("test")
	if err != nil {
		return false, err
	}
	fmt.Println("test")
	if isCatalogExist != (model.Catalog{}) {
		err := updateCatalog(catalog)
		if err != nil {
			return false, err
		}
		return false, nil
	} else {
		fmt.Println("test")
		err := insertCatalog(catalog)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func insertCatalog(catalog model.Catalog) error {

	Catalogs = append(Catalogs, catalog)

	return nil
}

func updateCatalog(catalog model.Catalog) error {
	for i, c := range Catalogs {
		if c.Id == catalog.Id {
			Catalogs[i] = catalog
		}
	}

	return nil
}

func DeleteCatalog(id uuid.UUID) error {

	for i, c := range Catalogs {
		if c.Id == id {
			Catalogs = append(Catalogs[:i], Catalogs[i+1:]...)
			break
		}
	}

	return nil
}

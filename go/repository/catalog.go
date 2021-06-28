package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sogilis/poc-wiismile/model"
	"github.com/sogilis/poc-wiismile/utils"
)

var catalogs = []model.Catalog{
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
	return catalogs, nil
}

func GetCatalogById(id uuid.UUID) (model.Catalog, error) {
	for _, catalog := range catalogs {
		if catalog.Id == id {
			return catalog, nil
		}
	}

	return model.Catalog{}, errors.New("No match ID")
}

func GetCatalogsCount() (int, error) {
	return len(catalogs), nil
}

func SaveCatalog(catalog model.Catalog) (bool, error) {
	isCatalogExist, err := GetCatalogById(catalog.Id)

	if err != nil {
		return false, err
	}

	if isCatalogExist != (model.Catalog{}) {
		err := updateCatalog(catalog)
		if err != nil {
			return false, err
		}
		return false, nil
	} else {
		err := insertCatalog(catalog)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func insertCatalog(catalog model.Catalog) error {

	catalogs = append(catalogs, catalog)

	return nil
}

func updateCatalog(catalog model.Catalog) error {
	for _, c := range catalogs {
		if c.Id == catalog.Id {
			c = catalog
		}
	}

	return nil
}

func DeleteCatalog(id uuid.UUID) error {

	for i, c := range catalogs {
		if c.Id == id {
			catalogs = append(catalogs[:i], catalogs[i+1:]...)
			break
		}
	}

	return nil
}

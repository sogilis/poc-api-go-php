package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sogilis/poc-wiismile/model"
	"github.com/sogilis/poc-wiismile/repository"
	"net/http"
)

func GetCatalogs(c *gin.Context) {

	catalogs, err := repository.GetCatalogList()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, catalogs)
}

func GetCatalogById(c *gin.Context) {
	catalogId, err := getCatalogId(c)

	catalog, err := repository.GetCatalogById(catalogId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, catalog)
}

func SaveCatalog(c *gin.Context) {
	var catalog model.Catalog
	_ = json.NewDecoder(c.Request.Body).Decode(&catalog)

	isCreated, err := repository.SaveCatalog(catalog)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if isCreated {
		c.Status(http.StatusCreated)
	} else {
		c.Status(http.StatusOK)
	}
}

func DeleteCatalog(c *gin.Context) {
	catalogId, err := getCatalogId(c)

	err = repository.DeleteCatalog(catalogId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}

func getCatalogId(c *gin.Context) (uuid.UUID, error) {
	id := c.Param("id")
	catalogId, err := uuid.Parse(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	return catalogId, err
}

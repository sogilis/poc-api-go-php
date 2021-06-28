package model

import "github.com/google/uuid"

// Catalog Model
type Catalog struct {
	Id           uuid.UUID `json:"id"`
	SupplierName string    `json:"supplierName"`
	Enabled      bool      `json:"enabled"`
	CreatedAt    string    `json:"createdAt"`
	CreatedBy    string    `json:"createdBy"`
}

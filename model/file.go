package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	TypeModel TYPE_MODEL `json:"typeModel"`
	Format    string     `json:"format"`
	Data      []byte     `json:"data"`
	ProductID string     `json:"productId"`
}

type TYPE_MODEL string

const (
	PRODUCT TYPE_MODEL = "PRODUCT"
	SHOP    TYPE_MODEL = "SHOP"
)

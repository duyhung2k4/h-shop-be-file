package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Format    string `json:"format"`
	IsAvatar  *bool  `json:"isAvatar"`
	Name      string `json:"name"`
	Size      uint64 `json:"size"`
	Data      []byte `json:"data"`
	ProductID string `json:"productId"`
}

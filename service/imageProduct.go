package service

import (
	"app/config"
	"app/model"

	"gorm.io/gorm"
)

type imageProductService struct {
	db *gorm.DB
}

type ImageProductService interface {
	GetImagesByProductId(productId string) ([]*model.File, error)
	GetAvatarByProductId(productId string) (*model.File, error)
}

func (s *imageProductService) GetImagesByProductId(productId string) ([]*model.File, error) {
	var images []*model.File

	if err := s.db.Model(&model.File{}).Where("product_id = ?", productId).Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func (s *imageProductService) GetAvatarByProductId(productId string) (*model.File, error) {
	var images *model.File

	if err := s.db.Model(&model.File{}).Where("product_id = ? AND is_avatar = ?", productId, true).Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func NewImageProductService() ImageProductService {
	return &imageProductService{
		db: config.GetDB(),
	}
}

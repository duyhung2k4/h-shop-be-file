package response

import "app/model"

type ImageProductResponse struct {
	Avatar *model.File  `json:"avatar"`
	Images []model.File `json:"images"`
}

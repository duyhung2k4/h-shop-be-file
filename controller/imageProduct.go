package controller

import (
	"app/dto/response"
	"app/model"
	"app/service"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type imageProductController struct {
	imageProductService service.ImageProductService
}

type ImageProductController interface {
	GetImagesbyProductId(w http.ResponseWriter, r *http.Request)
	GetAvatarByProductId(w http.ResponseWriter, r *http.Request)
}

func (c *imageProductController) GetImagesbyProductId(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	productId := query.Get("id")

	if productId == "" {
		badRequest(w, r, errors.New("product id not empty"))
	}

	images, err := c.imageProductService.GetImagesByProductId(productId)
	if err != nil {
		internalServerError(w, r, err)
	}

	imageResponse := response.ImageProductResponse{
		Images: []model.File{},
		Avatar: nil,
	}
	for _, item := range images {
		if item.IsAvatar != nil {
			imageResponse.Avatar = item
		} else {
			imageResponse.Images = append(imageResponse.Images, *item)
		}
	}

	res := Response{
		Data:    imageResponse,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *imageProductController) GetAvatarByProductId(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	productId := query.Get("id")

	if productId == "" {
		badRequest(w, r, errors.New("product id not empty"))
	}

	image, err := c.imageProductService.GetAvatarByProductId(productId)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    image,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func NewImageProductController() ImageProductController {
	return &imageProductController{
		imageProductService: service.NewImageProductService(),
	}
}

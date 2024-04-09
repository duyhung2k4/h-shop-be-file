package api

import (
	"app/config"
	"app/grpc/proto"
	"app/model"
	"context"

	"gorm.io/gorm"
)

type fileGRPC struct {
	db *gorm.DB
	proto.UnsafeFileServiceServer
}

func (g *fileGRPC) InsertFile(ctx context.Context, req *proto.InsertFileReq) (*proto.InsertFileRes, error) {
	var newFiles []model.File
	for _, data := range req.Data {
		newFiles = append(newFiles, model.File{
			ProductID: req.ProductId,
			Data:      data,
			TypeModel: model.PRODUCT,
		})
	}

	if err := g.db.Model(&model.File{}).Create(&newFiles).Error; err != nil {
		return nil, err
	}

	fileIds := []uint64{}

	for _, file := range newFiles {
		fileIds = append(fileIds, uint64(file.ID))
	}

	res := &proto.InsertFileRes{
		ProductId: req.ProductId,
		TypeModel: string(model.PRODUCT),
		FileIds:   fileIds,
	}

	return res, nil
}

func NewFileGRPC() proto.FileServiceServer {
	return &fileGRPC{
		db: config.GetDB(),
	}
}

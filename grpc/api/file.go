package api

import (
	"app/config"
	"app/grpc/proto"
	"app/model"
	"io"

	"gorm.io/gorm"
)

type fileGRPC struct {
	db *gorm.DB
	proto.UnsafeFileServiceServer
}

func (g *fileGRPC) InsertFile(stream proto.FileService_InsertFileServer) error {
	var newFiles []model.File

	for {
		result, err := stream.Recv()
		if err == io.EOF {
			break
		}

		newFiles = append(newFiles, model.File{
			ProductID: result.ProductId,
			Data:      result.Data,
			TypeModel: model.PRODUCT,
			Name:      result.Name,
			Format:    result.Format,
			Size:      uint64(len(result.Data)),
		})
	}

	if err := g.db.Model(&model.File{}).Create(&newFiles).Error; err != nil {
		return err
	}

	fileIds := []uint64{}

	for _, file := range newFiles {
		fileIds = append(fileIds, uint64(file.ID))
	}

	res := &proto.InsertFileRes{
		ProductId: newFiles[0].ProductID,
		TypeModel: string(model.PRODUCT),
		FileIds:   fileIds,
	}

	stream.SendAndClose(res)

	return nil
}

func NewFileGRPC() proto.FileServiceServer {
	return &fileGRPC{
		db: config.GetDB(),
	}
}

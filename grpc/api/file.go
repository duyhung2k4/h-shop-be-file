package api

import (
	"app/config"
	"app/grpc/proto"
	"app/model"
	"context"
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
		FileIds:   fileIds,
	}

	stream.SendAndClose(res)

	return nil
}

func (g *fileGRPC) DeleteFile(ctx context.Context, req *proto.DeleteFileReq) (*proto.DeleteFileRes, error) {
	listIds := req.Ids
	if len(listIds) == 0 {
		return &proto.DeleteFileRes{
			Mess: "",
		}, nil
	}
	if err := g.db.Model(&model.File{}).Unscoped().Delete(&model.File{}, listIds).Error; err != nil {
		return nil, err
	}
	return &proto.DeleteFileRes{
		Mess: "",
	}, nil
}

func (g *fileGRPC) GetFileIdsWithProductId(ctx context.Context, req *proto.GetFileIdsWithProductIdReq) (*proto.GetFileIdsWithProductIdRes, error) {
	var fileResults []GetFileIdsWithProductIdReturnQuery
	ids := []uint64{}

	if err := g.db.Model(&model.File{}).Select("id").Where("product_id = ?", req.ProductId).Scan(&fileResults).Error; err != nil {
		return nil, err
	}

	for _, file := range fileResults {
		ids = append(ids, uint64(file.Id))
	}

	return &proto.GetFileIdsWithProductIdRes{
		Ids: ids,
	}, nil
}

func (g *fileGRPC) InsertAvatarProduct(ctx context.Context, req *proto.InsertAvatarProductReq) (*proto.InsertAvatarProductRes, error) {
	var newVatar = model.File{
		Format:    req.Format,
		IsAvatar:  &model.TrueValue,
		Name:      req.Name,
		Size:      uint64(len(req.Data)),
		Data:      req.Data,
		ProductID: req.ProductId,
	}

	err := g.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.File{}).
			Where("product_id = ? AND is_avatar = ?", req.ProductId, true).Delete(&model.File{}).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.File{}).Create(&newVatar).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	res := proto.InsertAvatarProductRes{
		Id:        uint64(newVatar.ID),
		ProductId: newVatar.ProductID,
	}

	return &res, nil
}

func NewFileGRPC() proto.FileServiceServer {
	return &fileGRPC{
		db: config.GetDB(),
	}
}

type GetFileIdsWithProductIdReturnQuery struct {
	Id uint
}

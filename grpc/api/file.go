package api

import (
	"app/config"
	"app/grpc/proto"
	"context"

	"gorm.io/gorm"
)

type fileGRPC struct {
	db *gorm.DB
	proto.UnsafeFileServiceServer
}

func (g *fileGRPC) InsertFile(ctx context.Context, req *proto.InsertFileReq) (*proto.InsertFileRes, error) {
	return nil, nil
}

func NewFileGRPC() proto.FileServiceServer {
	return &fileGRPC{
		db: config.GetDB(),
	}
}

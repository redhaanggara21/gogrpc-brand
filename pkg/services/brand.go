package services

import (
	"context"
	"net/http"

	"github.com/redhaanggara21/go-grpc-brand-svc/pkg/db"
	"github.com/redhaanggara21/go-grpc-brand-svc/pkg/models"
	pb "github.com/redhaanggara21/go-grpc-brand-svc/pkg/pb"
)

type Server struct {
	H db.Handler
}

func (s *Server) CreateBrand(ctx context.Context, req *pb.CreateBrandRequest) (*pb.CreateBrandResponse, error) {
	var brand models.Brand
	brand.Brand_name = req.BrandName
	brand.Date_created = req.DateCreated

	if result := s.H.DB.Create(&brand); result.Error != nil {
		return &pb.CreateBrandResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	// fmt.Print(
	// 	&pb.CreateBrandResponse{
	// 		Status: http.StatusCreated,
	// 		Id:     1,
	// 	}
	// )

	return &pb.CreateBrandResponse{
		Status: http.StatusCreated,
		Id:     brand.Id,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var brand models.Brand

	if result := s.H.DB.First(&brand, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:          brand.Id,
		BrandName:   brand.Brand_name,
		DateCreated: brand.Date_created,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

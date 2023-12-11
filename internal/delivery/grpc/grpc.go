package grpc

import (
	"context"
	"fmt"
	"github.com/nightlord189/example-hasher/internal/config"
	"github.com/nightlord189/example-hasher/internal/entity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

type Handler struct {
	Config  config.Config
	Usecase IUsecase
}

func New(cfg config.Config, uc IUsecase) *Handler {
	return &Handler{
		Config:  cfg,
		Usecase: uc,
	}
}

func (h *Handler) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", h.Config.GRPCPort))

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	RegisterHasherServer(grpcServer, h)
	return grpcServer.Serve(listener)
}

func (h *Handler) GetHashes(ctx context.Context, request *HashRequest) (response *HashResponse, err error) {
	entities := make([]entity.HashRequestItem, 0, len(request.Items))

	for _, item := range request.Items {
		entityConverted, err := item.toEntity()
		if err != nil {
			return nil, fmt.Errorf("convert item error: %w", err)
		}
		entities = append(entities, entityConverted)
	}

	hashes, err := h.Usecase.GetHashes(ctx, entities)
	if err != nil {
		return nil, fmt.Errorf("hash error: %w", err)
	}

	convertedHashes := make([]*HashResponseItem, 0, len(hashes))
	for _, hash := range hashes {
		converted := responseItemFromEntity(&hash)
		convertedHashes = append(convertedHashes, &converted)
	}

	response = &HashResponse{
		Message: entity.SuccessResponseMessage,
		Items:   convertedHashes,
	}

	return response, nil
}

func (h *Handler) mustEmbedUnimplementedHasherServer() {

}

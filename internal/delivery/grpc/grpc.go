package grpc

import (
	"context"
	"fmt"
	"github.com/nightlord189/example-hasher/internal/entity"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

type Handler struct {
	Port    int
	Usecase IUsecase
}

func New(port int, uc IUsecase) *Handler {
	return &Handler{
		Port:    port,
		Usecase: uc,
	}
}

func (h *Handler) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", h.Port))

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	RegisterHasherServer(grpcServer, h)
	return grpcServer.Serve(listener)
}

func (h *Handler) GetHashes(ctx context.Context, request *HashRequest) (response *HashResponse, err error) {
	zerolog.Ctx(ctx).Info().Msg("new request")

	entities := make([]entity.HashRequestItem, 0, len(request.Items))

	for _, item := range request.Items {
		entityConverted, err := item.toEntity()
		if err != nil {
			zerolog.Ctx(ctx).Error().Msgf("convert item error: " + err.Error())
			return nil, fmt.Errorf("convert item error: %w", err)
		}
		entities = append(entities, entityConverted)
	}

	hashes, err := h.Usecase.GetHashes(ctx, entities)
	if err != nil {
		zerolog.Ctx(ctx).Error().Msgf("get hashes error: " + err.Error())
		return nil, fmt.Errorf("get hashes error: %w", err)
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

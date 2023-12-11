package grpc

import (
	"context"
	"github.com/nightlord189/example-hasher/internal/entity"
)

type IUsecase interface {
	GetHashes(ctx context.Context, req []entity.HashRequestItem) ([]entity.HashResponseItem, error)
}

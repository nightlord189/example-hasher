package usecase

import (
	"context"
	"fmt"
	"github.com/nightlord189/example-hasher/internal/entity"
	"github.com/nightlord189/example-hasher/pkg"
)

type Usecase struct {
}

func New() *Usecase {
	return &Usecase{}
}

func (u *Usecase) GetHashes(ctx context.Context, req []entity.HashRequestItem) ([]entity.HashResponseItem, error) {
	result := make([]entity.HashResponseItem, 0, len(req))
	for _, item := range req {
		var hashed string
		switch item.Type {
		case entity.HashTypeSHA256:
			hashed = pkg.GetSHA256Hash(item.Data)
		case entity.HashTypeSHA512:
			hashed = pkg.GetSHA512Hash(item.Data)
		default:
			return nil, fmt.Errorf("hash %s hash invalid type %s", item.ID, item.Type)
		}

		result = append(result, entity.HashResponseItem{
			ID:     item.ID,
			Result: hashed,
		})
	}

	return result, nil
}

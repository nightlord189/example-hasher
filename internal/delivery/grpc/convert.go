package grpc

import (
	"fmt"
	"github.com/nightlord189/example-hasher/internal/entity"
	"strings"
)

func (r *HashRequestItem) toEntity() (entity.HashRequestItem, error) {
	typeStr := strings.Replace(r.GetType().String(), "HASH_", "", 1)

	hashTypeParsed, err := entity.ParseHashType(typeStr)
	if err != nil {
		return entity.HashRequestItem{}, fmt.Errorf("parse hash_type error: %w", err)
	}

	return entity.HashRequestItem{
		ID:   r.GetId(),
		Type: hashTypeParsed,
		Data: r.GetData(),
	}, nil
}

func responseItemFromEntity(r *entity.HashResponseItem) HashResponseItem {
	return HashResponseItem{
		Id:     r.ID,
		Result: r.Result,
	}
}

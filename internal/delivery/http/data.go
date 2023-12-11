package http

import (
	"fmt"

	"github.com/nightlord189/example-hasher/internal/entity"
)

type HashRequest struct {
	Items []HashRequestItem `json:"items"`
}

type HashRequestItem struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func (r *HashRequestItem) toEntity() (entity.HashRequestItem, error) {
	hashTypeParsed, err := entity.ParseHashType(r.Type)
	if err != nil {
		return entity.HashRequestItem{}, fmt.Errorf("parse hash_type error: %w", err)
	}

	return entity.HashRequestItem{
		ID:   r.ID,
		Type: hashTypeParsed,
		Data: r.Data,
	}, err
}

type HashResponseItem struct {
	ID     string `json:"id"`
	Result string `json:"result"`
}

type HashResponse struct {
	Message string             `json:"message"`
	Items   []HashResponseItem `json:"items"`
}

func responseItemFromEntity(r *entity.HashResponseItem) HashResponseItem {
	return HashResponseItem{
		ID:     r.ID,
		Result: r.Result,
	}
}

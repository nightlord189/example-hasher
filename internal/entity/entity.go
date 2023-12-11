package entity

//go:generate go-enum --marshal

const SuccessResponseMessage = "success"

// HashType ENUM(SHA256, SHA512)
type HashType string

type HashRequestItem struct {
	ID   string
	Type HashType
	Data string
}

type HashResponseItem struct {
	ID     string
	Result string
}

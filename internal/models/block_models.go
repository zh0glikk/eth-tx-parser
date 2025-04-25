package models

type CreateBlockRequest struct {
	Number uint64 `json:"number"`
}

type SearchBlocksRequest struct {
	PageMetadata
}

type BlockResponse struct {
	ID     uint64 `json:"id"`
	Number uint64 `json:"number"`
}

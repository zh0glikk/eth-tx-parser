package models

type CreateAddressRequest struct {
	Address string `json:"address"`
}

type SearchAddressesRequest struct {
	PageMetadata
}

type AddressResponse struct {
	ID      uint64 `json:"id"`
	Address string `json:"address"`
}

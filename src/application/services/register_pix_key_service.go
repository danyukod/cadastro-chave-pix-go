package services

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/model/response"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database/repository"
)

type RegisterPixKeyService struct {
	pixKeyRepository repository.PixKeyRepository
}

func NewRegisterPixKeyService(
	pixKeyRepository repository.PixKeyRepository) *RegisterPixKeyService {
	return &RegisterPixKeyService{pixKeyRepository}
}

func (r *RegisterPixKeyService) Execute(request request.RegisterPixKeyRequest) (response.RegisterPixKeyResponse, error) {
	return response.RegisterPixKeyResponse{}, nil
}

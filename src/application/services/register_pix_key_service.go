package services

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/model/response"
)

type RegisterPixKeyService struct {
}

func NewRegisterPixKeyService() *RegisterPixKeyService {
	return &RegisterPixKeyService{}
}

func (r *RegisterPixKeyService) Execute(request request.RegisterPixKeyRequest) (response.RegisterPixKeyResponse, error) {
	return response.RegisterPixKeyResponse{}, nil
}

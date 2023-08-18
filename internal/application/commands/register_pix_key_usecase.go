package commands

import (
	"errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/application/commands/dto"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/model"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/aggregate"
	businesserros "github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/value_object"
	"github.com/danyukod/cadastro-chave-pix-go/internal/infrastructure/persistence"
)

type RegisterPixKeyUsecase interface {
	Execute(dto.RegisterPixKeyDTO) (model.PixKeyDomainInterface, error)
}

type RegisterPixKeyService struct {
	persistence persistence.PixKeyPersistenceInterface
}

func NewRegisterPixKeyService(
	persistence persistence.PixKeyPersistenceInterface) *RegisterPixKeyService {
	return &RegisterPixKeyService{persistence: persistence}
}

func (r *RegisterPixKeyService) Execute(dto dto.RegisterPixKeyDTO) (model.PixKeyDomainInterface, error) {
	var businessErrors businesserros.BusinessErrors

	var be *businesserros.BusinessErrors

	holderDomain, err := aggregate.NewHolderDomain(dto.AccountHolderName, dto.AccountHolderLastName)
	if checkErrors(err, be) {
		return nil, err
	}

	accoutDomain, err := aggregate.NewAccountDomain(dto.AccountNumber, dto.AgencyNumber, dto.AccountType, holderDomain)
	if checkErrors(err, be) {
		return nil, err
	}

	pixKeyDomain, err := model.NewPixKeyDomain(dto.PixKeyType, dto.PixKey, accoutDomain)
	if checkErrors(err, be) {
		return nil, err
	}

	if businessErrors.HasErrors() {
		return nil, businessErrors
	}

	err = r.VerifyIfPixKeyAlreadyExists(pixKeyDomain.GetPixKeyType().GetType(), pixKeyDomain.GetPixKey())
	if checkErrors(err, be) {
		return nil, err
	}

	if businessErrors.HasErrors() {
		return nil, businessErrors
	}

	pixKeyDomain, err = r.RegisterPixKey(pixKeyDomain)
	if err != nil {
		return nil, err
	}

	return pixKeyDomain, nil

}

func checkErrors(err error, be *businesserros.BusinessErrors) bool {
	var businessErrors businesserros.BusinessErrors

	if err == nil {
		return false
	}
	if errors.As(err, &be) {
		businessErrors = businesserros.AppendErrors(businessErrors, *be)
		return false
	}
	return true
}

func (r *RegisterPixKeyService) RegisterPixKey(pixKeyDomain model.PixKeyDomainInterface) (model.PixKeyDomainInterface, error) {
	pixKeyDomain, err := r.persistence.CreatePixKey(pixKeyDomain)
	if err != nil {
		return nil, err
	}

	return pixKeyDomain, nil
}

func (r *RegisterPixKeyService) VerifyIfPixKeyAlreadyExists(pixKeyType string, pixKey string) error {
	var businessErrors businesserros.BusinessErrors
	pixKeyDomain, err := r.persistence.FindPixKeyByKeyAndType(pixKeyType, pixKey)
	if err != nil {
		return err
	}

	if pixKeyDomain != nil {
		return businesserros.AddError(businessErrors, *businesserros.CreatePixKeyAlreadyExistsError(pixKey))
	}

	return nil
}

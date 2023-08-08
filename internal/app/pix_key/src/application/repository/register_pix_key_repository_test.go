package repository_test

import (
	"errors"
	application "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/ports/output"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/repository"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/pix_key"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/shared/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

type PixKeyPersistenceMock struct {
	output.PixKeyPersistenceInterface
	findPixKeyByKeyAndTypeFunc func(pixKeyType string, pixKey string) (pix_key.PixKeyDomainInterface, error)
	createPixKeyFunc           func(pixKeyDomain pix_key.PixKeyDomainInterface) (pix_key.PixKeyDomainInterface, error)
}

func (p PixKeyPersistenceMock) FindPixKeyByKeyAndType(pixKeyType string, pixKey string) (pix_key.PixKeyDomainInterface, error) {
	return p.findPixKeyByKeyAndTypeFunc(pixKeyType, pixKey)
}

func (p PixKeyPersistenceMock) CreatePixKey(pixKeyDomain pix_key.PixKeyDomainInterface) (pix_key.PixKeyDomainInterface, error) {
	return p.createPixKeyFunc(pixKeyDomain)
}

func TestRegisterPixKeyRepository_RegisterPixKey(t *testing.T) {
	t.Run("should return an handler when persistence layer returns an handler", func(t *testing.T) {

		pixKeyPersistenceMock := PixKeyPersistenceMock{
			findPixKeyByKeyAndTypeFunc: func(pixKeyType string, pixKey string) (pix_key.PixKeyDomainInterface, error) {
				return nil, nil
			},
			createPixKeyFunc: func(pixKeyDomain pix_key.PixKeyDomainInterface) (pix_key.PixKeyDomainInterface, error) {
				return nil, errors.New("handler creating pix key")
			},
		}

		repo := repository.NewRegisterPixKeyRepository(pixKeyPersistenceMock)

		pixKeyDomain, err := tests.PixKeyMockFactory()

		_, err = repo.RegisterPixKey(pixKeyDomain)

		assert.NotNil(t, err)
	})

	t.Run("should return a PixKeyDomain when persistence layer creates a new PixKeyDomain", func(t *testing.T) {
		pixKeyPersistenceMock := PixKeyPersistenceMock{
			findPixKeyByKeyAndTypeFunc: func(pixKeyType string, pixKey string) (pix_key.PixKeyDomainInterface, error) {
				return nil, nil
			},
			createPixKeyFunc: func(pixKeyDomain pix_key.PixKeyDomainInterface) (pix_key.PixKeyDomainInterface, error) {
				return tests.PixKeyMockFactory()
			},
		}

		repo := repository.NewRegisterPixKeyRepository(pixKeyPersistenceMock)

		pixKeyDomain, err := tests.PixKeyMockFactory()

		result, err := repo.RegisterPixKey(pixKeyDomain)

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return businessError when persistence layer find a PixKeyDomain", func(t *testing.T) {
		pixKeyPersistenceMock := PixKeyPersistenceMock{
			findPixKeyByKeyAndTypeFunc: func(pixKeyType string, pixKey string) (pix_key.PixKeyDomainInterface, error) {
				return tests.PixKeyMockFactory()
			},
			createPixKeyFunc: func(pixKeyDomain pix_key.PixKeyDomainInterface) (pix_key.PixKeyDomainInterface, error) {
				return nil, nil
			},
		}

		repo := repository.NewRegisterPixKeyRepository(pixKeyPersistenceMock)

		var businesError application.BusinessErrors

		err := repo.VerifyIfPixKeyAlreadyExists("cpf", "39357160876")
		assert.IsType(t, businesError, err)
		assert.Equal(t, "PixKey", err.(application.BusinessErrors)[0].Field())
		assert.Equal(t, "Chave pix ja cadastrada.", err.(application.BusinessErrors)[0].Error())
		assert.NotNil(t, err)
	})

	t.Run("should return error nil when persistence layer does not find a PixKeyDomain", func(t *testing.T) {
		pixKeyPersistenceMock := PixKeyPersistenceMock{
			findPixKeyByKeyAndTypeFunc: func(pixKeyType string, pixKey string) (pix_key.PixKeyDomainInterface, error) {
				return nil, nil
			},
			createPixKeyFunc: func(pixKeyDomain pix_key.PixKeyDomainInterface) (pix_key.PixKeyDomainInterface, error) {
				return nil, nil
			},
		}

		repo := repository.NewRegisterPixKeyRepository(pixKeyPersistenceMock)

		err := repo.VerifyIfPixKeyAlreadyExists("cpf", "39357160876")

		assert.Nil(t, err)
	})
}

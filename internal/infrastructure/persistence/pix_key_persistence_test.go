package persistence_test

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/infrastructure/persistence"
	"github.com/danyukod/cadastro-chave-pix-go/internal/infrastructure/persistence/entity"
	"github.com/danyukod/cadastro-chave-pix-go/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestPixKeyPersistence_CreatePixKey(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to set up test persistence: %v", err)
	}
	db.AutoMigrate(&entity.PixKeyEntity{})

	repo := persistence.NewPixKeyPersistence(db)

	pixKeyDomain, err := test.PixKeyMockFactory()
	if err != nil {
		t.Fatalf("Failed to create pixKeyDomain: %v", err)
	}

	pixKey, err := repo.CreatePixKey(pixKeyDomain)

	assert.NotNil(t, pixKey)
	assert.Equal(t, pixKeyDomain.GetPixKey(), pixKey.GetPixKey())
}

func TestPixKeyPersistence_FindPixKeyByType(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to set up test persistence: %v", err)
	}

	db.AutoMigrate(&entity.PixKeyEntity{})

	repo := persistence.NewPixKeyPersistence(db)

	pixKeyDomain, err := test.PixKeyMockFactory()
	if err != nil {
		t.Fatalf("Failed to create pixKeyDomain: %v", err)
	}

	pixKey, err := repo.CreatePixKey(pixKeyDomain)
	if err != nil {
		t.Fatalf("Failed to retrieve response: %v", err)
	}

	if pixKey.GetPixKey() != pixKeyDomain.GetPixKey() {
		t.Errorf("Expected response to be %s, got %s", pixKeyDomain.GetPixKey(), pixKey.GetPixKey())
	}

	pixKeyByType, err := repo.FindPixKeyByKeyAndType(pixKeyDomain.GetPixKeyType().GetType(), pixKey.GetPixKey())
	if err != nil {
		t.Fatalf("Failed to find pix key by type: %v", err)
	}

	assert.NotNil(t, pixKeyByType)
	assert.Equal(t, pixKeyDomain.GetPixKey(), pixKeyByType.GetPixKey())
}

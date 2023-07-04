package database_test

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database/entity"
	"github.com/danyukod/cadastro-chave-pix-go/src/shared/tests"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestPixKeyPersistence_CreatePixKey(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	db.AutoMigrate(&entity.PixKeyEntity{})

	repo := database.NewPixKeyPersistence(db)

	pixKeyDomain, err := tests.PixKeyMockFactory()
	if err != nil {
		t.Fatalf("Failed to create pixKeyDomain: %v", err)
	}

	pixKey, err := repo.CreatePixKey(pixKeyDomain)
	if err != nil {
		t.Fatalf("Failed to retrieve pixKey: %v", err)
	}

	if pixKey.GetPixKey() != pixKeyDomain.GetPixKey() {
		t.Errorf("Expected pixKey to be %s, got %s", pixKeyDomain.GetPixKey(), pixKey.GetPixKey())
	}
}

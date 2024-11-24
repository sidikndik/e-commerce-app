package repository

import (
	"e-commerce-app/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to create GORM DB: %v", err)
	}
	return gormDB, mock
}

func TestCustomerRepository_GetByCondition_GORM(t *testing.T) {
	db, mock := setupTestDB(t)
	defer func() { _ = mock.ExpectationsWereMet() }()

	customerRepo := NewCustomerRepository(db, nil)

	t.Run("successfully get customer by email", func(t *testing.T) {
		customer := model.Customer{
			Email: "johndoe@example.com",
		}

		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "password"}).
			AddRow(1, "John Doe", "johndoe@example.com", "123456789", "password123")

		mock.ExpectQuery(`SELECT \* FROM "customers" WHERE email = \$1 ORDER BY "customers"\."id" LIMIT \$2`).
			WithArgs(customer.Email, 1). // GORM secara default mengisi LIMIT dengan 1
			WillReturnRows(rows)

		result, err := customerRepo.GetByCondition(customer)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "John Doe", result.Name)
		assert.Equal(t, "johndoe@example.com", result.Email)
	})

	t.Run("successfully get customer by phone", func(t *testing.T) {
		customer := model.Customer{
			Phone: "123456789",
		}

		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "password"}).
			AddRow(1, "John Doe", "johndoe@example.com", "123456789", "password123")

		mock.ExpectQuery(`SELECT \* FROM "customers" WHERE phone = \$1 ORDER BY "customers"\."id" LIMIT \$2`).
			WithArgs(customer.Phone, 1).
			WillReturnRows(rows)

		result, err := customerRepo.GetByCondition(customer)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "John Doe", result.Name)
		assert.Equal(t, "123456789", result.Phone)
	})

	t.Run("customer not found", func(t *testing.T) {
		customer := model.Customer{
			Email: "unknown@example.com",
		}

		mock.ExpectQuery(`SELECT \* FROM "customers" WHERE email = \$1 ORDER BY "customers"\."id" LIMIT \$2`).
			WithArgs(customer.Email, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		result, err := customerRepo.GetByCondition(customer)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}

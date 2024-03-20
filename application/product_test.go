package application_test

import (
	"appproduct/application"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	// arrange
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	// act
	err := product.Enable()

	// assert
	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()

	require.Equal(t, err.Error(), "price cannot be zero")
}

func TestProduct_Disable(t *testing.T) {
	// arrange
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	// act
	err := product.Disable()

	// assert
	require.Nil(t, err)

	product.Price = 10

	err = product.Disable()

	require.Equal(t, err.Error(), "price must be zero")
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()

	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()

	require.Equal(t, err.Error(), "invalid status")

	product.Status = application.ENABLED
	_, err = product.IsValid()

	require.Nil(t, err)

	product.Price = -1
	_, err = product.IsValid()

	require.Equal(t, err.Error(), "price must be greater than zero")
}

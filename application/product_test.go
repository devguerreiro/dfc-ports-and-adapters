package application_test

import (
	"appproduct/application"
	"testing"

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

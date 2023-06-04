package bigcommerce

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func getClient() (Client, error) {
	var client Client
	err := godotenv.Load()
	if err != nil {
		return client, err
	}

	storeHash := os.Getenv("FS_STORE_HASH")
	xAuthToken := os.Getenv("FS_XAUTHTOKEN")

	client = NewClient("3", storeHash, xAuthToken)

	return client, nil
}

func TestGetProductById(t *testing.T) {
	fs, _ := getClient()

	productId := 193

	product, err := fs.getProduct(productId)

	if err != nil {
		t.Error(err)
	}

	if product.ID != productId {
		t.Error("Response-product id does not match repquest product id")
	}
}

func TestGetAllProducts(t *testing.T) {
	fs, _ := getClient()

	products, err := fs.GetAllProducts()
	if err != nil {
		t.Error(err)
	}
}

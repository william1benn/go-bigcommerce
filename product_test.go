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

	product, err := fs.GetProduct(productId)

	if err != nil {
		t.Error(err)
	}

	if product.ID != productId {
		t.Error("Response-product id does not match repquest product id")
	}
}

func TestGetAllProducts(t *testing.T) {
	fs, err := getClient()

	if err != nil {
		t.Error("error getting client")
	}

	products, _, err := fs.GetAllProducts(ProductQueryParams{})
	if err != nil {
		t.Error(err)
		return
	}

	if len(products) < 1 {
		t.Error("no products")
	}

}

func TestGetFullProductCatalog(t *testing.T) {
	fs, _ := getClient()

	products, err := fs.GetFullProductCatalog(250)
	if err != nil {
		t.Error(err)
		return
	}

	if len(products) != 69 {
		t.Error("did not fetch all products")
	}
}

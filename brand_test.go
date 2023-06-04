package bigcommerce

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func getClient() Client {
	var client Client

	godotenv.Load()

	storeHash := os.Getenv("FS_STOREHASH")
	xAuthToken := os.Getenv("FS_XAUTHTOKEN")

	client = newClient("3", storeHash, xAuthToken)

	return client

}

func TestGetBrand(t *testing.T) {
	fs := getClient()

	brandId := 1

	brand, err := fs.GetBrand(brandId)

	if err != nil {
		t.Error(err)
	}

	if brand.ID != brandId {
		t.Error("response brand-id does not match request brand id")
	}
}

func TestGetBrands(t *testing.T) {
	fs := getClient()

	brands, err := fs.GetBrands()

	if err != nil {
		t.Error(err)
	}
}

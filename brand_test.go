package bigcommerce

import (
	"testing"
)

func TestGetBrand(t *testing.T) {
	fs, _ := getClient()

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
	fs, _ := getClient()

	brands, _, err := fs.GetBrands()

	if err != nil {
		t.Error(err)
	}

	if len(brands) < 1 {
		t.Error("no brands")
	}

}

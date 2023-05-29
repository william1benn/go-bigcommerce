package bigcommerce

import (
	"fmt"
	"testing"
)

func TestGetProductById(t *testing.T) {
	fs := NewClient("", "", "")

	product, err := fs.getProductById(193)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(product)
}

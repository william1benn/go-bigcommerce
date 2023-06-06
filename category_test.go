package bigcommerce

import "testing"

func TestGetCategory(t *testing.T) {

	fs, _ := getClient()

	categoryIdDoesNotExist := 11

	_, err := fs.GetCategory(categoryIdDoesNotExist)

	if err == nil {
		t.Error("Expected Error")
	}

	categoryIdDoesExist := 27

	category, err := fs.GetCategory(categoryIdDoesExist)

	if err != nil {
		t.Error(err)
		return
	}

	if category.ID != categoryIdDoesExist {
		t.Error("reponse-category id soes not match request category id")
	}

}

func TestGetCategories(t *testing.T) {
	fs, _ := getClient()

	categories, _, err := fs.GetCategories(CategoryQueryParams{})

	if err != nil {
		t.Error(err)
	}

	if len(categories) < 1 {
		t.Error("no catgories")
	}
}

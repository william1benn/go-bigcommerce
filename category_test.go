package bigcommerce

import "testing"

func TestGetCategory(t *testing.T) {

	fs := getClient()

	categoryId := 11

	category, err := fs.GetCategory(categoryId)

	if err != nil {
		t.Error(err)
	}

	if category.ID != categoryId {
		t.Error("reponse-category id soes not match request category id")
	}

}

func TestGetCategories(t *testing.T) {
	fs := getClient()

	categories, err := fs.GetCategories()

	if err != nil {
		t.Error(err)
	}

	if len(categories) < 1 {
		t.Error("no catgories")
	}
}

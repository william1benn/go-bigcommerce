package bigcommerce

/*
	I have this idea where you select a parent product and you select some other existing products and they will be recreated as variants
	under the selected parent product. All products that became variants will be deleted and redirected

	first port of call is the ability to convert  from a product to a variant
*/
// TODO I actuall need to convert to a CreateProductVariantParams object...
func (product *Product) ToVariant(productID int) ProductVariant {
	return ProductVariant{}
}

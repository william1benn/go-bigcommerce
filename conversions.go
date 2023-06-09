package bigcommerce

/*
	I have this idea where you select a parent product and you select some other existing products and they will be recreated as variants
	under the selected parent product. All products that became variants will be deleted and redirected

	first port of call is the ability to convert  from a product to a variant
*/
// TODO I actuall need to convert to a CreateProductVariantParams object...
func (product *Product) ToVariantCreateParams(parentProductID int) ProductVariantCreateParams {
	return ProductVariantCreateParams{
		CostPrice:              product.CostPrice,
		Price:                  product.SalePrice,
		RetailPrice:            product.RetailPrice,
		Weight:                 product.Weight,
		Width:                  product.Width,
		Height:                 product.Height,
		Depth:                  product.Depth,
		IsFreeShipping:         product.IsFreeShipping,
		FixedCostShippingPrice: product.FixedCostShippingPrice,
		UPC:                    product.UPC,
		InventoryLevel:         product.InventoryLevel,
		InventoryWarningLevel:  product.InventoryWarningLevel,
		BinPickingNumber:       product.BinPickingNumber,
		GTIN:                   product.GTIN,
		MPN:                    product.MPN,
		ProductID:              parentProductID,
		SKU:                    product.SKU,
	}
}

func (product *Product) ToVariantCreateParamsWithOptions(parentProductID int, options []VariantOption) ProductVariantCreateParams {
	createVariantParams := product.ToVariantCreateParams(parentProductID)
	createVariantParams.OptionValues = &options
	return createVariantParams
}

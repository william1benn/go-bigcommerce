package bigcommerce

type Order struct {
	ID                                      int            `json:"id"`
	CustomerID                              int            `json:"customer_id"`
	DateCreated                             string         `json:"date_created"`
	DateModified                            string         `json:"date_modified"`
	DateShipped                             string         `json:"date_shipped"`
	StatusID                                int            `json:"status_id"`
	Status                                  string         `json:"status"`
	SubtotalExTax                           string         `json:"subtotal_ex_tax"`
	SubtotalIncTax                          string         `json:"subtotal_inc_tax"`
	SubtotalTax                             string         `json:"subtotal_tax"`
	BaseShippingCost                        string         `json:"base_shipping_cost"`
	ShippingCostExTax                       string         `json:"shipping_cost_ex_tax"`
	ShippingCostIncTax                      string         `json:"shipping_cost_inc_tax"`
	ShippingCostTax                         string         `json:"shipping_cost_tax"`
	ShippingCostTaxClassID                  int            `json:"shipping_cost_tax_class_id"`
	BaseHandlingCost                        string         `json:"base_handling_cost"`
	HandlingCostExTax                       string         `json:"handling_cost_ex_tax"`
	HandlingCostIncTax                      string         `json:"handling_cost_inc_tax"`
	HandlingCostTax                         string         `json:"handling_cost_tax"`
	HandlingCostTaxClassID                  int            `json:"handling_cost_tax_class_id"`
	BaseWrappingCost                        string         `json:"base_wrapping_cost"`
	WrappingCostExTax                       string         `json:"wrapping_cost_ex_tax"`
	WrappingCostIncTax                      string         `json:"wrapping_cost_inc_tax"`
	WrappingCostTax                         string         `json:"wrapping_cost_tax"`
	WrappingCostTaxClassID                  int            `json:"wrapping_cost_tax_class_id"`
	TotalExTax                              string         `json:"total_ex_tax"`
	TotalIncTax                             string         `json:"total_inc_tax"`
	TotalTax                                string         `json:"total_tax"`
	ItemsTotal                              int            `json:"items_total"`
	ItemsShipped                            int            `json:"items_shipped"`
	PaymentMethod                           string         `json:"payment_method"`
	PaymentProviderID                       string         `json:"payment_provider_id"`
	PaymentStatus                           string         `json:"payment_status"`
	RefundedAmount                          string         `json:"refunded_amount"`
	OrderIsDigital                          bool           `json:"order_is_digital"`
	StoreCreditAmount                       string         `json:"store_credit_amount"`
	GiftCertificateAmount                   string         `json:"gift_certificate_amount"`
	IPAddress                               string         `json:"ip_address"`
	IPAddressV6                             string         `json:"ip_address_v6"`
	GeoIPCountry                            string         `json:"geoip_country"`
	GeoIPCountryISO2                        string         `json:"geoip_country_iso2"`
	CurrencyID                              int            `json:"currency_id"`
	CurrencyCode                            string         `json:"currency_code"`
	CurrencyExchangeRate                    string         `json:"currency_exchange_rate"`
	DefaultCurrencyID                       int            `json:"default_currency_id"`
	DefaultCurrencyCode                     string         `json:"default_currency_code"`
	StaffNotes                              string         `json:"staff_notes"`
	CustomerMessage                         string         `json:"customer_message"`
	DiscountAmount                          string         `json:"discount_amount"`
	CouponDiscount                          string         `json:"coupon_discount"`
	ShippingAddressCount                    int            `json:"shipping_address_count"`
	IsDeleted                               bool           `json:"is_deleted"`
	EbayOrderID                             string         `json:"ebay_order_id"`
	CartID                                  string         `json:"cart_id"`
	BillingAddress                          BillingAddress `json:"billing_address"`
	IsEmailOptIn                            bool           `json:"is_email_opt_in"`
	CreditCardType                          interface{}    `json:"credit_card_type"`
	OrderSource                             string         `json:"order_source"`
	ChannelID                               int            `json:"channel_id"`
	ExternalSource                          interface{}    `json:"external_source"`
	Products                                URLResource    `json:"products"`
	ShippingAddresses                       URLResource    `json:"shipping_addresses"`
	Coupons                                 URLResource    `json:"coupons"`
	ExternalID                              interface{}    `json:"external_id"`
	ExternalMerchantID                      interface{}    `json:"external_merchant_id"`
	TaxProviderID                           string         `json:"tax_provider_id"`
	StoreDefaultCurrencyCode                string         `json:"store_default_currency_code"`
	StoreDefaultToTransactionalExchangeRate string         `json:"store_default_to_transactional_exchange_rate"`
	CustomStatus                            string         `json:"custom_status"`
	CustomerLocale                          string         `json:"customer_locale"`
	ExternalOrderID                         string         `json:"external_order_id"`
}

type BillingAddress struct {
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Company     string       `json:"company"`
	Street1     string       `json:"street_1"`
	Street2     string       `json:"street_2"`
	City        string       `json:"city"`
	State       string       `json:"state"`
	Zip         int          `json:"zip"`
	Country     string       `json:"country"`
	CountryISO2 string       `json:"country_iso2"`
	Phone       int          `json:"phone"`
	Email       string       `json:"email"`
	FormFields  []FormFields `json:"form_fields"`
}

type FormFields struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type URLResource struct {
	URL      string `json:"url"`
	Resource string `json:"resource"`
}

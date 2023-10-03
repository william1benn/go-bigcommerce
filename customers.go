// https://developer.bigcommerce.com/docs/rest-management/customers#create-customers
package bigcommerce

import (
	"encoding/json"
	"fmt"
	"strings"
)

type BcAuthentication struct {
	ForcePasswordReset bool   `json:"force_password_reset,omitempty"`
	NewPassword        string `json:"new_password,omitempty"`
}

type BcAddresses struct {
	Address1        string         `json:"address1,omitempty"`
	Address2        string         `json:"address2,omitempty"`
	AddressType     string         `json:"address_type,omitempty"`
	City            string         `json:"city,omitempty"`
	Company         string         `json:"company,omitempty"`
	CountryCode     string         `json:"country_code,omitempty"`
	FirstName       string         `json:"first_name,omitempty"`
	LastName        string         `json:"last_name,omitempty"`
	Phone           string         `json:"phone,omitempty"`
	PostalCode      string         `json:"postal_code,omitempty"`
	StateOrProvince string         `json:"state_or_province,omitempty"`
	FormFields      []BcFormFields `json:"form_fields,omitempty"`
}

type BcFormFields struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type BcAttributes struct {
	Id             int    `json:"id,omitempty"`
	CustomerId     int    `json:"customer_id,omitempty"`
	AttributeId    int    `json:"attribute_id,omitempty"`
	AttributeValue string `json:"attribute_value,omitempty"`
	DateCreated    string `json:"date_created,omitempty"`
	DateModified   string `json:"date_modified,omitempty"`
}

type BcAmount struct {
	Amount float64 `json:"amount,omitempty"`
}

type BcCustomer struct {
	Id                                      int              `json:"id"`
	Email                                   string           `json:"email"`
	FirstName                               string           `json:"first_name"`
	LastName                                string           `json:"last_name"`
	Company                                 string           `json:"company,omitempty"`
	Phone                                   string           `json:"phone,omitempty"`
	Notes                                   string           `json:"notes,omitempty"`
	TaxExemptCategory                       string           `json:"tax_exempt_category,omitempty"`
	CustomerGroupID                         int              `json:"customer_group_id,omitempty"`
	Authentication                          BcAuthentication `json:"authentication,omitempty"`
	Addresses                               []BcAddresses    `json:"addresses,omitempty"`
	AcceptsProductReviewAbandonedCartEmails bool             `json:"accepts_product_review_abandoned_cart_emails,omitempty"`
	StoreCreditAmounts                      []BcAmount       `json:"store_credit_amounts,omitempty"`
	OriginChannelID                         int              `json:"origin_channel_id,omitempty"`
	ChannelIds                              []int            `json:"channel_ids,omitempty"`
	FormFields                              []BcFormFields   `json:"form_fields,omitempty"`
	Attributes                              []BcAttributes   `json:"attributes,omitempty"`
	Meta                                    MetaData         `json:"meta,omitempty"`
}

type BcCustomerParams struct {
	IDIn            []int    `url:"id:in,omitempty,comma"`
	CompanyIn       []string `url:"company:in,omitempty,comma"`
	IDGreater       []int    `url:"id:greater,omitempty,comma"`
	NameIn          []string `url:"name:in,omitempty,comma"`
	EmailIn         []string `url:"email:in,omitempty,comma"`
	NameLike        []string `url:"name:like,omitempty"`
	Include         string   `url:"include,omitempty"`
	Page            int      `url:"page,omitempty"`
	Limit           int      `url:"limit,omitempty"`
	Sort            string   `url:"sort,omitempty"`
	CategoriesIn    []int    `url:"categories:in,omitempty,comma"`
	CustomerGroupIn []string `url:"customer_group_id:in,omitempty,comma"`
	DateCreated     string   `url:"date_created,omitempty,comma"`
	DateCreatedMin  string   `url:"date_created:max,omitempty,comma"`
	DateCreatedMax  string   `url:"date_created:min,omitempty,comma"`
	DateModified    string   `url:"date_modified,omitempty,comma"`
}

// GET /stores/{store_hash}/v3/customers
func (client *Client) CreateCustomer(params []BcCustomer) ([]BcCustomer, error) {
	type ResponseObject struct {
		Data []BcCustomer `json:"data"`
		Meta MetaData     `json:"meta"`
	}
	var CustomerResponse ResponseObject

	for _, items := range params {
		if strings.TrimSpace(items.Email) == "" || strings.TrimSpace(items.FirstName) == "" || strings.TrimSpace(items.LastName) == "" {
			return CustomerResponse.Data, fmt.Errorf("you are missing a required parameter check (email,firstname,lastname)")
		}
	}

	p, err := json.Marshal(params)
	if err != nil {
		return CustomerResponse.Data, err
	}

	path := client.BaseURL.JoinPath("customers").String()
	resp, _ := client.Post(path, p)

	err = expectStatusCode(200, resp)
	if err != nil {
		return CustomerResponse.Data, err
	}

	err = json.NewDecoder(resp.Body).Decode(&CustomerResponse)
	if err != nil {
		return CustomerResponse.Data, err
	}
	return CustomerResponse.Data, nil
}

// GET /stores/{store_hash}/v3/customers
func (client *Client) GetCustomers(queryParams BcCustomerParams) ([]BcCustomer, error) {
	type ResponseObject struct {
		Data []BcCustomer `json:"data"`
		Meta MetaData     `json:"meta"`
	}
	var CustomerResponse ResponseObject

	params, err := paramString(queryParams)
	if err != nil {
		return CustomerResponse.Data, fmt.Errorf(err.Error())
	}

	getCustomers := client.BaseURL.JoinPath("customers").String() + params

	resp, _ := client.Get(getCustomers)
	err = expectStatusCode(200, resp)
	if err != nil {
		return CustomerResponse.Data, fmt.Errorf(err.Error())
	}

	err = json.NewDecoder(resp.Body).Decode(&CustomerResponse)
	if err != nil {
		return CustomerResponse.Data, fmt.Errorf(err.Error())
	}

	return CustomerResponse.Data, nil
}

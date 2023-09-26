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

type BcFormFields []struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type BcAmount struct {
	Amount float64 `json:"amount,omitempty"`
}

type BcCustomer []struct {
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
	Meta                                    MetaData         `json:"meta,omitempty"`
}

// /stores/{store_hash}/v3/customers
func (client *Client) createCustomer(params BcCustomer) (BcCustomer, error) {
	type ResponseObject struct {
		Data BcCustomer `json:"data"`
		Meta MetaData   `json:"meta"`
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

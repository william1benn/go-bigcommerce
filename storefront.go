package bigcommerce

import (
	"encoding/json"
	"errors"
	"time"
)

type TokenInput struct {
	AllowedCorsOrigins []string `json:"allowed_cors_origins"`
	ChannelId          *int     `json:"channel_id"`
	ExpiresAt          *int64   `json:"expires_at"`
}

type StoreTokenResponse struct {
	Data Token    `json:"data"`
	Meta MetaData `json:"meta"`
}

type Token struct {
	JwtToken string `json:"token"`
}

// POST stores/{store_hash}/v3/storefront/api-token
func (client *Client) CreateToken(allowedCorsOrigins []string, channelId *int, Expire *int64) (Token, MetaData, error) {
	response := StoreTokenResponse{}
	if len(allowedCorsOrigins) > 2 || len(allowedCorsOrigins) == 0 {
		return response.Data, response.Meta, errors.New("you cannot have more than two urls or zero urls for allowed cors - required")
	}
	if Expire == nil {
		timeAdded := time.Now().Add(time.Hour * 24).Unix()
		Expire = &timeAdded
	}
	if channelId == nil {
		channelNum := 1
		channelId = &channelNum
	}

	tokenInput := &TokenInput{
		AllowedCorsOrigins: allowedCorsOrigins,
		ExpiresAt:          Expire,
		ChannelId:          channelId,
	}

	path := client.BaseURL.JoinPath("/storefront/api-token").String()
	tokenMarshal, err := json.Marshal(tokenInput)
	if err != nil {
		return response.Data, response.Meta, err
	}

	resp, _ := client.Post(path, tokenMarshal)
	if err != nil {
		return response.Data, response.Meta, err
	}

	err = expectStatusCode(200, resp)
	if err != nil {
		return response.Data, response.Meta, err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Data, response.Meta, err
	}
	return response.Data, response.Meta, nil
}

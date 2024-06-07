package bigcommerce

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func makeClient() (Client, error) {
	var client Client
	err := godotenv.Load()
	if err != nil {
		return client, err
	}

	storeHash := os.Getenv("FS_STORE_HASH")
	xAuthToken := os.Getenv("FS_XAUTHTOKEN")

	client = NewClient("3", storeHash, xAuthToken)

	return client, nil
}

func Test_GenerateStoretoken(t *testing.T) {
	c, _ := makeClient()
	channel := 3
	tokenRes, _, err := c.CreateToken([]string{"https://example.com"}, &channel, nil)
	if err != nil {
		t.Fail()
	}
	fmt.Println(tokenRes.JwtToken)
}

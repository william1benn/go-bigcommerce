package bigcommerce

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func makeClient(t *testing.T) (Client, error) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err.Error())
	}

	storeHash := os.Getenv("FS_STORE_HASH")
	xAuthToken := os.Getenv("FS_XAUTHTOKEN")

	client := NewClient("3", storeHash, xAuthToken)
	return client, nil
}

func Test_GenerateStoretoken(t *testing.T) {
	c, err := makeClient(t)
	if err != nil {
		t.Fatal(err.Error())
	}

	tokenRes, err := c.CreateToken([]string{"https://example.com"}, nil, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(tokenRes.JwtToken)
}

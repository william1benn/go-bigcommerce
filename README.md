This library provides an unofficial Go client for the bigcommerce REST API

### Installation:
```
go get github.com/seanomeara96/go-bigcommerce
```

### bigcommerce example usage:

```go
package main

import (
	"fmt"

  	"github.com/joho/godotenv"
  	bigcommerce "github.com/seanomeara96/go-bigcommerce"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	storeHash := os.Getenv("STORE_HASH")
	xAuthToken := os.Getenv("XAUTHTOKEN")

	store := bigcommerce.NewClient("3", storeHash, xAuthToken)

	products, err := store.GetFullProductCatalog(250)
	if err != nil {
		panic(err)
	}
	fmt.Println(products)
}

```

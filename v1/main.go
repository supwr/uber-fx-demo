package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

var httpClient *resty.Client

func init() {
	httpClient = resty.New()
}

func main() {
	ctx := context.Background()

	external := External{
		HttpClient: httpClient,
	}

	persons, _ := external.getRandomPersons(ctx)
	result, _ := json.Marshal(persons)

	fmt.Println(string(result))
}

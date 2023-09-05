package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/supwr/uber-fx-demo/entities"
	"go.uber.org/fx"
)

type ExternalRepository interface {
	getRandomPersons(ctx context.Context) ([]entities.PersonDTO, error)
}

func main() {
	ctx := context.Background()

	app := createApp(
		fx.Invoke(func(external ExternalRepository) {
			persons, _ := external.getRandomPersons(ctx)
			result, _ := json.Marshal(persons)

			fmt.Println(string(result))
		}),
		fx.Invoke(func(s fx.Shutdowner) { _ = s.Shutdown() }),
	)

	app.Run()
}

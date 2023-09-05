package main

import (
	"context"
	"encoding/json"
	"fmt"

	"go.uber.org/fx"
)

func main() {
	ctx := context.Background()

	app := createApp(
		fx.Invoke(func(external *External) {
			persons, _ := external.getRandomPersons(ctx)
			result, _ := json.Marshal(persons)

			fmt.Println(string(result))
		}),
		fx.Invoke(func(s fx.Shutdowner) { _ = s.Shutdown() }),
	)

	app.Run()
}

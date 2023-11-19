package requests

import (
	"context"
	"os"

	"github.com/machinebox/graphql"
)

// type Request interface {
// 	DoRequest()
// }

func DoRequest(query string, client *graphql.Client, response string) {
	request := graphql.NewRequest(query)
	request.Header.Set("Authorization", "Bearer "+os.Getenv("CAIDO_AUTH_TOKEN"))

	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}
}

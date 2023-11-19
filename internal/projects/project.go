package projects

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/machinebox/graphql"
)

func GetProjectInfo(client *graphql.Client) {
	query := `
	query {
		projects{
			id
			name
			version
			updatedAt
		}
	}
	`
	request := graphql.NewRequest(query)
	request.Header.Set("Authorization", "Bearer "+os.Getenv("CAIDO_AUTH_TOKEN"))

	var response QueryProjectsResponse
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 71; i++ {
		fmt.Printf("%s", "-")
	}
	fmt.Printf("\n|%5s |%10s |%10s |%15s |%20s |\n", "Idx", "Version", "Size", "Updated At", "Name")
	for i := 0; i < 71; i++ {
		fmt.Printf("%s", "-")
	}
	fmt.Println()

	for index, el := range response.Projects {
		fmt.Printf("|%5s |%10s |%10d |%15s |%20s |\n", strconv.Itoa(index), "v"+el.Version, el.Size, el.UpdatedAt.Format("2006-01-02"), el.Name)
	}

	var choose string
	fmt.Printf("\n > Choose a Project or Nothing to work on current Project : ")
	fmt.Scanln(&choose)
}

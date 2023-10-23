package main

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

	fmt.Println("Idx ", "Ver.  ", "ProjectName")
	for index, el := range response.Projects {
		fmt.Println("["+strconv.Itoa(index)+"] ", el.Version, el.Name)
	}
}

type QueryProjectsResponse struct {
	Projects []struct {
		Id        string
		Name      string
		Version   string
		UpdatedAt string
	}
}

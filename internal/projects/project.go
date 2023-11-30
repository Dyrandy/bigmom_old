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

	getCurrentProject(client)

	// var choose string
	// fmt.Printf("\n > Choose a Project or Nothing to work on current Project : ")
	// fmt.Scanln(&choose)

}

func getCurrentProject(client *graphql.Client) {
	query := `
	query {
		currentProject{
			id
			name
			version
			updatedAt
		}
	}
	`
	request := graphql.NewRequest(query)
	request.Header.Set("Authorization", "Bearer "+os.Getenv("CAIDO_AUTH_TOKEN"))

	var response QueryCurrentProjectResponse
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n > Curent Project : [ %s ]\n", response.CurrentProject.Name)
}

func InitProject(client *graphql.Client) {
	query1 := `
	mutation {
		createTamperRule(input: {
	  name: "If-None-Match",
	  collectionId: "1",
	  strategy: "REQUEST_HEADER",
	  matchTerm: "If-None-Match: .*",
	  replaceTerm: "",
	  isRegex: true,
	  isEnabled: true
	}){
	__typename
	}
	}
	`

	query2 := `
	mutation {
		createTamperRule(input: {
	  name: "If-Modified-Since",
	  collectionId: "1",
	  strategy: "REQUEST_HEADER",
	  matchTerm: "If-Modified-Since: .*",
	  replaceTerm: "",
	  isRegex: true,
	  isEnabled: true
	}){
	__typename
	}
	}
	`

	request1 := graphql.NewRequest(query1)
	request1.Header.Set("Authorization", "Bearer "+os.Getenv("CAIDO_AUTH_TOKEN"))
	request2 := graphql.NewRequest(query2)
	request2.Header.Set("Authorization", "Bearer "+os.Getenv("CAIDO_AUTH_TOKEN"))

	var response CreateTamperRuleResponse
	err1 := client.Run(context.Background(), request1, &response)
	err2 := client.Run(context.Background(), request2, &response)
	if err1 != nil && err2 != nil {
		panic(err1)
	}

	getProjectSettings(client)
}

func getProjectSettings(client *graphql.Client) {
	query := `
	query {
		tamperRuleCollection(id: "1"){
      rules{
        name
        strategy
        isRegex
        matchTerm
        replaceTerm
  		isEnabled
      }
    }
	}
	`

	request := graphql.NewRequest(query)
	request.Header.Set("Authorization", "Bearer "+os.Getenv("CAIDO_AUTH_TOKEN"))

	var response TamperRuleCollectionReponse
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 108; i++ {
		fmt.Printf("%s", "-")
	}
	fmt.Printf("\n|%15s |%20s |%25s |%15s |%10s |%10s |\n", "Strategy", "Name", "MatchTerm", "ReplaceTerm", "isRegex", "isEnabled")
	for i := 0; i < 108; i++ {
		fmt.Printf("%s", "-")
	}
	fmt.Println()

	for _, el := range response.Data.Rules {
		fmt.Printf("|%15s |%20s |%25s |%15s |%10t |%10t |\n", el.Strategy, el.Name, el.MatchTerm, el.ReplaceTerm, el.IsRegex, el.IsEnabled)
	}
	fmt.Println()
	fmt.Println(" > Setting Complete")
}

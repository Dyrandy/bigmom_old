package replays

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"

	"github.com/machinebox/graphql"
)

func GetProjectReplaySessions(client *graphql.Client) {
	query := `
	query{
		replaySessions{
		  nodes{
			name
		  }
		}
	  }
	`
	request := graphql.NewRequest(query)
	request.Header.Set("Authorization", "Bearer "+os.Getenv("CAIDO_AUTH_TOKEN"))

	var response QueryProjectSessionsResponse
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}
	for index, el := range response.ReplaySessions.Nodes {
		fmt.Println("["+strconv.Itoa(index+1)+"] ", el.Name)
	}
}

func GetReplaySessionData(client *graphql.Client, id string) {
	query := `
	query{
		replaySession(id: "` + id + `"){
		  name
		  entries{
			nodes{
			  request{
				host
				path
				query
				raw
			  }
			}
		  }
		}
	  }
	`
	request := graphql.NewRequest(query)
	request.Header.Set("Authorization", "Bearer "+os.Getenv("CAIDO_AUTH_TOKEN"))

	var response QueryProjectReplayDataResponse
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}

	for index, el := range response.ReplaySession.Entries.Nodes {
		// fmt.Println(index, el)
		base64Test, err2 := base64.StdEncoding.DecodeString(el.Request.Raw)
		if err2 != nil {
			panic(err2)
		}
		fmt.Printf("\n[%d]  Request Session\n%s\n", index, base64Test)
		// fmt.Println(index, base64Test)
	}
}

type QueryProjectSessionsResponse struct {
	ReplaySessions struct {
		Nodes []struct {
			Name string
		}
	}
}

type QueryProjectReplayDataResponse struct {
	ReplaySession struct {
		Name    string
		Entries struct {
			Nodes []struct {
				Request struct {
					Raw string
				}
			}
		}
	}
}

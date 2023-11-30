package replays

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/Dyrandy/bigmom/internal/menus"
	"github.com/Dyrandy/bigmom/internal/scanner"
	"github.com/machinebox/graphql"
)

func GetProjectReplaySessions(client *graphql.Client) {
	query := `
	query {
		replaySessionCollections{
			nodes{
        name
        sessions{
          name
          id 
        } 
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
	for _, el := range response.ReplaySessionCollections.Nodes {
		fmt.Println("[*]", el.Name)
		printBar(236)
		fmt.Printf("\n|%4s |%12s |%8s |%40s |%5s |%77s |%70s |\n", "Idx", "Session Name", "Method", "Host", "Port", "Path", "Query")
		printBar(236)
		for _, el2 := range el.Sessions {
			getReplaySessionData(client, el2.ID, el2.Name, false)
			// fmt.Println(el2.ID, el2.Name)
		}
	}
	var choose string
	fmt.Printf("\n\n > Which Session ID to Attack : ")
	fmt.Scanln(&choose)
	getReplaySessionData(client, choose, "Nothing", true)
}

func getReplaySessionData(client *graphql.Client, id string, sessionName string, isForTarget bool) {
	query := `
	query{
		replaySession(id: "` + id + `"){
			name
			activeEntry{
			  request{
				host
				path
				method
				query
				raw
				isTls
				fileExtension
				source
				port
		  
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
	el := response.ReplaySession.ActiveEntry
	if isForTarget != true {
		fmt.Printf("\n|%4s |%12s |%8s |%40s |%5d |%77s |%70s |", id, sessionName, el.Request.Method, el.Request.Host, el.Request.Port, el.Request.Path, el.Request.Query)
	} else {
		base64Test, err2 := base64.StdEncoding.DecodeString(el.Request.Raw)
		if err2 != nil {
			panic(err2)
		}
		menus.Logo()
		fmt.Printf("[*] %s\n", response.ReplaySession.Name)
		printBar(130)
		fmt.Printf("\n%s\n", base64Test)
		printBar(130)
		fmt.Println()
		menus.PrintAttackMenu()

		var choose string
		fmt.Printf("\n > Which Attack to Perform : ")
		fmt.Scanln(&choose)

		switch strings.ToUpper(choose) {
		case "Q":
			return
		case "W":
			fmt.Println("Not Implemented Yet")
			return
		case "E":
			fmt.Println("Not Implemented Yet")
			return
		case "R":
			scanner.DoRaceAttack(el.Request.Raw, el.Request.Method, el.Request.Host, el.Request.Path, el.Request.Query, el.Request.IsTLS)
		default:
			fmt.Println("Wrong Choice")
			return
		}
	}

	// for _, el := range response.ReplaySession.ActiveEntry.Request {
	// fmt.Println(index, el)
	// base64Test, err2 := base64.StdEncoding.DecodeString(el.Request.Raw)
	// if err2 != nil {
	// 	panic(err2)
	// }
	// fmt.Printf("\n[%d]  Request Session\n%s\n", index, base64Test)

	// fmt.Println(el.Request.Host, el.Request.Path, el.Request.Query)
	// }
	// fmt.Println()
	// printBar()
	// fmt.Println()
}

func printBar(length int) {
	for i := 0; i < length; i++ {
		fmt.Printf("%s", "-")
	}
}

package main

import (
	"fmt"
	"strings"

	"github.com/Dyrandy/bigmom/internal/menus"
	"github.com/Dyrandy/bigmom/internal/projects"
	"github.com/Dyrandy/bigmom/internal/replays"
	"github.com/machinebox/graphql"
)

var graphqlURL string = "http://127.0.0.1:8080/graphql"
var client *graphql.Client

func init() {
	client = graphql.NewClient(graphqlURL)
	menus.Logo()
}

func main() {
	fmt.Println("[*] Projects")
	projects.GetProjectInfo(client)
	menus.PrintMenu()

	var choose string

	for {
		fmt.Printf("\n > Input Choice : ")
		fmt.Scanln(&choose)

		switch strings.ToUpper(choose) {
		case "Q": // Quit
			return
		case "W": // Init Project
			menus.Logo()
			projects.InitProject(client)
			menus.PrintMenu()
		case "E": // Fuzz Target
			fmt.Println("Not Defined Yet")
			return
		case "R": // Vuln Testing
			menus.Logo()
			// fmt.Println("Not Defined Yet")
			replays.GetProjectReplaySessions(client)
			menus.PrintMenu()
			return
		default:
			menus.Logo()
			fmt.Println("Not a Valid Choice")
			menus.PrintMenu()
		}
	}
	// fmt.Println("\n[*] Sessions for Current Project")
	// replays.GetProjectReplaySessions(client)

	// var choose string
	// fmt.Print("[*] Which Session Do You Want?: ")
	// fmt.Scanln(&choose)
	// replays.GetReplaySessionData(client, choose)
}

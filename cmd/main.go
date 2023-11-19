package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Dyrandy/bigmom/internal/projects"
	_ "github.com/Dyrandy/bigmom/internal/replays"
	"github.com/machinebox/graphql"
)

var graphqlURL string = "http://127.0.0.1:8080/graphql"
var client *graphql.Client

func init() {
	client = graphql.NewClient(graphqlURL)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	printLogo()
}

func main() {
	fmt.Println("[*] Projects")
	projects.GetProjectInfo(client)

	// fmt.Println("\n[*] Sessions for Current Project")
	// replays.GetProjectReplaySessions(client)

	// var choose string
	// fmt.Print("[*] Which Session Do You Want?: ")
	// fmt.Scanln(&choose)
	// replays.GetReplaySessionData(client, choose)
}

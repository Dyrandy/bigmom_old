package main

import (
	"fmt"

	"github.com/machinebox/graphql"
)

func main() {
	PrintLogo()
	client := graphql.NewClient("http://127.0.0.1:8080/graphql")
	fmt.Println("[*] Projects")
	GetProjectInfo(client)

	fmt.Println("\n[*] Sessions for Current Project")
	GetProjectReplaySessions(client)

	var choose string
	fmt.Print("[*] Which Session Do You Want?: ")
	fmt.Scanln(&choose)
	GetReplaySessionData(client, choose)
}

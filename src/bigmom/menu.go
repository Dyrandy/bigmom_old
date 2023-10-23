package main

import "fmt"

func PrintLogo() {
	logo := `
   Welcome To BigMom, Caido-GraphQL-Client
      ____  _       __  ___              
     / __ )(_)___ _/  |/  /___  ____ ___ 
    / __  / / __ '/ /|_/ / __ \/ __ '__ \
   / /_/ / / /_/ / /  / / /_/ / / / / / /
  /_____/_/\__, /_/  /_/\____/_/ /_/ /_/ 
          /____/                            v1.0
	
	By: Dyrandy
		  `
	fmt.Println(logo)
}

func PrintMenu() {
	fmt.Println("[0] Get Project Information")
	fmt.Println("[1] Use Project")
}

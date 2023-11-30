package menus

import (
	"fmt"
	"os"
	"os/exec"
)

func printLogo() {
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
	fmt.Println("[Q] Quit    [W] Init Settings    [E] Fuzz Target     [R] Vuln Testing")
}

func PrintAttackMenu() {
	fmt.Println("[Q] Quit    [W] XSS    [E] Brute Force     [R] Race Condition")
}

func Logo() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	printLogo()
}

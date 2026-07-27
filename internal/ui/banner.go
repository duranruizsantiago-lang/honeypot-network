/*
banner.go

ASCII art banner for the hive honeypot network CLI
*/

package ui

import "fmt"

func PrintBanner() {
	art := `
  ██╗  ██╗██╗██╗   ██╗███████╗
  ██║  ██║██║██║   ██║██╔════╝
  ███████║██║██║   ██║█████╗
  ██╔══██║██║╚██╗ ██╔╝██╔══╝
  ██║  ██║██║ ╚████╔╝ ███████╗
  ╚═╝  ╚═╝╚═╝  ╚═══╝  ╚══════╝`

	fmt.Println(Cyan(art))
	fmt.Printf(
		"  %s %s\n\n",
		Dim("honeypot network"),
		Dim("v"+Version),
	)
}

const Version = "0.1.0"

package main

import (
	"fmt"
	"os"

	"github.com/KuiK-san/cli-golang/internal/clis/docker"
	"github.com/KuiK-san/cli-golang/internal/clis/git"
	"github.com/KuiK-san/cli-golang/internal/clis/monitor"
)

const VERSION string = "1.0"

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("Nenhum argumento encontrado")
		return
	}

	switch os.Args[1] {
	case "-v", "--version":
		fmt.Println("cli-server on version:", VERSION)
	case "git":
		git.Cli(os.Args[2:])
	case "container", "docker":
		docker.Cli(os.Args[2:])
	case "monitor":
		monitor.Cli(os.Args[2:])
	case "--help", "-h":
		helpDefault()
	default:
		fmt.Print(os.Args[1:], " não encontrado")
	}
}

func helpDefault() {
	fmt.Println("Você entrou no help")
}

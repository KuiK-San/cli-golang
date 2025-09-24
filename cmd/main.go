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
		fmt.Println("Nenhum comando encontrado")
		helpDefault()
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
		fmt.Print(os.Args[1:], "Não é um comando valido")
		helpDefault()
	}
}

func helpDefault() {
	fmt.Println("")
	fmt.Println(`CLI-Server - Ferramenta de linha de comando multifuncional

	USO:
	cli-server [COMANDO] [SUBCOMANDOS] [OPÇÕES]

	COMANDOS PRINCIPAIS:
	git         Operações de configuração do git
	docker      Operações de gerenciamento de containers
	monitor     Operações de monitoramento de servicos

	OPÇÕES GLOBAIS:
	-h, --help     Mostra esta mensagem de ajuda
	-v, --version  Mostra a versão da aplicação `,
	)
	fmt.Println("")
}

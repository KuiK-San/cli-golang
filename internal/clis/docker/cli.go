package docker

import (
	"fmt"
)

func Cli(params []string) {
	if len(params[1:]) < 1 {
		fmt.Println("Nenhum comando encontrado")
		helpDefault()
		return
	}

	switch params[1] {
	case "restart":
		restartContainer(params[2])
	case "restart-push":
		restartContainer(params[2])
	case "--help", "-h":
		helpDefault()
	default:
		fmt.Print(params[1:], "Não é um comando valido")
		helpDefault()
	}
}

func helpDefault() {
	fmt.Println("")
	fmt.Println(`CLI-Server - Ferramenta de linha de comando multifuncional

	USO:
	cli-server docker [OPÇÕES]

	COMANDOS PRINCIPAIS:
	restart   		Reinicia e rebuilda o container
	restart-push	Executa o push do repositório, reinicia e rebuilda o container
	
	OPÇÕES GLOBAIS:
	-h, --help     Mostra esta mensagem de ajuda`,
	)
	fmt.Println("")
}

func restartContainer(container string) {
	// TODO: fazer o restart do container
}

func pushRepo() {
	// TODO: fazer o pull do repositorio
}

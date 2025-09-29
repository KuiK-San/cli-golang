package git

import (
	"fmt"
	"os"
)

func Cli(params []string) {
	if len(os.Args[1:]) < 1 {
		fmt.Println("Nenhum comando encontrado")
		helpDefault()
		return
	}

	switch os.Args[1] {
	case "configure":
		// TODO: escreve eval "$(ssh-agent -s)" e ssh-add ~/.ssh/github no terminal atual do usuario
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
	cli-server git [SUBCOMANDOS] [OPÇÕES]

	COMANDOS PRINCIPAIS:
	configure   Configura o git na pasta atual
	
	OPÇÕES GLOBAIS:
	-h, --help     Mostra esta mensagem de ajuda`,
	)
	fmt.Println("")
}

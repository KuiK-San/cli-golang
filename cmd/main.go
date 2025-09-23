package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		return
	}

	if len(os.Args[1:]) < 1 {
		fmt.Println("Nenhum argumento encontrado")
		return
	}

	switch os.Args[1] {
	case "-v", "--version":
		fmt.Println("cli-server on version:", os.Getenv("VERSION"))
	default:
		fmt.Print(os.Args[1:], " não encontrado")
	}
}

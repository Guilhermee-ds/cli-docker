package cmd

import (
	"log"
	"os"
	"strings"
    "github.com/kyokomi/emoji"
)

// Função para gerar o Dockerfile com base nas escolhas do usuário
func generateDockerfile(language string) {
    file, err := os.Create("Dockerfile")
    if err != nil {
        emoji.Println(":x: Erro ao criar o Dockerfile:", err)
        return
    }
    defer file.Close()

    // Criação do conteúdo do Dockerfile
    dockerfileContent := getDockerfileContent(language)

    // Escreve o conteúdo no Dockerfile
    _, err = file.WriteString(dockerfileContent)
    if err != nil {
        emoji.Println(":x: Erro ao escrever no Dockerfile:", err)
        return
    }

    emoji.Println(":page_facing_up: :white_check_mark: Dockerfile gerado com sucesso!")
}


func readDockerfile(filename string) string{
    content, err := os.ReadFile(filename)

    if err != nil {
        log.Fatalf("Erro ao ler o Dockerfile: %v", err)
    }
    return string(content)
}

// Função para obter o conteúdo do Dockerfile dependendo da linguagem escolhida
func getDockerfileContent(language string) string {
    switch strings.ToLower(language) {
    case "go":
        return readDockerfile("dockerfiles/go.Dockerfile")
    case "python":
        return readDockerfile("dockerfiles/python.Dockerfile")
    case "node.js":
        return  readDockerfile("dockerfiles/node.Dockerfile")

    default:
        return "FROM ubuntu:latest\nCMD [\"bash\"]\n"
    }
}


package cmd

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/spf13/cobra"
    "github.com/kyokomi/emoji"
)

var rootCmd = &cobra.Command{
    Use:   "docker-cli-app",
    Short: "Uma CLI para gerar Dockerfiles e construir imagens Docker",
    Run: func(cmd *cobra.Command, args []string) {
        // Lógica de interação com o usuário aqui
        reader := bufio.NewReader(os.Stdin)

        // Pergunta sobre a linguagem
        emoji.Print(":question: Deseja incluir uma linguagem de programação? (sim/não): ")
        languageAnswer, _ := reader.ReadString('\n')
        languageAnswer = strings.TrimSpace(languageAnswer)

        var language string
        if strings.ToLower(languageAnswer) == "sim" {
            fmt.Println("Escolha uma linguagem:")
            emoji.Println(":one: 🐹 Go")
            emoji.Println(":two: 🌳 Node.js")
            emoji.Println(":three: 🐍 Python")
            emoji.Print(":point_right: Sua escolha: ")
            languageChoice, _ := reader.ReadString('\n')
            languageChoice = strings.TrimSpace(languageChoice)

            switch languageChoice {
            case "1":
                language = "go"
            case "2":
                language = "node.js"
            case "3":
                language = "python"
            default:
                emoji.Println(":x: Opção inválida, usando Go como padrão.")
                language = "go"
            }
        }

        // Pergunta sobre banco de dados
        emoji.Print(":question: Deseja incluir um banco de dados? (sim/não): ")
        databaseAnswer, _ := reader.ReadString('\n')
        databaseAnswer = strings.TrimSpace(databaseAnswer)

        var database string
        if strings.ToLower(databaseAnswer) == "sim" {
            fmt.Println("Escolha um banco de dados:")
            emoji.Println(":one: 🐘 PostgreSQL")
            emoji.Println(":two: MySQL")
            emoji.Print(":point_right: Sua escolha: ")
            databaseChoice, _ := reader.ReadString('\n')
            databaseChoice = strings.TrimSpace(databaseChoice)

            switch databaseChoice {
            case "1":
                database = "Postgres"
            case "2":
                database = "MySQL"
            default:
            emoji.Println(":x: Opção inválida, sem banco de dados.")
            }
        }

        // Pergunta sobre cache
        emoji.Print(":question: Deseja incluir uma ferramenta de cache? (sim/não): ")
        cacheAnswer, _ := reader.ReadString('\n')
        cacheAnswer = strings.TrimSpace(cacheAnswer)

        var cache string
        if strings.ToLower(cacheAnswer) == "sim" {
            fmt.Println("Escolha uma ferramenta de cache:")
            emoji.Println(":one: 🧠 Redis")
            emoji.Println(":two: Memcached")
            emoji.Print(":point_right: Sua escolha: ")
            cacheChoice, _ := reader.ReadString('\n')
            cacheChoice = strings.TrimSpace(cacheChoice)

            switch cacheChoice {
            case "1":
                cache = "Redis"
            case "2":
                cache = "Memcached"
            default:
                emoji.Println(":x: Opção inválida, sem cache.")
            }
        }
        emoji.Printf("\n:gear: Configurações selecionadas:\nLinguagem: %s\nBanco de Dados: %s\nCache: %s\n", language, database, cache)

        // Gera o Dockerfile
        generateDockerfile(language)

        // Gera o docker-compose.yml
        generateDockerCompose(language, database, cache)

        // Constrói a imagem Docker
        buildDockerImage()
    },
}

// Execute inicia o rootCmd
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

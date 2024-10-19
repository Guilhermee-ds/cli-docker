package cmd

import (
    "fmt"
    "os"
    "strings"
    "github.com/kyokomi/emoji"
)

// Função para gerar o arquivo docker-compose.yml
func generateDockerCompose(language, database, cache string) {
    file, err := os.Create("docker-compose.yml")
    if err != nil {
        emoji.Println(":x: Erro ao criar o arquivo docker-compose.yml:", err)
        return
    }
    defer file.Close()

    // Gerar conteúdo do docker-compose.yml
    dockerComposeContent := `
version: '3.8'
services:
`

    // Adicionando o serviço da aplicação
    dockerComposeContent += fmt.Sprintf(`
  app:
    build: .
    ports:
      - "8080:8080"
`)
  // Adicionando o serviço do banco de dados, se selecionado
    if database != "" {
        if strings.ToLower(database) == "mysql" {
            dockerComposeContent += `
  mysql:
    image: mysql:8.0
    environment:
      - MYSQL_ROOT_PASSWORD=root
      - MYSQL_DATABASE=mydb
    volumes:
      - mysql_data:/var/lib/mysql
`
        } else if strings.ToLower(database) == "postgres" {
            dockerComposeContent += `
  postgres:
    image: postgres:14-alpine
    environment:
      - POSTGRES_USER=root
      - POSTGRES_PASSWORD=root
    volumes:
      - postgres_data:/var/lib/postgres/data
`
        }
    }

if cache != "" {
    if strings.ToLower(cache) == "redis" {
        dockerComposeContent += fmt.Sprintf(`
  redis:
    image: %s
    environment:
      - REDIS_PASSWORD=secret
    volumes:
      - redis_data:/data
`, getCacheImage(cache))
    } else if strings.ToLower(cache) == "memcached" {
        dockerComposeContent += fmt.Sprintf(`
  memcached:
    image: %s
    volumes:
      - memcached_data:/data
`, getCacheImage(cache))
    }
}

// Adicionando volumes
dockerComposeContent += `
volumes:
`

// Adicionar volume para o banco de dados
if database != "" {
    dockerComposeContent += fmt.Sprintf("  %s_data:\n", strings.ToLower(database))
}

// Adicionar volume para o cache, se for Redis ou Memcached
if cache != "" {
    if strings.ToLower(cache) == "redis" {
        dockerComposeContent += "  redis_data:\n"
    } else if strings.ToLower(cache) == "memcached" {
        dockerComposeContent += "  memcached_data:\n"
    }
}

    // Escreve o conteúdo no docker-compose.yml
    _, err = file.WriteString(dockerComposeContent)
    if err != nil {

        emoji.Println(":x: Erro ao escrever no docker-compose.yml:", err)
        return
    }
    emoji.Println(":page_facing_up: :white_check_mark: docker-compose.yml gerado com sucesso!")
}

// Função para obter a imagem do banco de dados baseado na escolha 
func getDatabaseImage(database string) string {
    switch strings.ToLower(database) {
    case "postgres":
        return "postgres:14-alpine"
    case "mysql":
        return "mysql:8.0"
    default:
        return ""
    }
}

// Função para obter a imagem de cache baseado na escolha
func getCacheImage(cache string) string {
    switch strings.ToLower(cache) {
    case "redis":
        return "redis:7-alpine"
    case "memcached":
        return "memcached:1.6-alpine"
    default:
        return ""
    }
}


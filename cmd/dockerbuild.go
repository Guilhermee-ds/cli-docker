package cmd

import (
    "fmt"
    "os/exec"
    "time"
    "os"
    "github.com/kyokomi/emoji"
)

func showLoader(done chan bool) {
    loaderChars := []string{"|", "/", "-", "\\"}
    i := 0
    for {
        select {
        case <-done:
            return
        default:
            emoji.Printf("\r%s :wrench: Criando...", loaderChars[i%len(loaderChars)])
            i++
            time.Sleep(100 * time.Millisecond)
        }
    }
}

// Função para construir a imagem Docker usando o Dockerfile
func buildDockerImage() {

    done := make(chan bool)
    go showLoader(done)

    cmd := exec.Command("docker", "build", "-t", "my-custom-image", ".")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()

    done <- true
    fmt.Println("\r")

    if err != nil {
        emoji.Println(":x: Erro ao construir a imagem Docker:", err)
    } else {
        emoji.Println(":whale: :white_check_mark: Imagem Docker construída com sucesso!")
    }
}


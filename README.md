
# CLI - Criação de DockerFile e Docker-Compose

A CLI gera automaticamente um Dockerfile e um docker-compose com as configurações desejadas, criando uma imagem enxuta usando Alpine e Slim. O usuário pode testar facilmente os bancos de dados e o Redis! 🔧🐳

Desenvolvi uma CLI que permite aos usuários selecionar:

- Linguagens: Go, Node.js e Python 🖥️
- Banco de Dados: PostgreSQL e MySQL 🗄️
- Ferramenta de Cache: Redis e Memcache ⚡

## Tecnologias Utilizadas:

- Golang
- Docker

## Dependências:

- Golang >= go1.22.1
- Docker
- Docker Compose
- Redis CLI (Opcional)

## Como Rodar

Depois de ter todas a dependências informadas anteriormente

Instale as bibliotecas do projeto

```shell
go mod tidy

```

Rode a aplicação (verifique se o Docker está iniciado na sua máquina).

```shell
go run main.go

```
Prontinho, agora selecione a linguagem, banco de dados e a ferramenta de cache. 😊

verifique se a imagem foi gerada.

```shell
sudo Docker images

```

No diretório onde o Dockerfile e o docker-compose foram gerados (raiz do projeto), rode esse comando

```shell
docker-compose up --build -d

```
Agora, verifique se os containers estão em execução

```shell
sudo Docker ps

```

Caso o banco de dados selecionado e a ferramenta de cache estejam em execução, teste-os.

Para testar as ferramentas.

PostgreSQL
- Host = localhost
- password = root
- user = root
- DB_name = postgres


Redis
- Host = localhost
- password = secret
- user = root

O que pode mudar é o host. Para saber o host dos containers, basta rodar este comando
```shell
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' nome-do-container

```


## Finalização
Parabéns! A documentação do projeto foi concluída com sucesso. As ferramentas desenvolvidas estão prontas para uso e foram testadas para garantir sua funcionalidade e eficiência.

Esperamos que este projeto seja útil e facilite o trabalho de muitos desenvolvedores. Continuarei a aprimorar e atualizar as ferramentas conforme necessário.

Boas práticas e feliz codificação 😊!

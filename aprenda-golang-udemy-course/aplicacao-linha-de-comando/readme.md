# Aplicação de linha de comando GOLANG

## Descricao

- (1) A aplicação irá receber um domínio/endereço web em ipv4 público do site
- (2) A aplicação era receber um domínio / endereço web e retornar o servidor em qual o site está hospedado

## Passos

- go mod init linha-de-comando ( inicializar o gerenciador de pacotes/dependencias)
- go get github.com/urfave/cli
- go build

## Comandos

### Via go cli

- (1) go run main.go ip --host www.google.com.br
- (2) go run main.go servidores --host www.google.com.br

### Via executavel

- (1) ./linha-de-comando ip --host www.google.com.br
- (2) ./linha-de-comando servidores --host www.google.com.br

## Anotações

- go.sum == package.json-lock
- go.mod == package.json

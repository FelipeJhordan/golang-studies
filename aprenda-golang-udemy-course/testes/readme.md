# Testes com GO

## Comando para rodar os testes

- go test ( rodar os testes do pacote que esta localizado )
- go test ./... ( rodar todos os testes em todos os sub-diretórios)
- go test -v ./... (rodar testes em modo verboso)
- go test ./... --cover (rodar testes e gerar um log de cobertura)
- go test --coverprofile cobertura.txt ( criar um arquivo txt com log de cobertura)
- go tool cover --func=cobertura.txt ( mostrar um log no terminal sobre quais linhas estão cobertas)
- go tool cover --html=cobertura.txt ( mostrar um html contendo um log sobre o que está sendo coberto ou não)

## Observações

- O arquivo de teste precisa ter o suffixo "\_test"
- A função de teste propriamente dita, precisa estar nesse formato TestXxxxx

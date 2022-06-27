# Canais / Channels

## Descrição

Maneira mais utilizada para sincronizar as goroutines, ou para fazer elas rodarem ao mesmo tempo.

## Como utilizar os canais

### Como criar

```
canal:= make(chan string)
```

### Como utilizar 1

```
canal:= make(chan string)
go escrever("Olá mundo", canal)
mensagem := <- cana
fmt.Println(mensagem)

func escrever(texto string, canal chan string) {
    for i:=0 i < 5; i++ {
          canal <- texto
          time.Sleep(time.Second)
    }
}
```

### Como utilizar 2 ( com close )

```
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for i := 0; i < 5; i++ {
		mensagem, isOpen := <-canal
		if !isOpen {
			break
		}
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
```

### Como utilizar 3 ( com for of)

```
func main() {
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")



	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
```

## Minhas Anotações

- operacoes bloqueantes
- parece seguir o padrão observer

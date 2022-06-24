# Loops

## Não existe while, do while, foreach em go, todo é um for ou utiliza da palavra for...

## While

### Não existe um comando interno chamado "while", no seu lugar é utilizado uma sintaxe do for:

```
    for i < 10 {
        i++
        fmt.Println("incrementando i")
        time.Sleep(time.Second)
    }
```

## For padrão em outras linguagens

```
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}
```

## Utilização de Range ( tipo for of ou for in)

```
    nomes := []string{"Felipe", "Jhordan", "Maciel"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
    for _, nome := range nomes {
	fmt.Println(nome)
	}
    for indice, letra := range "PALAVRA" {
	fmt.Println(indice, string(letra))
	}
```

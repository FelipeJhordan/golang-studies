# Structs

## Structs é o mais perto que temos de uma classe em GO, podemos criar um struct a partir do

```
type Usuario struct {
	nome     string
	idade    uint8
	endereco Endereco
}
```

## Podemos declarar um struct

var usuario Usuario

## Podemos fazer a alocação apartir de uma inferência

```
usuario:= Usuario{
    "Felipe",
    21,
    Endereco{
        logradouro: "Rua monica machiyama",
        numero: 213,
    },
}
```

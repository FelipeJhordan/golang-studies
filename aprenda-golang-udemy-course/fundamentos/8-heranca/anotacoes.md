## "Heranca" em go

para conseguir fazer um struct herdar atributos de outro struct é necessário apenas declarar o tipo do struct pai

```
type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}
```

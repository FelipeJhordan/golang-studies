# Metodos em structs

## Você precisa referenciar a qual struct o metodo vai ser vinculado

```
type Usuario struct {
	nome  string
	idade uint8
}


func (u Usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados \n", u.nome)
}
func (u Usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

```

## Para alterar algum valor da "instancia", você precisa criar um método com sinatura por ponteiro

```
func (u *Usuario) fazerAniversario() {
	u.idade++
}
```

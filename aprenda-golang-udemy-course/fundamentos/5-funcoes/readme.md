# Funções em GO

## Particularidades

### Caso os parametros de uma função sejam todos de um mesmo tipo, da para você só especificar no "final"

`func somar(n1, n2 int8): int8 {}`

### Funções em go podem ter 2 ou mais retornos ( isso é muito estranho)

```
func calculosMatematicos(n1, n2 int8) (int8, int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	subtracao1 := n1 - n2 + 1
	return soma, subtracao, subtracao1
}
```

### Você pode ou não recebir todos os retornos...

```
resultadoSoma, _, subtracao1 := calculosMatematicos(10, 15)
```

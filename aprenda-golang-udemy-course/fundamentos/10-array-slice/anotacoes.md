# Array e Slice

## Array

### Array no go aceita apenas um tipo/struct, e também precisa definir na sua declaração.

```
	var array1 [5]string
	array1[0] = "Posição 0"
	array := [5]string{"Posição1", "Posiçã
```

## Slice

### Slice é mais parecido com uma lista ou array no javascript por exemplo, você não precisa especificar o tamanho de um slice, e também tem um método append ( push ) onde você adiciona novos itens.

```
	slice := []int{10, 11, 12, 15}
	slice = append(slice, 18)
```

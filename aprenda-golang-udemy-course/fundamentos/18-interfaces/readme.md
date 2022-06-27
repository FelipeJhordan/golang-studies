# Interfaces em go

## O uso de interfaces em go é implicita ( não precisamos declarar durante a criação da classe/struct)

```
type forma interface {
	area() float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}



```

## Com interfaces dá para você tirar aproveito do polimorfismo

```
func escreverArea(f forma) {
	fmt.Println("A área da forma é %0.2f", f.area())
}
```

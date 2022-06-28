package formas

import "testing"

func TestArea(t *testing.T) {
	t.Run("Retangulo", func (t *testing.T) {
		ret := Retangulo{
			10, 12
		}
		areaEsperada := 120
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.FatalF("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func (t *testing.T) {
		circulo := Circulo{10}

		areaEsperada := float64(math.PI * 100)
		areaRecebida := circulo.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Quadrado", func (t *testing.T) {
		Qua := Quadrado{3, 4}
		areaEsperada := float64(12)
		areaRecebida := Qua.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area recebida %f é diferente da area esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area recebida %f é diferente da area esperada %f", areaRecebida, areaEsperada)
		}
	})
}

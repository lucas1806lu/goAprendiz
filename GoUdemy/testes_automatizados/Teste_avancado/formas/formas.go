package formas

import "math"



type Forma interface {
	Area() float64
	//perim() float64
}

type Quadrado struct {
	Largura, Altura float64
}
type Circulo struct {
	Raio float64
}

func (q Quadrado) Area() float64 {
	return q.Largura * q.Altura
}
//func (q quadrado) perim() float64 {
//	return 2*q.largura + 2*q.altura
//}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}
//func (c círculo) perim() float64 {
//	return 2 * math.Pi * c.raio
//}

//func Medir(g Forma) {
//	fmt.Println(g)
//	fmt.Println(g.Area())
	//fmt.Println(g.perim())
//}
//func main() {
//	q := Quadrado{Largura: 3, Altura: 4}
//	c := Circulo{Raio: 5}

//	Medir(q)
//	Medir(c)
//}

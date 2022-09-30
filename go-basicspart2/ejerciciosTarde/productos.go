package main

type Tienda struct {
	Products []producto
}

type producto struct {
	Type  string
	Name  string
	Price float64
}

type ProductController interface {
	CalcularCosto() float64
}

func (p producto) CalcularCosto() float64 {

	switch p.Type {
	case "Mediano":
		return p.Price * 0.03
	case "Grande":
		return p.Price*0.06 + 2500
	default:
		return 0
	}
}

type Ecommerce interface {
	Total() float64
	Agregar(producto)
}

func (t Tienda) Total() float64 {
	total := 0.0
	for _, v := range t.Products {
		total += v.Price + v.CalcularCosto()
	}
	return total
}

func nuevoProducto(tipo, nombre string, precio float64) producto {
	return producto{tipo, nombre, precio}
}

func nuevaTienda() Ecommerce {
	return &Tienda{}

}

func (t *Tienda) Agregar(p producto) {
	t.Products = append(t.Products, p)
}

func main() {

}

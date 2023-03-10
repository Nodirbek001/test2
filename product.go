package main

type Product struct {
	Name     string
	Quantity int
	Price    int
}

type ProductList []Product

func NewProduct(name string, quantity int, price int) Product {
	return Product{
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}
}
func (l *ProductList) remove(index int) {
	*l = append((*l)[:index], (*l)[:index+1]...)

}

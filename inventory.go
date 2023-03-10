package main

type Inventory struct {
	Products ProductList
}

func NewInventory(products ProductList) Inventory {
	return Inventory{
		Products: products,
	}
}

func (i *Inventory) Searche(prodecteName string) (Product,bool){
	for _, product := range i.Products {
		if product.Name==prodecteName {
			return product, true
		}
	}
	
	return Product{}, false
}
func(i *Inventory) TakeProduct(name string, quantity int){
	for inx:= range i.Products{
		if i.Products[inx].Name== name{
			i.Products[inx].Quantity-=quantity
			if i.Products[inx].Quantity==0 {
				i.Products.remove(inx)
			}
		}
	}
}
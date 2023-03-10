package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Store struct{
	Inventory Inventory
	Cash int
}

type SellProductRequest struct{
	ProductName string
	Quantity int
}


func NewStore(sum int ,inventory Inventory) Store{
	return Store{
		Inventory: inventory,
		Cash:  sum,
	}
}


func (s *Store) addToCash(sum int){
	s.Cash+=sum
}
func (s *Store)Sell(r SellProductRequest){
	product, exists:=s.Inventory.Searche(r.ProductName)
	if !exists {
		fmt.Printf("product with this name not fount")

		return
	}

	if product.Quantity<r.Quantity {
		fmt.Printf("not enough products: only %d left\n",product.Quantity)
	}
	s.Inventory.TakeProduct(r.ProductName, r.Quantity)
	s.addToCash(r.Quantity*product.Price)

}

func( s *Store) printStats(){
	w:=tabwriter.NewWriter(os.Stdout, 10,8,1,' ', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w,"inventory")
	fmt.Fprintln(w,"name\tquantity,\tprice")
	for _,prduct:=range s.Inventory.Products{
		fmt.Fprintf(w,"%s\t%d\t%d\n", prduct.Name,prduct.Quantity, prduct.Price)
	}
	w.Flush()
	fmt.Println("Cash", s.Cash)
}


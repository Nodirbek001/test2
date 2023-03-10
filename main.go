package main
var prducts =ProductList{
	
	{
		Name: "Apple",
		Quantity: 10,
		Price: 10,
	},
	{
		Name: "Banan",
		Quantity: 5,
		Price: 5,
	},
	{
		Name: "Book",
		Quantity: 6,
		Price: 6,
	},
	
}
func main(){
 i:=NewInventory(prducts)
 s:=NewStore(1000,i)
 s.printStats()
 s.Sell(SellProductRequest{
	ProductName: "Apple",
	Quantity: 2,
 })

}
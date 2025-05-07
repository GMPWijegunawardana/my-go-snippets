package main
import "fmt"


func main(){
	var nameofproduct1 string = "Laptop"
	var price1 float64 = 1050.99
	var quantity1 int = 3
	var inStock1 bool = true

	totalprice1 := price1 * float64(quantity1)

	fmt.Println("Product Name:\n", nameofproduct1)
	fmt.Println("Price:\n", price1)
	fmt.Println("Quantity:\n", quantity1)
	fmt.Println("In Stock:\n", inStock1)

	if inStock1 {
		fmt.Println("Total Cost:\n", totalprice1)
	} else {
		fmt.Println("Product is out of stock.")
	}

	var nameofproduct2 string = "dekstop"
	var price2 float64 = 12050.99
	var quantity2 int = 5
	var inStock2 bool = true
	
	totalprice2 := price2 * float64(quantity2)

	fmt.Println("Product Name:\n", nameofproduct2)
	fmt.Println("Price:\n", price2)	
	fmt.Println("Quantity:\n", quantity2)
	fmt.Println("In Stock:\n", inStock2)

	if inStock2 {
		fmt.Println("Total Cost:\n", totalprice2)
	} else {
		fmt.Println("Product is out of stock.")

}
}
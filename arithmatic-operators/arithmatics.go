package main 

import "fmt"

func main() {

	a := 10
	b := 23

	sum := a + b
	sub := a - b
	Product := a * b
	division := a / b
	remainder := a % b

	fmt.Println("Sum:", sum)
	fmt.Println("Subtraction:", sub)	
	fmt.Println("Product:", Product)
	fmt.Println("Division:", division)
	fmt.Println("Remainder:", remainder)
	fmt.Println("Increment:", a+1)
	fmt.Println("Decrement:", b-5)

	// using += , -= symbols

	//compound assignment operators
	c :=15

	c += 5
	fmt.Println("Increment:", c)

	c-= 5
	fmt.Println("Decrement:", c)

	/* increment oparator
	c++
	
	c-- (decrement operator)
	
	*/

	//convertions 

	var intNum int = 10
	var floatNum float64 = 20.5

	Sum := intNum + int(floatNum) // mekedi warahan athule dammama float ekath int ekak widiyata gannawa 
	fmt.Println(Sum)

	// eth mekedi enne 30 witharai dashama 5 print wen  na
	

	Sum2 := float64(intNum) + floatNum 
	fmt.Println(Sum2)  // mehema dammama 30.5 widiyata eno
	
		





}

package main

import "fmt"

func main()  {
	fmt.Println("estes são os numeros divisiveis por 3")
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			
			fmt.Println(i)
		}
		
	}
}
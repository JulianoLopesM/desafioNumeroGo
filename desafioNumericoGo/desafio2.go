package main

import "fmt"

func main()  {
	for i := 0; i <= 100; i++ {
		switch  {
		case i%3 == 0 && i%5==0:
			fmt.Println( i ,"pinpam")
		case i%3==0:
			fmt.Println(i ,"pin")
		case i%5 == 0:
			fmt.Println(i , "pam" )
			
			
		}
	}
	
}
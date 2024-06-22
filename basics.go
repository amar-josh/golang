package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main(){
	fmt.Printf("My favourit number is: %d\n", rand.Intn(10));
	fmt.Printf("Square toot: %f \n", math.Sqrt(16))
	fmt.Printf("%v \n",math.Pi)
	fmt.Println(math.Pi)
}
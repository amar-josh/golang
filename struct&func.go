// package main

// import "fmt"

// type Data struct {
// 	a int
// 	b int
// }

// func NewData(a, b int) Data {
// 	return Data{a, b}
// }

// func (r Data) Area() int {
// 	area := r.a * r.b
// 	return area
// }

// func (r Data) Perimeter() int {
// 	return 2 * (r.a * r.b)

// }

// func main() {
// 	r := NewData(2, 3)
// 	fmt.Printf("area of circle: %v\n", r.Area())
// 	fmt.Printf("area of circle: %v", r.Perimeter())
// }
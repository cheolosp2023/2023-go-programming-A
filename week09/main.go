package main

import "fmt"

func swap(n1 *int, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp //
}

func main() {
	a := 10
	b := 20 //var b int = 20

	c := &a //var c *int = &a
	fmt.Printf("%T\n", c)
	fmt.Println(a, b)
	swap(&a, &b)      //pass by pointer
	fmt.Println(a, b) //포인터를 사용해서 a와 b의 순서 바꾸기
}

/*
포인터 사용
package main

import "fmt"

func double(n *int) {
	*n = *n * 2
}

func main() {
	var a int = 6
	double(&a)
	fmt.Println(a)
}
*/

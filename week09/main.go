package main

import "fmt"

func main() {
	a := 10
	//var b int = 20
	var pa *int = &a

	fmt.Printf("%T %T\n", &a, pa)
	fmt.Printf("%x %x %x\n", &a, pa, &pa) //주소값 출력(pa와 &pa 다름)
	fmt.Println(&a, pa, &pa)              //위와 같음(포인터는 메모리 주소값 출력)
	fmt.Println(*pa)
	pa = &b //b의 메모리 주소를 받음
	fmt.Println(*pa)
}

/*
package main

import "fmt"

func swap(n1 *int, n2 *int) {
	//fmt.Println(n1, n2) //16진수의 주소값(포인터는 주소값 반환)
	temp := *n1
	*n1 = *n2
	*n2 = temp
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
*/
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

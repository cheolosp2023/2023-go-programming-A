package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	//var a int 			//declaration	
	//a =9					//assign value
	
	//var a int = 9			//declaration & assign value
	//var a = 9				//declaration & assign value

	a := 9 					//short form.declaration & assign value
	fmt.Println(a, reflect.TypeOf(a))

	fmt.Println(reflect.TypeOf('Z'))		//rune, int32
	fmt.Println(reflect.TypeOf('99'))
	fmt.Println(reflect.TypeOf('2,7'))
	fmt.Println(reflect.TypeOf('false'))
	fmt.Println(reflect.TypeOf("Go!"))
	fmt.Println('A', 'a', '0', '김', '인', '하')  //65, 97, 48
	fmt.Println(math.Ceil((3.17))
	fmt.Println(math.Floor((3.17))
	//fmt.Println(strings.Title("오픈소스 프로그래밍~"))
	fmt.Println(strings.Title("opensource programming~ \"go\"!")
}

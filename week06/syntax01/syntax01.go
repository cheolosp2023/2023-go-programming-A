package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.Floor(2.71)) //내림
	fmt.Println(math.Ceil(2.71))  //올림
	fmt.Println(math.Round(2.71)) //반올림
	fmt.Println("Hello Go~")
	fmt.Println(strings.Title("go git github java")) //단어의 첫 글자는 대문자, 나머지는 소문자

	//변수선언, 변수이름, 변수타입
	var c rune = '가' //"d"=문자열 'd'=rune=한 글자 유니코드(UTF-8) 값(rune은 int32의 별명 4byte)
	fmt.Println(c)
	fmt.Printf("%c\n", c) //문자 그대로 출력
	fmt.Printf("%T\n", c) //타입 출력

	//var a int16 = 7
	//var a = 7
	a := 7 //위와 동일한 선언
	fmt.Println(a)
	fmt.Printf("%d\n", a)
	fmt.Printf("%T\n", a)

	b := 7.3
	b = 5 //다른 값이 들어갈 순 있지만 타입이 변하진 않음(float->int는 가능 5.0으로 바뀜 하지만 int->float은 오류)
	fmt.Printf("%d\n", b)
	fmt.Printf("%T\n", b)
	println("----------------")
	e := 7
	fmt.Println(e)
	var d float64 = 5.3
	e = int(d) //Type Conversion, Casting(형 변환)
	fmt.Printf("%d\n", e)
	fmt.Printf("%T\n", e)

	f := false
	fmt.Printf("%c\n", f)
	fmt.Printf("%T\n", f)
}

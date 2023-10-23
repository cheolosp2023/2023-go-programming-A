package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	texts := "G@ M@ney~"
	fmt.Println(texts)

	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)

	//변수명은 영문자로 시작
	//대문자로 시작 경우 다른 패키지에서 접근 가능, 소문자로 시작은 동일 패키지에서만 접근 가능
	a := 7
	var c rune = '가'
	var b float64
	var d int
	var e rune
	var f bool   //bool의 zero value는 false
	var g string //string의 zero value는 null
	var studentId string
	var i int32

	//zero value(변수에 값을 할당하지 않았을 때 기본적으로 갖는 값)
	fmt.Println(b)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	//naming conventuion
	fmt.Println(studentId)
	fmt.Println(i)

	println("-------------")
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)

	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(g))

	println("-------------")
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", c)
	println("-------------")

}

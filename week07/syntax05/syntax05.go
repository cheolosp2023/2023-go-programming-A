package main

import "fmt"

// shadowing
func main() {
	//자료타입을 변수명으로 사용
	//var float64 float64 = 9.1
	//var test float64 = 7.9 //float64(자료타입)로 변수를 설정하면 shadow가 일어나서 오류 발생
	//fmt.Println(float64)
	var test1 float64 = 9.1
	var test2 float64 = 7.9
	fmt.Println(test1 + test2)

	//패키지를 변수명으로 사용
	//var fmt string = "inha" //fmt로 변수 설정한 경우에도 오류 발생
	//fmt.Println(fmt)
	var univ string = "inha"
	fmt.Println(univ)

	//함수를 변수명으로 사용
	//var append string = "functions" //append로 변수 설정한 경우
	//var f = append([]string{}, "함수")
	var f1 string = "functions"
	var f2 = append([]string{}, "함수")
	fmt.Println(f1)
	fmt.Println(f2)

}

/*
package main

import (
	"fmt"
	"log"
	"os"
)

// 입력(0과 1처리)된 수의 소수 판정 프로그램 ver0.9
func main() {
	var number int

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number)

	if err != nil {
		log.Fatal(err)
	}

	for number < 2 {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
		os.Exit(0)
	}

	isPrime := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/
/*
package main

import (

	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

)

// 입력된 수의 소수 판정 프로그램 ver0.7

	func main() {
		fmt.Print("정수 입력 : ")
		rd := bufio.NewReader(os.Stdin)
		in, err := rd.ReadString('\n')

		if err != nil { //에러가 발생하면
			log.Fatal(err)
		}

		in = strings.TrimSpace(in)
		number, err := strconv.Atoi(in)

		if err != nil {
			log.Fatal(err)
		}

		isPrime := true

		for i := 2; i < number; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
			fmt.Print(i, " ")
		}

		if isPrime {
			fmt.Println(number, "는(은) 소수입니다.")
		} else {
			fmt.Println(number, "는(은) 소수가 아닙니다.")
		}
	}
*/

/*
package main

import (
	"fmt"
	"log"
)

// 입력(fmt 패키지의 Scanln())된 수의 소수 판정 프로그램 ver0.8
func main() {
	var number int
	fmt.Print("정수 입력 : ")
	n, err := fmt.Scanln(&number) //n=입력받은 개수

	fmt.Println(n)

	if err != nil {
		log.Fatal(err)
	}

	isPrime := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//seed 생성용 패키지

/*
	func main() {
		//seed 설정(시간 설정->시간이 흐를 때마다 다른 랜덤 수(다시 주사위를 돌리면 다른 수가 나오게 설정))
		seed := time.Now().Unix()
		rand.Seed(seed)

		for i := 1; i < 6; i++ {
			dice := rand.Intn(6) + 1 //rand.Intn(n)0 ~ n-1 사이의 랜덤 수
			fmt.Println(dice)
		}
	}
*/

/*
// 난수 추출된 수의 소수 판정 프로그램 ver0.1
// 소수 : 약수가 0과 자기 자신
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0

	number := rand.Intn(150) + 2 //2~ 151 사이의 수
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 1; i < number; i++ {
		if number%i == 0 {
			count++
		}
	}
	if count == 2 {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/

/*
// 난수 추출된 수의 소수 판정 프로그램 ver0.2
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0
	number := rand.Intn(150) + 2 //2~ 151 사이의 수
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ { //1과 number일 때 loop 돌지 않음
		if number%i == 0 {
			//count++
			count = count + 1
		}
	}
	if count == 0 {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/

/*
// 난수 추출된 수의 소수 판정 프로그램 ver0.3
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true //true이면 소수
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
		}
	}

	if isPrime == true {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/

// 난수 추출된 수의 소수 판정 프로그램 ver0.4
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true //true이면 소수
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
		fmt.Print(i, " ") //앞에서 (2,3 ...) 나누어 떨어지면 break -> 속도 향상
	}

	if isPrime { //비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}

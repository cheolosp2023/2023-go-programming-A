package main

import (
	"bufio"
	"fmt"
	"log" //Fatal 함수 사용(에러 메세지 출력, 프로그램 종료)
	"os"
	"strconv" //TrimSpace 함수 사용
	"strings" //ParseInt
)

func main() {
	fmt.Print("단 입력 : ")
	rd := bufio.NewReader(os.Stdin) //키보드로 입력 읽기
	//in, _ := rd.ReadString('\n')  //, _ : 에러가 발생한 상황에서 받을 수 있는 부분(_은 무시하고 진행)
	in, err := rd.ReadString('\n') //nil이면 오류x nil이 아니면 오류

	if err != nil { //에러가 발생하면
		log.Fatal(err)
	}

	in = strings.TrimSpace(in)
	//dan, err := strconv.ParseInt(in, 8, 32) //두 번째 칸 : 몇 진수 인지 설정
	dan, err := strconv.Atoi(in) //알아서 변형

	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < 10; i++ {
		//fmt.Println(dan, " * ", i, " = ", (int(dan) * i)) //ParseInt인 경우 형변환(디테일한 설정 가능)
		//fmt.Println(dan, " * ", i, " = ", (dan * i))      //Atoi인 경우 형변환 x
		fmt.Printf("%d * %d = %d\n", dan, i, (dan * i))
	}

	//go에는 while문 없음. 다른 언어의 while문 구현
	i := 1
	for i < 10 {
		fmt.Printf("%d * %d = %d\n", dan, i, (dan * i))
		i++
	}
}

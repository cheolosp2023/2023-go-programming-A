package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	now = time.Now()
	//var year int = now.Year()  아래와 같음
	year := now.Year()
	var month string = now.Month().String()
	fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())
}

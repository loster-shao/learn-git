package main

import (
	"fmt"
	"time"
)



func main() {
	var time1,time2,time3,time4,time5 int64
	var sc string
	fmt.Scanf("%d",&time1 )
	fmt.Scanf("%d",&time2 )
	fmt.Scanf("%d",&time3 )
	fmt.Scanf("%d",&time4 )
	fmt.Scanf("%d",&time5 )
	fmt.Scanf("%s",&sc )
	if time1!=0 {
		fmt.Println("input ok!")
	}else {
		fmt.Println("no input")
	}
	if time2!=0 {
		fmt.Println("input ok!")
	}else {
		fmt.Println("no input")
	}
	if time3!=0 {
		fmt.Println("input ok!")
	}else {
		fmt.Println("no input")
	}
	if time4!=0 {
		fmt.Println("input ok!")
	}else {
		fmt.Println("no input")
	}
	if time5!=0 {
		fmt.Println("input ok!")
	}else {
		fmt.Println("no input")
	}
		Tm1 := time.Unix(time1, 1)
		Tm2 := time.Unix(time2, 1)
		Tm3 := time.Unix(time3, 1)
		Tm4 := time.Unix(time4, 1)
		Tm5 := time.Unix(time5, 1)
	fmt.Println(sc)
		if sc == "result"{
			fmt.Println("ok")
			fmt.Println(Tm1, "\n", Tm2, "\n", Tm3, "\n", Tm4, "\n", Tm5)
		}else {
			fmt.Println("no")
		}
	fmt. Println(sc=="result")
	//if strings.EqualFold(SC, "go") {
	//	fmt.Println(" ok!")
	//	fmt.Println(Tm1, "\n", Tm2, "\n", Tm3, "\n", Tm4, "\n", Tm5)
	//} else {
	//	fmt.Println("no ")
	//}

}

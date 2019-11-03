package main

import ("fmt")

var k int

func main() {
	fmt.Println("请输入你要求素数的个数")
	k=string("n")
	get_prime(k)
}

func get_prime(n int) {
	origin, wait := make(chan int), make(chan struct{})
	go FilterPrime(origin, wait)
	for i := 2; i <= n; i++ {
		origin <- i
	}
	close(origin)
	<-wait
}

func FilterPrime(seq chan int, wait chan struct{}) {
	prime, ok := <-seq
	if !ok {
		close(wait)
		return
	}
	fmt.Println(prime)        //打印素数
	out := make(chan int)
	go FilterPrime(out, wait)     //启动goroutine
	for num := range seq {
		if num % prime != 0 {
			out <- num
		}
	}
	close(out)
}
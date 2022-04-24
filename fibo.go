package main

import (
	"flag"
	"fmt"
	"os"
)

func fibo(x, y, z int) {
	fibonacci := []int{y, z}
	fmt.Printf("1: %d\n2: %d\n", fibonacci[0], fibonacci[1])
	for i := 2; i < x; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
		fmt.Printf("%d: %d\n", i+1, fibonacci[i])
	}
}

func main() {
	var help bool
	flag.BoolVar(&help, "h", false, "Help with using program")
	flag.Parse()
	first := flag.Int("first", 0, "First number of Fibonacci seqence")
	secound := flag.Int("secound", 1, "Secound number of Fibonacci seqence")
	num := flag.Int("num", 10, "How many numbers to calculate")
	if help == true {
		fmt.Println("You have to provide 3 arguments - first number, secound number, how many numbers to calculate\nExample:\t./fibo first=0 secound=1 num=10\n")
		os.Exit(1)
	}
	fibo(*num, *first, *secound)
	os.Exit(0)
}

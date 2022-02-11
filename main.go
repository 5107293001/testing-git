package main
import ("fmt")
const PI = 3.14
var i,j,k = 1,2.2,"abc"
func main () {
	//x := 10
	// const PI = 3.14 //constant global variable
	// fmt.Println ("hello go")// global variable
	// var  age uint8 = 100
	// var B float32 = 1.223
	// fmt.Println("value of a =", age)
	// fmt.Println("Value of b=", B)
	//i = 2
	// var i,j,k int
	// i=1
	// j=2
	// k=3
	// fmt.Println(i,j,k,PI,x)

	// y := "10.10"
	// fmt.Printf("Type of y = %T\n",y)

//example ()
}

func example2()  {
	/*
	000 = 0
	001 = 1
	010 = 2
	011 = 3
	100 = 4
	101 = 5
	110 = 6
	111 = 7
	1000 = 8 
	*/
	var a = 2
	a = a << 0

	fmt.Println(a)
	var b,c = 2,1
	d := b & c
	fmt.Println(d)
	}
	func example () {
		age := 80
		fmt.Println("IN example =",age ,PI)
	}



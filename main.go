package main
import ("fmt")
const PI = 3.14
var i,j,k = 1,2.2,"abc"
func main () {
	const PI = 3.14 //constant global variable
	fmt.Println ("hello go")// global variable
	var age uint8 = 100
	var b float32 = 1.223
	fmt.Println("value of a =", age)
	fmt.Println(b)
	var i,j,k int
	i=1
	j=2
	k=3
	fmt.Println(i,j,k,PI)

	y := "abc"
	fmt.Printf("Type of y = %T\n",y)

}
func exmaple() {
	var age uint8 = 80 //local variable
	fmt.Println("In example=",age,)
}
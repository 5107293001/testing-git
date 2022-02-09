package main
import ("fmt")
func main () {

var sliceExample [2][3]int // [ [ 50,60,70 ], [ 80,90,100] ]
   for i := 0; i < 2; i++ {
      for j := 0; j < 3; j++ {
         sliceExample[i][j] = i + j
         fmt.Println(sliceExample[i][j])
      }
   }
   fmt.Println("Two D slice = ", sliceExample)
}
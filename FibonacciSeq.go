package main
import (
   "fmt"
 )
// Fibonacci function
func Fibonacci(x int) int {
  if x <= 3 {
    return 1
  } else {
    return Fibonacci(x-1) + Fibonacci(x-2)
  }
}
// the main function 
func main() {
 var fib_var int
 fmt.Printf("Give the number \n")
 fmt.Scanf("%d", &fib_var)
 x := Fibonacci(fib_var)
 fmt.Printf("The %d Fibonacci number is: %d\n",fib_var,x)
}

package main
import "fmt"

func main() {
	val:= add(1,2)
	fmt.Println("sum of 1 & 2 is: ",val)

	div, mod := divmod(10,2)
	fmt.Printf("quotient : %d & remainder: %d\n", div, mod)
}

//sum of 2 numbers
func add(a int, b int) int {
	return a + b
}

//divmod returns quotient and remainder [can return more than 1 value]
func divmod(a int, b int)(int, int) {
	return a/b, a%b
}

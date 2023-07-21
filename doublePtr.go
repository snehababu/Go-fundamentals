package main
import "fmt"

func main() {
	val := 10
	doublePtr(&val)
	fmt.Println(val)
}

func doublePtr(n *int) {
	*n *= 2 
}

/*
o/p : 20
which means that the value has been changed into 20 with use of pointers
*/
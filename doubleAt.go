package main
import "fmt"

func main(){
	values := []int{1,2,3,4} //declared a slice with no fixed size
	doubleAt(values, 2)
	fmt.Println("slice: ",values)

	val := 10
	double(val)
	fmt.Println("double: ",val)
}

func doubleAt(values []int, i int) {
	values[i] *=2
}

/*
the return slice will have the element in 2nd position doubled
Eg: [1,2,6,4] (3 * 2 = 6)

The slice will behave like a pointer meaning we will create the copy of the slice but it points to the same location in the memory so that the effect is reflected outside of the function
*/

func double(n int) {
	n *= 2
}
/*
o/p : 10
Go is passing parameter by value to the function meaning it is creating a copy of the integer and passing it to the double function
*/

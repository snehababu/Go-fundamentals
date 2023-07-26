package main
import "fmt"

func main(){
	worker()
}

func worker() {
	r1,err := acquire("A")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	defer release(r1) 
	/*defer is a defered function called meaning the release will be called only when the worker function exits */
	fmt.Println("worker")

	//what if we have one more defer
	r2, err := acquire("B")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	defer release(r2) 
}

func acquire(name string) (string, error) {
	return name, nil //returns a name and possible error
}

func release(name string) {
	fmt.Printf("Cleaning up %s\n", name) //release the resource
}

/*
o/p:
worker
cleaning up A

once the worker function exust then only the defer function is called
*/
package print

import "fmt"

//PrintLn function for print values
func PrintLn(values ...interface{}) {
	for i, v := range values {
		fmt.Printf("Value %v Index %d", v, i)
	}
}

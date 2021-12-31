package print

import "fmt"

//PrintLn function for print values
func PrintLn(values ...interface{}) {
	for _, v := range values {
		fmt.Println(v)
	}
}

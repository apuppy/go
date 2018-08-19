// calc.go
package main

import "os" // 用于获得命令行参数os.Args
import "fmt"
import "simplemath"
import "strconv"

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe command are :\n add \tAddition of two values.\n\tsqrt\tSquare root of a non-gegative value.")
}

func main() {
	args := os.Args
	if len(args) < 2 {
		Usage()
		return
	}
	
	switch args[1] {
		case "add":
			if len(args) != 4 {
				fmt.Println("USAGE: calc add <integer1> <integer2>")
				return
			}
			
			v1, err1 := strconv.Atoi(args[2])
			v2, err2 := strconv.Atoi(args[3])
			if err1 != nil || err2 != nil {
				fmt.Println("USAGE: calc add <integer1> <integer2>")
				return
			}
			ret := simplemath.Add(v1,v2)
			fmt.Println("Result: ", ret)
		case "sqrt":
			if len(args) != 3 {
				fmt.Println("USAGE: calc sqrt <integer>")
				return
			}
			v, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("USAGE: calc sqrt <integer>")
				return
			}
			ret := simplemath.Sqrt(v)
			fmt.Println("Result: ",ret)
		default:
			Usage()
	}
}

package main

import "fmt"
import "github.com/mattn/go-colorable"

func main() {
	out := colorable.NewColorableStdout()
	fmt.Fprintf(out, "ʕ◔ϖ◔ʔ"+"\x1b[31m"+"💓"+"\x1b[0m"+"ʕ◔ϖ◔ʔ")
}

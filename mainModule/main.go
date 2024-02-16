package main

import (
	"fmt"
	"module"
	p11 "module1/package1"
	p12 "module1/package2"
	p21 "module2/package1"
	p22 "module2/package2"
)

func main() {
	fmt.Println("main")
	module.Package()
	p11.Package()
	p12.Package()
	p21.Package()
	p22.Package()
}

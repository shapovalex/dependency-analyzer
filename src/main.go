package main

import (
	"flag"
	"fmt"
)

func main() {
	var lFlag = flag.String("l", "", "Specify framework to use. Allowed values: pypi")
	var fFlag = flag.String("f", "", "Specify file/files to analyze")
	flag.Parse()
	fmt.Println(*lFlag)
	fmt.Println(*fFlag)
}

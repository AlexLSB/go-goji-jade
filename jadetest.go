package main

import (
	"fmt"
	"io/ioutil"

	"goji.io"
	"goji.io/pat"
)

func main() {
	dat, err := ioutil.ReadFile("templates/main.jade")
	if err != nil {
		fmt.Printf("ReadFile error: %v", err)
		return
	}

	tmpl, err := jade.Parse("name_of_tpl", string(dat))
	if err != nil {
		fmt.Printf("Parse error: %v", err)
		return
	}

	fmt.Printf("\nOutput:\n\n%s", hpp.PrPrint(tmpl))
}


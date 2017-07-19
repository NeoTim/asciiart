package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Print(renderASCIIToUnicode(string(input)))
}

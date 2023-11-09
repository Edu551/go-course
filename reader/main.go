package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	// t, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// if err := f.Close(); err != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)
	// }

	io.Copy(os.Stdout, f)
	// os.Stdout.Write(t)
}

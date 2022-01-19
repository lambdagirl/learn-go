package main

import "fmt"

const boilingF = 212

// this constant is a package-level declaration, throughtout all the files of the packages.

func main() {
	var f = boilingF // this variable is local to the function main
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
}

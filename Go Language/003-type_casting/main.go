package main

import (
	"fmt"
	"strconv"
)

/*
bool

string

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
	 // represents a Unicode code point

float32 float64

complex64 complex128
*/

func main() {
	var mystring = "1"
	var x = 10
	var f float32 = 2.0
	fmt.Println(mystring, x, f)
	number, _ := strconv.Atoi(mystring)
	fmt.Println(number)
}

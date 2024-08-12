package main

import (
	"fmt"
	"strings"
)
type StringProcessor interface {
	Length() int
	WordCount() int
}
type MyString struct {
	Text string
}

func (ms MyString) Length() int {
	return len(ms.Text)
}
func (ms MyString) WordCount() int {
	words := strings.Fields(ms.Text)
	return len(words)
}
type Formatter interface {
	ToUpper() string
	ToLower() string
}
type MyFormatter struct {
	Text string
}
type PointerOperations interface {
	Increment()
	Decrement()
}

type IntPointer struct {
	Value *int
}


func (ip *IntPointer) Increment() {
	*ip.Value++
}
func (ip *IntPointer) Decrement() {
	*ip.Value--
}
func (mf MyFormatter) ToUpper() string {
	return strings.ToUpper(mf.Text)
}
func (mf MyFormatter) ToLower() string {
	return strings.ToLower(mf.Text)
}

func main() {
	//1
	str1 := MyString{Text: "Hello, my brather "}
	fmt.Println("Length of string:", str1.Length())
	fmt.Println("Word count:", str1.WordCount())
	//2
	for1 := MyFormatter{Text: "Hello,Brather"}
	fmt.Println("Original string:", for1.Text)
	fmt.Println("Upper case:", for1.ToUpper())
	fmt.Println("Lower case:", for1.ToLower())
	//3
	num1 := 10
	intP1 := IntPointer{Value: &num1}
	fmt.Println("Initial value:", *intP1.Value)
	intP1.Increment()
	fmt.Println("After Increment:", *intP1.Value)
	intP1.Decrement()
	fmt.Println("After Decrement:", *intP1.Value)}
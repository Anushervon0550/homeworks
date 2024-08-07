package main

import (
	"fmt"
	"strings"
)
func Substring(s, substr string) bool {
    return strings.Contains(s, substr)}
func INDEX(s, substr string)int{
	return strings.Index(s,substr)}
func bb(original, old, new string) string {
		return strings.ReplaceAll(original, old, new)
}
func main() {
	//1
	src1 := "Privet "
	src2 := "Mir!"
	src3 := src1 + src2
	fmt.Println(src3)
	//2
	A:="Hello, world!"
	bytes := []byte(A)
	fmt.Println(bytes)
	runes :=[]rune(A)
	fmt.Println(runes)
	//3,4
	str := "Hello, world!"
    substr := "world"
    fmt.Println("Содержит ли строка подстроку?", Substring(str, substr))
	index:= INDEX(str, substr)
	if index != -1 {
        fmt.Printf("Подстрока найдена на индексе: %d\n", index)
    } else {
        fmt.Println("Подстрока не найдена")
    }
	//5
	original := "Hello, world! Hello, Go!"
    old := "Hello"
    new := "Hi"
    result := bb(original, old, new)
    fmt.Println(result)
}


